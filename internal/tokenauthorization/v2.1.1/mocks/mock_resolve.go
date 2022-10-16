package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	tokenauthorizationMocks "github.com/satimoto/go-datastore/pkg/tokenauthorization/mocks"
	user "github.com/satimoto/go-datastore/pkg/user/mocks"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1/mocks"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	tokenauthorization "github.com/satimoto/go-ocpi/internal/tokenauthorization/v2.1.1"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *tokenauthorization.TokenAuthorizationResolver {
	return &tokenauthorization.TokenAuthorizationResolver{
		Repository:          tokenauthorizationMocks.NewRepository(repositoryService),
		AsyncService:        services.AsyncService,
		NotificationService: services.NotificationService,
		ConnectorResolver:   connector.NewResolver(repositoryService),
		EvseResolver:        evse.NewResolver(repositoryService),
		UserRepository:      user.NewRepository(repositoryService),
	}
}
