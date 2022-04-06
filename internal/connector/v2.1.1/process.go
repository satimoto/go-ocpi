package connector

import (
	"context"

	"github.com/satimoto/go-datastore/db"
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
				connectorParams.TariffID = util.SqlNullString(dto.TariffID)
			}

			if dto.LastUpdated != nil {
				connectorParams.LastUpdated = *dto.LastUpdated
			}

			connector, err = r.Repository.UpdateConnectorByUid(ctx, connectorParams)
		} else {
			connectorParams := NewCreateConnectorParams(evseID, dto)

			connector, err = r.Repository.CreateConnector(ctx, connectorParams)
		}

		return &connector
	}

	return nil
}

func (r *ConnectorResolver) ReplaceConnectors(ctx context.Context, evseID int64, dto []*ConnectorDto) {
	if dto != nil {
		for _, connector := range dto {
			if connector.Id != nil {
				r.ReplaceConnector(ctx, evseID, *connector.Id, connector)
			}
		}
	}
}
