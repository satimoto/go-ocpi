package connector

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type ConnectorRepository interface {
	CreateConnector(ctx context.Context, arg db.CreateConnectorParams) (db.Connector, error)
	GetConnectorByUid(ctx context.Context, arg db.GetConnectorByUidParams) (db.Connector, error)
	GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error)
	GetEvse(ctx context.Context, id int64) (db.Evse, error)
	ListConnectors(ctx context.Context, evseID int64) ([]db.Connector, error)
	UpdateConnectorByUid(ctx context.Context, arg db.UpdateConnectorByUidParams) (db.Connector, error)
	UpdateEvseLastUpdated(ctx context.Context, arg db.UpdateEvseLastUpdatedParams) error
	UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error
}

type ConnectorResolver struct {
	Repository ConnectorRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ConnectorResolver {
	repo := ConnectorRepository(repositoryService)
	return &ConnectorResolver{repo}
}

func (r *ConnectorResolver) ReplaceConnector(ctx context.Context, evseID int64, uid string, payload *ConnectorPayload) *db.Connector {
	if payload != nil {
		connector, err := r.Repository.GetConnectorByUid(ctx, db.GetConnectorByUidParams{
			EvseID: evseID,
			Uid:    uid,
		})

		if err == nil {
			connectorParams := db.NewUpdateConnectorByUidParams(connector)

			if payload.Standard != nil {
				connectorParams.Standard = *payload.Standard
			}

			if payload.Format != nil {
				connectorParams.Format = *payload.Format
			}

			if payload.PowerType != nil {
				connectorParams.PowerType = *payload.PowerType
			}

			if payload.Amperage != nil {
				connectorParams.Amperage = *payload.Amperage
			}

			if payload.Voltage != nil {
				connectorParams.Voltage = *payload.Voltage
			}

			if payload.TariffID != nil {
				connectorParams.TariffID = util.SqlNullString(payload.TariffID)
			}

			if payload.LastUpdated != nil {
				connectorParams.LastUpdated = *payload.LastUpdated
			}

			connector, err = r.Repository.UpdateConnectorByUid(ctx, connectorParams)
		} else {
			connectorParams := NewCreateConnectorParams(evseID, payload)

			connector, err = r.Repository.CreateConnector(ctx, connectorParams)
		}

		return &connector
	}

	return nil
}

func (r *ConnectorResolver) ReplaceConnectors(ctx context.Context, evseID int64, payload []*ConnectorPayload) {
	if payload != nil {
		for _, connector := range payload {
			if connector.Id != nil {
				r.ReplaceConnector(ctx, evseID, *connector.Id, connector)
			}
		}
	}
}
