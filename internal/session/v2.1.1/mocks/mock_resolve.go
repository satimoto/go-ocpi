package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	chargingperiod "github.com/satimoto/go-ocpi-api/internal/chargingperiod/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/mocks"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *ocpi.OCPIRequester) *session.SessionResolver {
	repo := session.SessionRepository(repositoryService)

	return &session.SessionResolver{
		Repository:                 repo,
		ChargingPeriodResolver:     chargingperiod.NewResolver(repositoryService),
		LocationResolver:           location.NewResolver(repositoryService, requester),
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
		VersionDetailResolver:      versiondetail.NewResolver(repositoryService, requester),
	}
}
