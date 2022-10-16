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
	"github.com/satimoto/go-ocpi/internal/service"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1/mocks"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *cdr.CdrResolver {
	return &cdr.CdrResolver{
		Repository:             cdrMocks.NewRepository(repositoryService),
		OcpiService:            services.OcpiService,
		CalibrationResolver:    calibration.NewResolver(repositoryService),
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService, services),
		NodeRepository:         node.NewRepository(repositoryService),
		SessionRepository:      session.NewRepository(repositoryService),
		TariffResolver:         tariff.NewResolver(repositoryService, services),
		TokenRepository:        token.NewRepository(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService, services),
	}
}
