package nlcon

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
	TARIFF_UID_TEMPLATE = "connect.ned-%v"
)

func (r *NlConService) StartService(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting NlCon Service")
	r.shutdownCtx = shutdownCtx
	r.waitGroup = waitGroup
}

func (r *NlConService) Run(ctx context.Context, credential db.Credential, lastUpdated time.Time) {
	/*
	 * Update NlCon tariffs
	 * Get the updated EVSE tarifflist for NlCon.
	 * Loop through the EVSE list and update the connector tariffs.
	 */

	r.waitGroup.Add(1)
	log.Printf("Start NlCon Tariff sync")

	r.updateConnectors(ctx, credential)

	log.Printf("End NlCon Tariff sync")
	r.waitGroup.Done()
}

func (r *NlConService) updateConnectors(ctx context.Context, credential db.Credential) {
	url := "https://plonweuconspofa01.azurewebsites.net/api/v1/nl/con/tariffs"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		metrics.RecordError("OCPI311", "Error requesting nlcon prices", err)
		util.LogHttpRequest("OCPI311", url, request, false)
		return
	}

	response, err := r.HTTPRequester.Do(request)

	if err != nil {
		metrics.RecordError("OCPI312", "Error in nlcon prices response", err)
		util.LogHttpResponse("OCPI312", url, response, true)
		return
	}

	nlConTariffsDto, err := UnmarshalDto(response.Body)
	defer response.Body.Close()

	if err != nil {
		metrics.RecordError("OCPI313", "Error unmarshaling", err)
		util.LogHttpResponse("OCPI313", url, response, true)
		return
	}

	tariffMap := make(map[string]db.Tariff)

	for _, nlConTariffDto := range nlConTariffsDto {
		if _, ok := tariffMap[nlConTariffDto.TariffName]; !ok {
			if tariff, err := r.updateTariff(ctx, credential, nlConTariffDto); err == nil {
				tariffMap[nlConTariffDto.TariffName] = tariff
			}
		}

		if tariff, ok := tariffMap[nlConTariffDto.TariffName]; ok {
			connectors, err := r.ConnectorRepository.ListConnectorsByEvseID(ctx, util.SqlNullString(nlConTariffDto.EvseID))

			if err != nil {
				metrics.RecordError("OCPI314", "Error listing connectors by evse id", err)
				log.Printf("OCPI314: EvseID=%v", nlConTariffDto.EvseID)
				continue
			}

			for _, connector := range connectors {
				connectorParams := param.NewUpdateConnectorByEvseParams(connector)
				connectorParams.IsPublished = true
				connectorParams.TariffID = util.SqlNullString(tariff.Uid)

				if connector.TariffID.String != connectorParams.TariffID.String {
					_, err = r.ConnectorRepository.UpdateConnectorByEvse(ctx, connectorParams)

					if err != nil {
						metrics.RecordError("OCPI314", "Error updating connector", err)
						log.Printf("OCPI315: Params=%#v", connectorParams)
					}
				}
			}
		}
	}
}

func (r *NlConService) updateTariff(ctx context.Context, credential db.Credential, nlConTariffDto *NlConTariffDto) (db.Tariff, error) {
	likeTariffUid := fmt.Sprintf("%%%s", nlConTariffDto.TariffName)
	tariff, err := r.TariffRepository.GetTariffLikeUid(ctx, likeTariffUid)

	if err != nil {
		// Create the tariff
		tariffUid := fmt.Sprintf(TARIFF_UID_TEMPLATE, nlConTariffDto.TariffName)

		tariffParams := db.CreateTariffParams{
			Uid:          tariffUid,
			CredentialID: credential.ID,
			Currency:     "EUR",
			LastUpdated:  time.Now().UTC(),
		}

		tariff, err = r.TariffRepository.CreateTariff(ctx, tariffParams)

		if err != nil {
			metrics.RecordError("OCPI315", "Error creating tariff", err)
			log.Printf("OCPI315: Params=%#v", tariffParams)
			return db.Tariff{}, err
		}
	}

	// Construct the elements
	var elementsDto []*dto.ElementDto

	if nlConTariffDto.StartTariff > 0 {
		elementDto := dto.ElementDto{}
		elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
			Type:     db.TariffDimensionFLAT,
			Price:    nlConTariffDto.StartTariff,
			StepSize: 1,
		})

		elementsDto = append(elementsDto, &elementDto)
	}

	if nlConTariffDto.EnergyTariff > 0 {
		elementDto := dto.ElementDto{}
		elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
			Type:     db.TariffDimensionENERGY,
			Price:    nlConTariffDto.EnergyTariff,
			StepSize: 1,
		})

		elementsDto = append(elementsDto, &elementDto)
	}

	if nlConTariffDto.ParkingTariff != nil {
		elementDto := dto.ElementDto{}
		elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
			Type:     db.TariffDimensionTIME,
			Price:    *nlConTariffDto.ParkingTariff,
			StepSize: int32(util.DefaultInt(nlConTariffDto.ParkingStepSize, 1)),
		})

		if nlConTariffDto.ParkingGracePeriodInMinutes != nil {
			elementDto.Restrictions = &dto.ElementRestrictionDto{
				MinDuration: nlConTariffDto.ParkingGracePeriodInMinutes,
			}
		}

		elementsDto = append(elementsDto, &elementDto)
	}

	r.ElementResolver.ReplaceElements(ctx, tariff, elementsDto)

	return tariff, nil
}
