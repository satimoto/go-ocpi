package connector

import (
	"github.com/satimoto/go-datastore/pkg/connector"
	"github.com/satimoto/go-datastore/pkg/db"
)

type ConnectorResolver struct {
	Repository connector.ConnectorRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ConnectorResolver {
	return &ConnectorResolver{
		Repository: connector.NewRepository(repositoryService),
	}
}
