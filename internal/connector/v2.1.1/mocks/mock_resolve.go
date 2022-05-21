package mocks

import (
	connectorMocks "github.com/satimoto/go-datastore/pkg/connector/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	connector "github.com/satimoto/go-ocpi-api/internal/connector/v2.1.1"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *connector.ConnectorResolver {
	return &connector.ConnectorResolver{
		Repository: connectorMocks.NewRepository(repositoryService),
	}
}
