package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	connector "github.com/satimoto/go-ocpi-api/connector/v2.1.1"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *connector.ConnectorResolver {
	repo := connector.ConnectorRepository(repositoryService)

	return &connector.ConnectorResolver{
		Repository: repo,
	}
}