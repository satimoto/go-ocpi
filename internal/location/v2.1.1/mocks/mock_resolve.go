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
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *location.LocationResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *location.LocationResolver {
	return &location.LocationResolver{
		Repository:             locationMocks.NewRepository(repositoryService),
		OcpiRequester:          ocpiRequester,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		EnergyMixResolver:      energymix.NewResolver(repositoryService),
		EvseResolver:           evse.NewResolver(repositoryService),
		DisplayTextResolver:    displaytext.NewResolver(repositoryService),
		GeoLocationResolver:    geolocation.NewResolver(repositoryService),
		ImageResolver:          image.NewResolver(repositoryService),
		OpeningTimeResolver:    openingtime.NewResolver(repositoryService),
		TariffResolver:         tariff.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionDetailResolver:  versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
