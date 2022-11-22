package htb

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

const (
	TARIFF_UID_TEMPLATE = "htb.solutions-%v"
)

func (r *HtbService) StartService(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting HasToBe Service")
	r.shutdownCtx = shutdownCtx
	r.waitGroup = waitGroup
}

func (r *HtbService) Run(ctx context.Context, credential db.Credential, lastUpdated time.Time) {
	/*
	 * Update HasToBe tariffs
	 * Get the HasToBe tariff list from the database and update the tariffs if changed.
	 * Get the updated EVSE list from HasToBe.
	 * Loop through the EVSE list and update the connector tariffs.
	 */

	r.waitGroup.Add(1)
	log.Printf("Start HasToBe Tariff sync")

	if ok := r.updateHtbTariffs(ctx, credential, lastUpdated); ok {
		r.updateConnectors(ctx)
	}

	log.Printf("End HasToBe Tariff sync")
	r.waitGroup.Done()
}

func (r *HtbService) updateConnectors(ctx context.Context) {
	url := "https://smart.htb.solutions/api/community/price_information"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		metrics.RecordError("OCPI296", "Error requesting htb prices", err)
		util.LogHttpRequest("OCPI296", url, request, false)
		return
	}

	response, err := r.HTTPRequester.Do(request)

	if err != nil {
		metrics.RecordError("OCPI297", "Error in htb prices response", err)
		util.LogHttpResponse("OCPI297", url, response, true)
		return
	}

	priceInformationDto, err := UnmarshalDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		metrics.RecordError("OCPI298", "Error unmarshaling", err)
		util.LogHttpResponse("OCPI298", url, response, true)
		return
	}

	for _, priceInformation := range priceInformationDto {
		connectors, err := r.ConnectorRepository.ListConnectorsByEvseID(ctx, util.SqlNullString(priceInformation.EvseID))

		if err != nil {
			metrics.RecordError("OCPI299", "Error listing connectors by evse id", err)
			log.Printf("OCPI299: EvseID=%v", priceInformation.EvseID)
			continue
		}

		tariffID := util.SqlNullString(nil)

		if priceInformation.Conditions != nil && priceInformation.Conditions.Rate != nil {
			tariffID = util.SqlNullString(fmt.Sprintf(TARIFF_UID_TEMPLATE, *priceInformation.Conditions.Rate))
		}

		for _, connector := range connectors {
			connectorParams := param.NewUpdateConnectorByEvseParams(connector)
			connectorParams.TariffID = tariffID

			if connector.TariffID.String != connectorParams.TariffID.String {
				_, err = r.ConnectorRepository.UpdateConnectorByEvse(ctx, connectorParams)

				if err != nil {
					metrics.RecordError("OCPI300", "Error updating connector", err)
					log.Printf("OCPI300: Params=%#v", connectorParams)
					continue
				}
			}
		}
	}
}

func (r *HtbService) updateHtbTariffs(ctx context.Context, credential db.Credential, lastUpdated time.Time) bool {
	htbTariffs, err := r.DataImportRepository.ListHtbTariffs(ctx)

	if err != nil {
		metrics.RecordError("OCPI294", "Error listing htb tariffs", err)
		return false
	}

	for _, htbTariff := range htbTariffs {
		if htbTariff.LastUpdated.After(lastUpdated) {
			r.updateTariff(ctx, credential, htbTariff)
		}
	}

	return true

}

func (r *HtbService) updateTariff(ctx context.Context, credential db.Credential, htbTariff db.HtbTariff) {
	tariffUid := fmt.Sprintf(TARIFF_UID_TEMPLATE, htbTariff.Name)
	tariff, err := r.TariffRepository.GetTariffByUid(ctx, tariffUid)

	if err != nil {
		// Create the tariff
		tariffParams := db.CreateTariffParams{
			Uid:                      tariffUid,
			CredentialID:             credential.ID,
			Currency:                 htbTariff.Currency,
			IsIntermediateCdrCapable: false,
			LastUpdated:              time.Now().UTC(),
		}

		tariff, err = r.TariffRepository.CreateTariff(ctx, tariffParams)

		if err != nil {
			metrics.RecordError("OCPI295", "Error creating tariff", err)
			log.Printf("OCPI295: Params=%#v", tariffParams)
			return
		}
	}

	// Construct the elements
	var elementsDto []*dto.ElementDto

	if htbTariff.FlatPrice.Valid {
		elementDto := dto.ElementDto{}
		elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
			Type:     db.TariffDimensionFLAT,
			Price:    htbTariff.TimePrice.Float64,
			StepSize: 1,
		})

		elementsDto = append(elementsDto, &elementDto)
	}

	if htbTariff.EnergyPrice.Valid {
		elementDto := dto.ElementDto{}
		elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
			Type:     db.TariffDimensionENERGY,
			Price:    htbTariff.EnergyPrice.Float64,
			StepSize: 1,
		})

		elementsDto = append(elementsDto, &elementDto)
	}

	if htbTariff.TimePrice.Valid {
		elementDto := dto.ElementDto{}
		elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
			Type:     db.TariffDimensionTIME,
			Price:    htbTariff.TimePrice.Float64,
			StepSize: 1,
		})

		if htbTariff.TimeMinDuration.Valid {
			elementDto.Restrictions = &dto.ElementRestrictionDto{
				MinDuration: &htbTariff.TimeMinDuration.Int32,
			}
		}

		elementsDto = append(elementsDto, &elementDto)
	}

	r.ElementResolver.ReplaceElements(ctx, tariff, elementsDto)
}
