package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	businessdetail "github.com/satimoto/go-ocpi-api/internal/businessdetail/mocks"
	displaytext "github.com/satimoto/go-ocpi-api/internal/displaytext/mocks"
	energymix "github.com/satimoto/go-ocpi-api/internal/energymix/mocks"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1/mocks"
	geolocation "github.com/satimoto/go-ocpi-api/internal/geolocation/mocks"
	image "github.com/satimoto/go-ocpi-api/internal/image/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	openingtime "github.com/satimoto/go-ocpi-api/internal/openingtime/mocks"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *ocpi.OCPIRequester) *location.LocationResolver {
	repo := location.LocationRepository(repositoryService)

	return &location.LocationResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		EnergyMixResolver:      energymix.NewResolver(repositoryService),
		EvseResolver:           evse.NewResolver(repositoryService),
		DisplayTextResolver:    displaytext.NewResolver(repositoryService),
		GeoLocationResolver:    geolocation.NewResolver(repositoryService),
		ImageResolver:          image.NewResolver(repositoryService),
		OpeningTimeResolver:    openingtime.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService, requester),
	}
}
