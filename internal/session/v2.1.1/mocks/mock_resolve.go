package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	node "github.com/satimoto/go-datastore/pkg/node/mocks"
	sessionMocks "github.com/satimoto/go-datastore/pkg/session/mocks"
	chargingperiod "github.com/satimoto/go-ocpi-api/internal/chargingperiod/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1/mocks"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *session.SessionResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *session.SessionResolver {
	return &session.SessionResolver{
		Repository:                 sessionMocks.NewRepository(repositoryService),
		OcpiRequester:              ocpiRequester,
		ChargingPeriodResolver:     chargingperiod.NewResolver(repositoryService),
		LocationResolver:           location.NewResolverWithServices(repositoryService, ocpiRequester),
		NodeRepository:             node.NewRepository(repositoryService),
		TokenResolver:              token.NewResolverWithServices(repositoryService, ocpiRequester),
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
		VersionDetailResolver:      versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
