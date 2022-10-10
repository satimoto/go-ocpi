package connector

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/util"
)

func (r *ConnectorResolver) ReplaceConnector(ctx context.Context, evse db.Evse, uid string, connectorDto *dto.ConnectorDto) *db.Connector {
	if connectorDto != nil {
		connector, err := r.Repository.GetConnectorByEvse(ctx, db.GetConnectorByEvseParams{
			EvseID: evse.ID,
			Uid:    uid,
		})

		if err == nil {
			connectorParams := param.NewUpdateConnectorByEvseParams(connector)

			if connectorDto.Standard != nil {
				connectorParams.Standard = *connectorDto.Standard
			}

			if connectorDto.Format != nil {
				connectorParams.Format = *connectorDto.Format
			}

			if connectorDto.PowerType != nil {
				connectorParams.PowerType = *connectorDto.PowerType
			}

			if connectorDto.Amperage != nil {
				connectorParams.Amperage = *connectorDto.Amperage
			}

			if connectorDto.Voltage != nil {
				connectorParams.Voltage = *connectorDto.Voltage
			}

			if connectorDto.TariffID != nil {
				connectorParams.TariffID = dbUtil.SqlNullString(connectorDto.TariffID)
			}

			if connectorDto.LastUpdated != nil {
				connectorParams.LastUpdated = connectorDto.LastUpdated.Time()
			}

			connectorParams.Identifier = dbUtil.SqlNullString(GetConnectorIdentifier(evse, connectorDto))
			connectorParams.Publish = true
			connectorParams.Wattage = util.CalculateWattage(connectorParams.PowerType, connectorParams.Voltage, connectorParams.Amperage)
			updatedConnector, err := r.Repository.UpdateConnectorByEvse(ctx, connectorParams)

			if err != nil {
				dbUtil.LogOnError("OCPI079", "Error updating connector", err)
				log.Printf("OCPI079: Params=%#v", connectorParams)
				return nil
			}

			connector = updatedConnector
		} else {
			connectorParams := NewCreateConnectorParams(evse, connectorDto)
			connector, err = r.Repository.CreateConnector(ctx, connectorParams)

			if err != nil {
				dbUtil.LogOnError("OCPI080", "Error creating connector", err)
				log.Printf("OCPI080: Params=%#v", connectorParams)
				return nil
			}
		}

		return &connector
	}

	return nil
}

func (r *ConnectorResolver) ReplaceConnectors(ctx context.Context, evse db.Evse, connectorsDto []*dto.ConnectorDto) {
	if connectorsDto != nil {
		if connectors, err := r.Repository.ListConnectors(ctx, evse.ID); err == nil {
			connectorMap := make(map[string]db.Connector)

			for _, connector := range connectors {
				connectorMap[connector.Uid] = connector
			}

			for _, connectorDto := range connectorsDto {
				if connectorDto.Id != nil {
					r.ReplaceConnector(ctx, evse, *connectorDto.Id, connectorDto)
					delete(connectorMap, *connectorDto.Id)
				}
			}

			for _, connector := range connectorMap {
				connectorParams := param.NewUpdateConnectorByEvseParams(connector)
				connectorParams.Publish = false
				_, err := r.Repository.UpdateConnectorByEvse(ctx, connectorParams)

				if err != nil {
					dbUtil.LogOnError("OCPI113", "Error updating connector", err)
					log.Printf("OCPI113: Params=%#v", connectorParams)
				}
			}
		}
	}
}