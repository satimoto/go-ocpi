package itbec

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

const (
	TARIFF_UID_TEMPLATE = "be.charge-%v"
)

func (r *ItBecService) StartService(shutdownCtx context.Context, waitGroup *sync.WaitGroup) {
	log.Printf("Starting ItBec Service")
	r.shutdownCtx = shutdownCtx
	r.waitGroup = waitGroup
}

func (r *ItBecService) Run(ctx context.Context, credential db.Credential, lastUpdated time.Time) {
	/*
	 * Update ItBec tariffs
	 * Get the updated EVSE tarifflist for ItBec.
	 * Loop through the Connectors list and update the connector tariffs.
	 */

	r.waitGroup.Add(1)
	log.Printf("Start ItBec Tariff sync")

	r.updateTariffs(ctx, credential)
	r.updateConnectors(ctx, credential)

	log.Printf("End ItBec Tariff sync")
	r.waitGroup.Done()
}

func selectTariffUid(connector db.Connector) string {
	if connector.PowerType == db.PowerTypeDC {
		if connector.Voltage < 100 {
			return "a6f27f41-da68-4871-a8b3-bf4f42e6c227"
		} else if connector.Voltage >= 150 {
			return "d7797205-0c2f-4cf3-a5b3-9d858a64e72f"
		}

		return "d7797205-0c2f-4cf3-a5b3-9d858a64e72f"
	}

	return "cb1ee4e8-df82-4a42-90b8-f23bd35b32af"
}

func (r *ItBecService) updateConnectors(ctx context.Context, credential db.Credential) {
	listConnectorsByPartyAndCountryCodeParams := db.ListConnectorsByPartyAndCountryCodeParams{
		CountryCode: util.SqlNullString("IT"),
		PartyID:     util.SqlNullString("BEC"),
	}

	connectors, err := r.ConnectorRepository.ListConnectorsByPartyAndCountryCode(ctx, listConnectorsByPartyAndCountryCodeParams)

	if err != nil {
		metrics.RecordError("OCPI335", "Error listing connectors", err)
		log.Printf("OCPI335: Params=%#v", listConnectorsByPartyAndCountryCodeParams)
		return
	}

	for _, connector := range connectors {
		tariffUid := fmt.Sprintf("%%%s", selectTariffUid(connector))

		if !connector.TariffID.Valid || connector.TariffID.String != tariffUid {
			connectorParams := param.NewUpdateConnectorByEvseParams(connector)
			connectorParams.IsPublished = true
			connectorParams.TariffID = util.SqlNullString(tariffUid)

			_, err = r.ConnectorRepository.UpdateConnectorByEvse(ctx, connectorParams)

			if err != nil {
				metrics.RecordError("OCPI336", "Error updating connector", err)
				log.Printf("OCPI336: Params=%#v", connectorParams)
			}
		}
	}
}

func (r *ItBecService) updateTariffs(ctx context.Context, credential db.Credential) {
	r.updateTariff(ctx, credential, &ItBecTariffDto{
		TariffName:   "cb1ee4e8-df82-4a42-90b8-f23bd35b32af",
		EnergyTariff: 0.72,
	})
	r.updateTariff(ctx, credential, &ItBecTariffDto{
		TariffName:   "a6f27f41-da68-4871-a8b3-bf4f42e6c227",
		EnergyTariff: 0.77,
	})
	r.updateTariff(ctx, credential, &ItBecTariffDto{
		TariffName:   "d7797205-0c2f-4cf3-a5b3-9d858a64e72f",
		EnergyTariff: 0.83,
	})
}

func (r *ItBecService) updateTariff(ctx context.Context, credential db.Credential, itBecTariffDto *ItBecTariffDto) (db.Tariff, error) {
	tariffUid := fmt.Sprintf(TARIFF_UID_TEMPLATE, itBecTariffDto.TariffName)
	tariff, err := r.TariffRepository.GetTariffByUid(ctx, tariffUid)

	if err != nil {
		// Create the tariff
		tariffParams := db.CreateTariffParams{
			Uid:          tariffUid,
			CredentialID: credential.ID,
			Currency:     "EUR",
			LastUpdated:  time.Now().UTC(),
		}

		tariff, err = r.TariffRepository.CreateTariff(ctx, tariffParams)

		if err != nil {
			metrics.RecordError("OCPI337", "Error creating tariff", err)
			log.Printf("OCPI375: Params=%#v", tariffParams)
			return db.Tariff{}, err
		}

		// Construct the elements
		var elementsDto []*dto.ElementDto

		if itBecTariffDto.EnergyTariff > 0 {
			elementDto := dto.ElementDto{}
			elementDto.PriceComponents = append(elementDto.PriceComponents, &dto.PriceComponentDto{
				Type:     db.TariffDimensionENERGY,
				Price:    itBecTariffDto.EnergyTariff,
				StepSize: 1,
			})

			elementsDto = append(elementsDto, &elementDto)
		}

		r.ElementResolver.ReplaceElements(ctx, tariff, elementsDto)
	}

	return tariff, nil
}
