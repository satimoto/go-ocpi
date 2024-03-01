package deise

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
	TARIFF_UID_TEMPLATE = "compleo-%v"
)

func (r *DeIseService) StartService(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting DeIse Service")
	r.shutdownCtx = shutdownCtx
	r.waitGroup = waitGroup
}

func (r *DeIseService) Run(ctx context.Context, credential db.Credential, lastUpdated time.Time) {
	/*
	 * Update DeIse tariffs
	 * Get the updated EVSE tarifflist for DeIse.
	 * Loop through the EVSE list and update the connector tariffs.
	 */

	r.waitGroup.Add(1)
	log.Printf("Start DeIse Tariff sync")

	r.updateConnectors(ctx, credential)

	log.Printf("End DeIse Tariff sync")
	r.waitGroup.Done()
}

func (r *DeIseService) updateConnectors(ctx context.Context, credential db.Credential) {
	url := "https://api.services-emobility.com/prices/public-offer/DE_ISE/OCPI/current"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		metrics.RecordError("OCPI316", "Error requesting DeIse prices", err)
		util.LogHttpRequest("OCPI316", url, request, false)
		return
	}

	response, err := r.HTTPRequester.Do(request)

	if err != nil {
		metrics.RecordError("OCPI317", "Error in DeIse prices response", err)
		util.LogHttpResponse("OCPI317", url, response, true)
		return
	}

	deIseTariffsDto, err := UnmarshalDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		metrics.RecordError("OCPI318", "Error unmarshaling", err)
		util.LogHttpResponse("OCPI318", url, response, true)
		return
	}

	tariffMap := make(map[string]db.Tariff)

	for _, deIseTariffDto := range deIseTariffsDto {
		if len(deIseTariffDto.EvseID) > 0 {
			if _, ok := tariffMap[deIseTariffDto.ActivePriceSet]; !ok {
				if tariff, err := r.updateTariff(ctx, credential, deIseTariffDto); err == nil {
					tariffMap[deIseTariffDto.ActivePriceSet] = tariff
				}
			}

			if tariff, ok := tariffMap[deIseTariffDto.ActivePriceSet]; ok {
				connectors, err := r.ConnectorRepository.ListConnectorsByEvseID(ctx, util.SqlNullString(deIseTariffDto.EvseID))

				if err != nil {
					metrics.RecordError("OCPI319", "Error listing connectors by evse id", err)
					log.Printf("OCPI319: EvseID=%v", deIseTariffDto.EvseID)
					continue
				}

				for _, connector := range connectors {
					connectorParams := param.NewUpdateConnectorByEvseParams(connector)
					connectorParams.IsPublished = true
					connectorParams.TariffID = util.SqlNullString(tariff.Uid)

					if connector.TariffID.String != connectorParams.TariffID.String {
						_, err = r.ConnectorRepository.UpdateConnectorByEvse(ctx, connectorParams)

						if err != nil {
							metrics.RecordError("OCPI320", "Error updating connector", err)
							log.Printf("OCPI320: Params=%#v", connectorParams)
						}
					}
				}
			}
		}
	}
}

func (r *DeIseService) updateTariff(ctx context.Context, credential db.Credential, deIseTariffDto *DeIseTariffDto) (db.Tariff, error) {
	likeTariffUid := fmt.Sprintf("%%%s", deIseTariffDto.ActivePriceSet)
	tariff, err := r.TariffRepository.GetTariffLikeUid(ctx, likeTariffUid)

	if err != nil {
		// Create the tariff
		tariffUid := fmt.Sprintf(TARIFF_UID_TEMPLATE, deIseTariffDto.ActivePriceSet)

		tariffParams := db.CreateTariffParams{
			Uid:          tariffUid,
			CredentialID: credential.ID,
			Currency:     "EUR",
			LastUpdated:  time.Now().UTC(),
		}

		tariff, err = r.TariffRepository.CreateTariff(ctx, tariffParams)

		if err != nil {
			metrics.RecordError("OCPI321", "Error creating tariff", err)
			log.Printf("OCPI321: Params=%#v", tariffParams)
			return db.Tariff{}, err
		}
	}

	// Construct the elements
	var elementsDto []*dto.ElementDto

	if deIseTariffDto.ActiveSessionPrice > 0.0 {
		elementDto := dto.ElementDto{}
		elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
			Type:     db.TariffDimensionFLAT,
			Price:    deIseTariffDto.ActiveSessionPrice,
			StepSize: 1,
		})

		elementsDto = append(elementsDto, &elementDto)
	}

	if deIseTariffDto.ActiveKwhPrice > 0.0 {
		elementDto := dto.ElementDto{}
		elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
			Type:     db.TariffDimensionENERGY,
			Price:    deIseTariffDto.ActiveKwhPrice,
			StepSize: 1,
		})

		elementsDto = append(elementsDto, &elementDto)
	}

	if deIseTariffDto.ActiveMinutePrice > 0 {
		elementDto := dto.ElementDto{}
		elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
			Type:     db.TariffDimensionSESSIONTIME,
			Price:    deIseTariffDto.ActiveMinutePrice,
			StepSize: 1,
		})

		if deIseTariffDto.ActiveMinuteDelay > 0 {
			elementDto.Restrictions = &dto.ElementRestrictionDto{
				MinDuration: &deIseTariffDto.ActiveMinuteDelay,
			}
		}

		elementsDto = append(elementsDto, &elementDto)
	}

	r.ElementResolver.ReplaceElements(ctx, tariff, elementsDto)

	return tariff, nil
}
