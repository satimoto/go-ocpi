package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	tokenauthorizationMocks "github.com/satimoto/go-datastore/pkg/tokenauthorization/mocks"
	connector "github.com/satimoto/go-ocpi-api/internal/connector/v2.1.1/mocks"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1/mocks"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *tokenauthorization.TokenAuthorizationResolver {
	return &tokenauthorization.TokenAuthorizationResolver{
		Repository:        tokenauthorizationMocks.NewRepository(repositoryService),
		ConnectorResolver: connector.NewResolver(repositoryService),
		EvseResolver:      evse.NewResolver(repositoryService),
	}
}
