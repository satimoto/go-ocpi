package mocks

import (
	cdrMocks "github.com/satimoto/go-datastore/pkg/cdr/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	node "github.com/satimoto/go-datastore/pkg/node/mocks"
	session "github.com/satimoto/go-datastore/pkg/session/mocks"
	token "github.com/satimoto/go-datastore/pkg/token/mocks"
	calibration "github.com/satimoto/go-ocpi/internal/calibration/mocks"
	cdr "github.com/satimoto/go-ocpi/internal/cdr/v2.1.1"
	chargingperiod "github.com/satimoto/go-ocpi/internal/chargingperiod/mocks"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1/mocks"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *cdr.CdrResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *cdr.CdrResolver {
	return &cdr.CdrResolver{
		Repository:             cdrMocks.NewRepository(repositoryService),
		OcpiRequester:          ocpiRequester,
		CalibrationResolver:    calibration.NewResolver(repositoryService),
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolverWithServices(repositoryService, ocpiRequester),
		NodeRepository:         node.NewRepository(repositoryService),
		SessionRepository:      session.NewRepository(repositoryService),
		TariffResolver:         tariff.NewResolverWithServices(repositoryService, ocpiRequester),
		TokenRepository:        token.NewRepository(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
