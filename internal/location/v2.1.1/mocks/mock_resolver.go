package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	businessdetail "github.com/satimoto/go-ocpi-api/internal/businessdetail/mocks"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1/mocks"
	displaytext "github.com/satimoto/go-ocpi-api/internal/displaytext/mocks"
	energymix "github.com/satimoto/go-ocpi-api/internal/energymix/mocks"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1/mocks"
	geolocation "github.com/satimoto/go-ocpi-api/internal/geolocation/mocks"
	image "github.com/satimoto/go-ocpi-api/internal/image/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	openingtime "github.com/satimoto/go-ocpi-api/internal/openingtime/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *location.LocationResolver {
	repo := location.LocationRepository(repositoryService)

	return &location.LocationResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService, requester),
		EnergyMixResolver:      energymix.NewResolver(repositoryService),
		EvseResolver:           evse.NewResolver(repositoryService),
		DisplayTextResolver:    displaytext.NewResolver(repositoryService),
		GeoLocationResolver:    geolocation.NewResolver(repositoryService),
		ImageResolver:          image.NewResolver(repositoryService),
		OpeningTimeResolver:    openingtime.NewResolver(repositoryService),
	}
}
