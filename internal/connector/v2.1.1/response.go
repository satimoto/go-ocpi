package connector

import (
	"context"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type ConnectorPayload struct {
	Id                 *string             `json:"id"`
	Standard           *db.ConnectorType   `json:"standard"`
	Format             *db.ConnectorFormat `json:"format"`
	PowerType          *db.PowerType       `json:"power_type"`
	Voltage            *int32              `json:"voltage"`
	Amperage           *int32              `json:"amperage"`
	TariffID           *string             `json:"tariff_id,omitempty"`
	TermsAndConditions *string             `json:"terms_and_conditions,omitempty"`
	LastUpdated        *time.Time          `json:"last_updated"`
}

func (r *ConnectorPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewConnectorPayload(connector db.Connector) *ConnectorPayload {
	return &ConnectorPayload{
		Id:                 &connector.Uid,
		Standard:           &connector.Standard,
		Format:             &connector.Format,
		PowerType:          &connector.PowerType,
		Voltage:            &connector.Voltage,
		Amperage:           &connector.Amperage,
		TariffID:           util.NilString(connector.TariffID.String),
		TermsAndConditions: util.NilString(connector.TermsAndConditions.String),
		LastUpdated:        &connector.LastUpdated,
	}
}

func NewCreateConnectorParams(evseID int64, payload *ConnectorPayload) db.CreateConnectorParams {
	return db.CreateConnectorParams{
		EvseID:             evseID,
		Uid:                *payload.Id,
		Standard:           *payload.Standard,
		Format:             *payload.Format,
		PowerType:          *payload.PowerType,
		Voltage:            *payload.Voltage,
		Amperage:           *payload.Amperage,
		TariffID:           util.SqlNullString(payload.TariffID),
		TermsAndConditions: util.SqlNullString(payload.TermsAndConditions),
		LastUpdated:        *payload.LastUpdated,
	}
}

func (r *ConnectorResolver) CreateConnectorPayload(ctx context.Context, connector db.Connector) *ConnectorPayload {
	return NewConnectorPayload(connector)
}

func (r *ConnectorResolver) CreateConnectorListPayload(ctx context.Context, connectors []db.Connector) []*ConnectorPayload {
	list := []*ConnectorPayload{}
	for _, connector := range connectors {
		list = append(list, r.CreateConnectorPayload(ctx, connector))
	}
	return list
}
