package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1"
	chargingperiod "github.com/satimoto/go-ocpi-api/internal/chargingperiod/mocks"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *cdr.CdrResolver {
	repo := cdr.CdrRepository(repositoryService)

	return &cdr.CdrResolver{
		Repository:             repo,
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService, requester),
		LocationResolver:       location.NewResolver(repositoryService, requester),
		TariffResolver:         tariff.NewResolver(repositoryService, requester),
	}
}
