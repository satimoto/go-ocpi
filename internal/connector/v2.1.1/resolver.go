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
