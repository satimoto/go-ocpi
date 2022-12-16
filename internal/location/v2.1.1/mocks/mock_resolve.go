package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	locationMocks "github.com/satimoto/go-datastore/pkg/location/mocks"
	businessdetail "github.com/satimoto/go-ocpi/internal/businessdetail/mocks"
	displaytext "github.com/satimoto/go-ocpi/internal/displaytext/mocks"
	energymix "github.com/satimoto/go-ocpi/internal/energymix/mocks"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1/mocks"
	geolocation "github.com/satimoto/go-ocpi/internal/geolocation/mocks"
	image "github.com/satimoto/go-ocpi/internal/image/mocks"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	openingtime "github.com/satimoto/go-ocpi/internal/openingtime/mocks"
	party "github.com/satimoto/go-ocpi/internal/party/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1/mocks"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *location.LocationResolver {
	return &location.LocationResolver{
		Repository:             locationMocks.NewRepository(repositoryService),
		OcpiService:            services.OcpiService,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		DisplayTextResolver:    displaytext.NewResolver(repositoryService),
		EnergyMixResolver:      energymix.NewResolver(repositoryService),
		EvseResolver:           evse.NewResolver(repositoryService, services),
		GeoLocationResolver:    geolocation.NewResolver(repositoryService),
		ImageResolver:          image.NewResolver(repositoryService),
		OpeningTimeResolver:    openingtime.NewResolver(repositoryService),
		PartyResolver:          party.NewResolver(repositoryService),
		TariffResolver:         tariff.NewResolver(repositoryService, services),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService, services),
	}
}
