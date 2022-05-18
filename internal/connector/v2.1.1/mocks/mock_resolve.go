package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	connector "github.com/satimoto/go-ocpi-api/internal/connector/v2.1.1"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *connector.ConnectorResolver {
	repo := connector.ConnectorRepository(repositoryService)

	return &connector.ConnectorResolver{
		Repository: repo,
	}
}
