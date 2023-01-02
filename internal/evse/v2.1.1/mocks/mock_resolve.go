package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	evseMocks "github.com/satimoto/go-datastore/pkg/evse/mocks"
	location "github.com/satimoto/go-datastore/pkg/location/mocks"
	node "github.com/satimoto/go-datastore/pkg/node/mocks"
	session "github.com/satimoto/go-datastore/pkg/session/mocks"
	tariff "github.com/satimoto/go-datastore/pkg/tariff/mocks"
	tokenauthorization "github.com/satimoto/go-datastore/pkg/tokenauthorization/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1/mocks"
	displaytext "github.com/satimoto/go-ocpi/internal/displaytext/mocks"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	geolocation "github.com/satimoto/go-ocpi/internal/geolocation/mocks"
	image "github.com/satimoto/go-ocpi/internal/image/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *evse.EvseResolver {
	recordEvseStatusPeriods := util.GetEnvBool("RECORD_EVSE_STATUS_PERIODS", true)

	return &evse.EvseResolver{
		Repository:                   evseMocks.NewRepository(repositoryService),
		OcpiService:                  services.OcpiService,
		ConnectorResolver:            connector.NewResolver(repositoryService),
		DisplayTextResolver:          displaytext.NewResolver(repositoryService),
		GeoLocationResolver:          geolocation.NewResolver(repositoryService),
		ImageResolver:                image.NewResolver(repositoryService),
		LocationRepository:           location.NewRepository(repositoryService),
		NodeRepository:               node.NewRepository(repositoryService),
		SessionRepository:            session.NewRepository(repositoryService),
		TariffRespository:            tariff.NewRepository(repositoryService),
		TokenAuthorizationRepository: tokenauthorization.NewRepository(repositoryService),
		VersionDetailResolver:        versiondetail.NewResolver(repositoryService, services),
		RecordEvseStatusPeriods:      recordEvseStatusPeriods,
	}
}
