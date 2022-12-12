package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	node "github.com/satimoto/go-datastore/pkg/node/mocks"
	sessionMocks "github.com/satimoto/go-datastore/pkg/session/mocks"
	token "github.com/satimoto/go-datastore/pkg/token/mocks"
	tokenauthorization "github.com/satimoto/go-datastore/pkg/tokenauthorization/mocks"
	chargingperiod "github.com/satimoto/go-ocpi/internal/chargingperiod/mocks"
	command "github.com/satimoto/go-ocpi/internal/command/v2.1.1/mocks"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	session "github.com/satimoto/go-ocpi/internal/session/v2.1.1"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *session.SessionResolver {
	return &session.SessionResolver{
		Repository:                   sessionMocks.NewRepository(repositoryService),
		OcpiService:                  services.OcpiService,
		AsyncService:                 services.AsyncService,
		ChargingPeriodResolver:       chargingperiod.NewResolver(repositoryService),
		CommandResolver:              command.NewResolver(repositoryService, services),
		LocationResolver:             location.NewResolver(repositoryService, services),
		NodeRepository:               node.NewRepository(repositoryService),
		TokenRepository:              token.NewRepository(repositoryService),
		TokenAuthorizationRepository: tokenauthorization.NewRepository(repositoryService),
		VersionDetailResolver:        versiondetail.NewResolver(repositoryService, services),
	}
}
