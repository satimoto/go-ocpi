package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	calibration "github.com/satimoto/go-ocpi-api/internal/calibration/mocks"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1"
	chargingperiod "github.com/satimoto/go-ocpi-api/internal/chargingperiod/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1/mocks"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *cdr.CdrResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *cdr.CdrResolver {
	repo := cdr.CdrRepository(repositoryService)

	return &cdr.CdrResolver{
		Repository:             repo,
		OcpiRequester:          ocpiRequester,
		CalibrationResolver:    calibration.NewResolver(repositoryService),
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolverWithServices(repositoryService, ocpiRequester),
		TariffResolver:         tariff.NewResolverWithServices(repositoryService, ocpiRequester),
		TokenResolver:          token.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionDetailResolver:  versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
