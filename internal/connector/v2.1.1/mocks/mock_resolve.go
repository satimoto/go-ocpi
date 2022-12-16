package mocks

import (
	connectorMocks "github.com/satimoto/go-datastore/pkg/connector/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1"
	party "github.com/satimoto/go-ocpi/internal/party/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *connector.ConnectorResolver {
	return &connector.ConnectorResolver{
		Repository:    connectorMocks.NewRepository(repositoryService),
		PartyResolver: party.NewResolver(repositoryService),
	}
}
