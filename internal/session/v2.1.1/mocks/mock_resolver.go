package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	chargingperiod "github.com/satimoto/go-ocpi-api/internal/chargingperiod/mocks"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *session.SessionResolver {
	repo := session.SessionRepository(repositoryService)

	return &session.SessionResolver{
		Repository:             repo,
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService, requester),
		LocationResolver:       location.NewResolver(repositoryService, requester),
	}
}
