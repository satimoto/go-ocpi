package connector

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func (r *ConnectorResolver) CreateConnectorDto(ctx context.Context, connector db.Connector) *dto.ConnectorDto {
	return dto.NewConnectorDto(connector)
}

func (r *ConnectorResolver) CreateConnectorListDto(ctx context.Context, connectors []db.Connector) []*dto.ConnectorDto {
	list := []*dto.ConnectorDto{}
	
	for _, connector := range connectors {
		list = append(list, r.CreateConnectorDto(ctx, connector))
	}

	return list
}
