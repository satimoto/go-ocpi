package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	calibration "github.com/satimoto/go-ocpi-api/internal/calibration/mocks"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1"
	chargingperiod "github.com/satimoto/go-ocpi-api/internal/chargingperiod/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *cdr.CdrResolver {
	repo := cdr.CdrRepository(repositoryService)

	return &cdr.CdrResolver{
		Repository:             repo,
		CalibrationResolver:    calibration.NewResolver(repositoryService),
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService, requester),
		TariffResolver:         tariff.NewResolver(repositoryService, requester),
	}
}
