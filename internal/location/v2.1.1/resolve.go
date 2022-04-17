package location

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/energymix"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
	"github.com/satimoto/go-ocpi-api/internal/image"
	"github.com/satimoto/go-ocpi-api/internal/openingtime"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
)

type LocationRepository interface {
	CreateLocation(ctx context.Context, arg db.CreateLocationParams) (db.Location, error)
	DeleteLocationDirections(ctx context.Context, locationID int64) error
	DeleteLocationImages(ctx context.Context, locationID int64) error
	DeleteRelatedLocations(ctx context.Context, locationID int64) error
	GetLocation(ctx context.Context, id int64) (db.Location, error)
	GetLocationByLastUpdated(ctx context.Context, arg db.GetLocationByLastUpdatedParams) (db.Location, error)
	GetLocationByUid(ctx context.Context, uid string) (db.Location, error)
	ListEvses(ctx context.Context, locationID int64) ([]db.Evse, error)
	ListFacilities(ctx context.Context) ([]db.Facility, error)
	ListLocationDirections(ctx context.Context, locationID int64) ([]db.DisplayText, error)
	ListLocationFacilities(ctx context.Context, locationID int64) ([]db.Facility, error)
	ListLocationImages(ctx context.Context, locationID int64) ([]db.Image, error)
	ListLocations(ctx context.Context) ([]db.Location, error)
	ListRelatedLocations(ctx context.Context, locationID int64) ([]db.GeoLocation, error)
	SetLocationDirection(ctx context.Context, arg db.SetLocationDirectionParams) error
	SetLocationFacility(ctx context.Context, arg db.SetLocationFacilityParams) error
	SetLocationImage(ctx context.Context, arg db.SetLocationImageParams) error
	SetRelatedLocation(ctx context.Context, arg db.SetRelatedLocationParams) error
	UnsetLocationFacilities(ctx context.Context, locationID int64) error
	UpdateLocationByUid(ctx context.Context, arg db.UpdateLocationByUidParams) (db.Location, error)
	UpdateLocationLastUpdated(ctx context.Context, arg db.UpdateLocationLastUpdatedParams) error
}

type LocationResolver struct {
	Repository LocationRepository
	*businessdetail.BusinessDetailResolver
	*displaytext.DisplayTextResolver
	*energymix.EnergyMixResolver
	*evse.EvseResolver
	*geolocation.GeoLocationResolver
	*image.ImageResolver
	*openingtime.OpeningTimeResolver
	*versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *LocationResolver {
	repo := LocationRepository(repositoryService)
	return &LocationResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		DisplayTextResolver:    displaytext.NewResolver(repositoryService),
		EnergyMixResolver:      energymix.NewResolver(repositoryService),
		EvseResolver:           evse.NewResolver(repositoryService),
		GeoLocationResolver:    geolocation.NewResolver(repositoryService),
		ImageResolver:          image.NewResolver(repositoryService),
		OpeningTimeResolver:    openingtime.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService),
	}
}
