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
	metrics "github.com/satimoto/go-ocpi/internal/metric"
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

	r.updateConnectors(ctx)

	log.Printf("End NlCon Tariff sync")
	r.waitGroup.Done()
}

func (r *NlConService) updateConnectors(ctx context.Context) {
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

	for _, nlConTariffDto := range nlConTariffsDto {
		likeTariffUid := fmt.Sprintf("%%%s", nlConTariffDto.TariffName)

		if tariff, err := r.TariffRepository.GetTariffLikeUid(ctx, likeTariffUid); err == nil {
			connectors, err := r.ConnectorRepository.ListConnectorsByEvseID(ctx, util.SqlNullString(nlConTariffDto.EvseID))

			if err != nil {
				metrics.RecordError("OCPI314", "Error listing connectors by evse id", err)
				log.Printf("OCPI314: EvseID=%v", nlConTariffDto.EvseID)
				continue
			}

			for _, connector := range connectors {
				connectorParams := param.NewUpdateConnectorByEvseParams(connector)
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
