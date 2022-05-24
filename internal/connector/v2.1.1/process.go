package connector

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func (r *ConnectorResolver) ReplaceConnector(ctx context.Context, evseID int64, uid string, dto *ConnectorDto) *db.Connector {
	if dto != nil {
		connector, err := r.Repository.GetConnectorByUid(ctx, db.GetConnectorByUidParams{
			EvseID: evseID,
			Uid:    uid,
		})

		if err == nil {
			connectorParams := db.NewUpdateConnectorByUidParams(connector)

			if dto.Standard != nil {
				connectorParams.Standard = *dto.Standard
			}

			if dto.Format != nil {
				connectorParams.Format = *dto.Format
			}

			if dto.PowerType != nil {
				connectorParams.PowerType = *dto.PowerType
			}

			if dto.Amperage != nil {
				connectorParams.Amperage = *dto.Amperage
			}

			if dto.Voltage != nil {
				connectorParams.Voltage = *dto.Voltage
			}

			if dto.TariffID != nil {
				connectorParams.TariffID = dbUtil.SqlNullString(dto.TariffID)
			}

			if dto.LastUpdated != nil {
				connectorParams.LastUpdated = *dto.LastUpdated
			}

			connectorParams.Wattage = util.CalculateWattage(connectorParams.PowerType, connectorParams.Voltage, connectorParams.Amperage)
			updatedConnector, err := r.Repository.UpdateConnectorByUid(ctx, connectorParams)

			if err != nil {
				dbUtil.LogOnError("OCPI079", "Error updating connector", err)
				log.Printf("OCPI079: Params=%v", connectorParams)
				return nil
			}

			connector = updatedConnector
		} else {
			connectorParams := NewCreateConnectorParams(evseID, dto)
			connector, err = r.Repository.CreateConnector(ctx, connectorParams)

			if err != nil {
				dbUtil.LogOnError("OCPI080", "Error creating connector", err)
				log.Printf("OCPI080: Params=%v", connectorParams)
				return nil
			}
		}

		return &connector
	}

	return nil
}

func (r *ConnectorResolver) ReplaceConnectors(ctx context.Context, evseID int64, dto []*ConnectorDto) {
	for _, connector := range dto {
		if connector.Id != nil {
			r.ReplaceConnector(ctx, evseID, *connector.Id, connector)
		}
	}
}
