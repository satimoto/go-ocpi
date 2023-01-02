package connector

import (
	"github.com/satimoto/go-datastore/pkg/connector"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/party"
)

type ConnectorResolver struct {
	Repository    connector.ConnectorRepository
	PartyResolver *party.PartyResolver
}

func NewResolver(repositoryService *db.RepositoryService) *ConnectorResolver {
	return &ConnectorResolver{
		Repository:    connector.NewRepository(repositoryService),
		PartyResolver: party.NewResolver(repositoryService),
	}
}
