package connector

import (
	"github.com/satimoto/go-datastore/pkg/connector"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/party"
)

type ConnectorResolver struct {
	Repository      connector.ConnectorRepository
	PartyRepository party.PartyRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ConnectorResolver {
	return &ConnectorResolver{
		Repository:      connector.NewRepository(repositoryService),
		PartyRepository: party.NewRepository(repositoryService),
	}
}
