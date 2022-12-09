package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	location "github.com/satimoto/go-datastore/pkg/location/mocks"
	node "github.com/satimoto/go-datastore/pkg/node/mocks"
	session "github.com/satimoto/go-datastore/pkg/session/mocks"
	tariff "github.com/satimoto/go-datastore/pkg/tariff/mocks"
	tokenauthorizationMocks "github.com/satimoto/go-datastore/pkg/tokenauthorization/mocks"
	user "github.com/satimoto/go-datastore/pkg/user/mocks"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1/mocks"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	tokenauthorization "github.com/satimoto/go-ocpi/internal/tokenauthorization/v2.1.1"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *tokenauthorization.TokenAuthorizationResolver {
	return &tokenauthorization.TokenAuthorizationResolver{
		Repository:            tokenauthorizationMocks.NewRepository(repositoryService),
		OcpiService:           services.OcpiService,
		AsyncService:          services.AsyncService,
		NotificationService:   services.NotificationService,
		ConnectorResolver:     connector.NewResolver(repositoryService),
		EvseResolver:          evse.NewResolver(repositoryService),
		LocationRepository:    location.NewRepository(repositoryService),
		NodeRepository:        node.NewRepository(repositoryService),
		SessionRepository:     session.NewRepository(repositoryService),
		TariffRespository:     tariff.NewRepository(repositoryService),
		UserRepository:        user.NewRepository(repositoryService),
		VersionDetailResolver: versiondetail.NewResolver(repositoryService, services),
	}
}
