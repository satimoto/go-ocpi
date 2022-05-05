package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	calibration "github.com/satimoto/go-ocpi-api/internal/calibration/mocks"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1"
	chargingperiod "github.com/satimoto/go-ocpi-api/internal/chargingperiod/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1/mocks"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OCPIRequester) *cdr.CdrResolver {
	repo := cdr.CdrRepository(repositoryService)

	return &cdr.CdrResolver{
		Repository:             repo,
		OCPIRequester:          ocpiRequester,
		CalibrationResolver:    calibration.NewResolver(repositoryService),
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService, ocpiRequester),
		TariffResolver:         tariff.NewResolver(repositoryService, ocpiRequester),
		TokenResolver:          token.NewResolver(repositoryService, ocpiRequester),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService, ocpiRequester),
	}
}
