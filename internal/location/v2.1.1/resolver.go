package location

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	"github.com/satimoto/go-ocpi-api/internal/energymix"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
	"github.com/satimoto/go-ocpi-api/internal/image"
	"github.com/satimoto/go-ocpi-api/internal/openingtime"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type LocationRepository interface {
	CreateLocation(ctx context.Context, arg db.CreateLocationParams) (db.Location, error)
	DeleteLocationDirections(ctx context.Context, locationID int64) error
	DeleteLocationImages(ctx context.Context, locationID int64) error
	DeleteRelatedLocations(ctx context.Context, locationID int64) error
	GetLocation(ctx context.Context, id int64) (db.Location, error)
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
	*credential.CredentialResolver
	*displaytext.DisplayTextResolver
	*energymix.EnergyMixResolver
	*evse.EvseResolver
	*geolocation.GeoLocationResolver
	*image.ImageResolver
	*openingtime.OpeningTimeResolver
}

func NewResolver(repositoryService *db.RepositoryService) *LocationResolver {
	repo := LocationRepository(repositoryService)
	return &LocationResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService),
		DisplayTextResolver:    displaytext.NewResolver(repositoryService),
		EnergyMixResolver:      energymix.NewResolver(repositoryService),
		EvseResolver:           evse.NewResolver(repositoryService),
		GeoLocationResolver:    geolocation.NewResolver(repositoryService),
		ImageResolver:          image.NewResolver(repositoryService),
		OpeningTimeResolver:    openingtime.NewResolver(repositoryService),
	}
}

func (r *LocationResolver) ReplaceLocation(ctx context.Context, uid string, payload *LocationPayload) *db.Location {
	if payload != nil {
		location, err := r.Repository.GetLocationByUid(ctx, uid)
		geoLocationID := util.SqlNullInt64(util.NilInt64(location.GeoLocationID))
		energyMixID := util.NilInt64(location.EnergyMixID)
		openingTimeID := util.NilInt64(location.OpeningTimeID)
		operatorID := util.NilInt64(location.OperatorID)
		ownerID := util.NilInt64(location.OwnerID)
		suboperatorID := util.NilInt64(location.SuboperatorID)
		geomPoint := &location.Geom

		if payload.Coordinates != nil {
			geomPoint = r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, payload.Coordinates)
		}

		if !geoLocationID.Valid || geomPoint == nil {
			return nil
		}

		if payload.EnergyMix != nil {
			r.EnergyMixResolver.ReplaceEnergyMix(ctx, energyMixID, payload.EnergyMix)
		}

		if payload.OpeningTimes != nil {
			r.OpeningTimeResolver.ReplaceOpeningTime(ctx, openingTimeID, payload.OpeningTimes)
		}

		if payload.Operator != nil {
			r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, operatorID, payload.Operator)
		}

		if payload.Owner != nil {
			r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, ownerID, payload.Owner)
		}

		if payload.Suboperator != nil {
			r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, suboperatorID, payload.Suboperator)
		}

		if err == nil {
			locationParams := NewUpdateLocationByUidParams(location)
			locationParams.Geom = *geomPoint
			locationParams.GeoLocationID = geoLocationID.Int64
			locationParams.EnergyMixID = util.SqlNullInt64(energyMixID)
			locationParams.OpeningTimeID = util.SqlNullInt64(openingTimeID)
			locationParams.OperatorID = util.SqlNullInt64(operatorID)
			locationParams.OwnerID = util.SqlNullInt64(ownerID)
			locationParams.SuboperatorID = util.SqlNullInt64(suboperatorID)

			if payload.Address != nil {
				locationParams.Address = *payload.Address
			}

			if payload.City != nil {
				locationParams.City = *payload.City
			}

			if payload.ChargingWhenClosed != nil {
				locationParams.ChargingWhenClosed = *payload.ChargingWhenClosed
			}

			if payload.Country != nil {
				locationParams.Country = *payload.Country
			}

			if payload.LastUpdated != nil {
				locationParams.LastUpdated = *payload.LastUpdated
			}

			if payload.PostalCode != nil {
				locationParams.PostalCode = *payload.PostalCode
			}

			if payload.Name != nil {
				locationParams.Name = util.SqlNullString(payload.Name)
			}

			if payload.TimeZone != nil {
				locationParams.TimeZone = util.SqlNullString(payload.TimeZone)
			}

			if payload.Type != nil {
				locationParams.Type = *payload.Type
			}

			location, err = r.Repository.UpdateLocationByUid(ctx, locationParams)
		} else {
			locationParams := NewCreateLocationParams(payload)
			locationParams.Geom = *geomPoint
			locationParams.GeoLocationID = geoLocationID.Int64
			locationParams.EnergyMixID = util.SqlNullInt64(energyMixID)
			locationParams.OpeningTimeID = util.SqlNullInt64(openingTimeID)
			locationParams.OperatorID = util.SqlNullInt64(operatorID)
			locationParams.OwnerID = util.SqlNullInt64(ownerID)
			locationParams.SuboperatorID = util.SqlNullInt64(suboperatorID)

			location, err = r.Repository.CreateLocation(ctx, locationParams)
		}

		if payload.Directions != nil {
			r.replaceDirections(ctx, location.ID, payload)
		}

		if payload.Facilities != nil {
			r.replaceFacilities(ctx, location.ID, payload)
		}

		if payload.Evses != nil {
			r.replaceEvses(ctx, location.ID, payload)
		}

		if payload.Images != nil {
			r.replaceImages(ctx, location.ID, payload)
		}

		if payload.RelatedLocations != nil {
			r.replaceRelatedLocations(ctx, location.ID, payload)
		}

		return &location
	}

	return nil
}

func (r *LocationResolver) replaceDirections(ctx context.Context, locationID int64, payload *LocationPayload) {
	r.Repository.DeleteLocationDirections(ctx, locationID)

	for _, directionPayload := range payload.Directions {
		displayTextParams := displaytext.NewCreateDisplayTextParams(directionPayload)

		if displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams); err == nil {
			r.Repository.SetLocationDirection(ctx, db.SetLocationDirectionParams{
				LocationID:    locationID,
				DisplayTextID: displayText.ID,
			})
		}
	}
}

func (r *LocationResolver) replaceEvses(ctx context.Context, locationID int64, payload *LocationPayload) {
	r.EvseResolver.ReplaceEvses(ctx, locationID, payload.Evses)
}

func (r *LocationResolver) replaceFacilities(ctx context.Context, locationID int64, payload *LocationPayload) {
	r.Repository.UnsetLocationFacilities(ctx, locationID)

	if facilities, err := r.Repository.ListFacilities(ctx); err == nil {
		filteredFacilities := []*db.Facility{}

		for _, facility := range facilities {
			if util.StringsContainString(payload.Facilities, facility.Text) {
				filteredFacilities = append(filteredFacilities, &facility)
			}
		}

		for _, facility := range filteredFacilities {
			r.Repository.SetLocationFacility(ctx, db.SetLocationFacilityParams{
				LocationID: locationID,
				FacilityID: facility.ID,
			})
		}
	}
}

func (r *LocationResolver) replaceImages(ctx context.Context, locationID int64, payload *LocationPayload) {
	r.Repository.DeleteLocationImages(ctx, locationID)

	for _, imagePayload := range payload.Images {
		imageParams := image.NewCreateImageParams(imagePayload)

		if image, err := r.ImageResolver.Repository.CreateImage(ctx, imageParams); err == nil {
			r.Repository.SetLocationImage(ctx, db.SetLocationImageParams{
				LocationID: locationID,
				ImageID:    image.ID,
			})
		}
	}
}

func (r *LocationResolver) replaceRelatedLocations(ctx context.Context, locationID int64, payload *LocationPayload) {
	r.Repository.DeleteRelatedLocations(ctx, locationID)

	for _, relatedLocation := range payload.RelatedLocations {
		geoLocationParams := geolocation.NewCreateGeoLocationParams(relatedLocation)

		if geoLocation, err := r.GeoLocationResolver.Repository.CreateGeoLocation(ctx, geoLocationParams); err == nil {
			r.Repository.SetRelatedLocation(ctx, db.SetRelatedLocationParams{
				LocationID:    locationID,
				GeoLocationID: geoLocation.ID,
			})
		}
	}
}
