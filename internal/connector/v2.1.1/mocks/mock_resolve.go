package mocks

import (
	connectorMocks "github.com/satimoto/go-datastore/pkg/connector/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	party "github.com/satimoto/go-datastore/pkg/party/mocks"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *connector.ConnectorResolver {
	return &connector.ConnectorResolver{
		Repository:      connectorMocks.NewRepository(repositoryService),
		PartyRepository: party.NewRepository(repositoryService),
	}
}
