package location

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/location"
	"github.com/satimoto/go-ocpi/internal/businessdetail"
	"github.com/satimoto/go-ocpi/internal/displaytext"
	"github.com/satimoto/go-ocpi/internal/energymix"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/geolocation"
	"github.com/satimoto/go-ocpi/internal/image"
	"github.com/satimoto/go-ocpi/internal/openingtime"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type LocationResolver struct {
	Repository             location.LocationRepository
	OcpiRequester          *transportation.OcpiRequester
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	DisplayTextResolver    *displaytext.DisplayTextResolver
	EnergyMixResolver      *energymix.EnergyMixResolver
	EvseResolver           *evse.EvseResolver
	GeoLocationResolver    *geolocation.GeoLocationResolver
	ImageResolver          *image.ImageResolver
	OpeningTimeResolver    *openingtime.OpeningTimeResolver
	TariffResolver         *tariff.TariffResolver
	VersionDetailResolver  *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *LocationResolver {
	return &LocationResolver{
		Repository:             location.NewRepository(repositoryService),
		OcpiRequester:          ocpiRequester,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		DisplayTextResolver:    displaytext.NewResolver(repositoryService),
		EnergyMixResolver:      energymix.NewResolver(repositoryService),
		EvseResolver:           evse.NewResolver(repositoryService),
		GeoLocationResolver:    geolocation.NewResolver(repositoryService),
		ImageResolver:          image.NewResolver(repositoryService),
		OpeningTimeResolver:    openingtime.NewResolver(repositoryService),
		TariffResolver:         tariff.NewResolver(repositoryService, ocpiRequester),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService, ocpiRequester),
	}
}
