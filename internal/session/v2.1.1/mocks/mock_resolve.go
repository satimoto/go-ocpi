package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	node "github.com/satimoto/go-datastore/pkg/node/mocks"
	sessionMocks "github.com/satimoto/go-datastore/pkg/session/mocks"
	token "github.com/satimoto/go-datastore/pkg/token/mocks"
	tokenauthorization "github.com/satimoto/go-datastore/pkg/tokenauthorization/mocks"
	chargingperiod "github.com/satimoto/go-ocpi/internal/chargingperiod/mocks"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1/mocks"
	session "github.com/satimoto/go-ocpi/internal/session/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
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
		VersionDetailResolver:      versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
		TokenRepository:              token.NewRepository(repositoryService),
		TokenAuthorizationRepository: tokenauthorization.NewRepository(repositoryService),
	}
}
