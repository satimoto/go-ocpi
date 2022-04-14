package connector

import (
	"context"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

type ConnectorDto struct {
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

func (r *ConnectorDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewConnectorDto(connector db.Connector) *ConnectorDto {
	return &ConnectorDto{
		Id:                 &connector.Uid,
		Standard:           &connector.Standard,
		Format:             &connector.Format,
		PowerType:          &connector.PowerType,
		Voltage:            &connector.Voltage,
		Amperage:           &connector.Amperage,
		TariffID:           util.NilString(connector.TariffID),
		TermsAndConditions: util.NilString(connector.TermsAndConditions),
		LastUpdated:        &connector.LastUpdated,
	}
}

func (r *ConnectorResolver) CreateConnectorDto(ctx context.Context, connector db.Connector) *ConnectorDto {
	return NewConnectorDto(connector)
}

func (r *ConnectorResolver) CreateConnectorListDto(ctx context.Context, connectors []db.Connector) []*ConnectorDto {
	list := []*ConnectorDto{}
	for _, connector := range connectors {
		list = append(list, r.CreateConnectorDto(ctx, connector))
	}
	return list
}
