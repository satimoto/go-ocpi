package location

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/displaytext"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
	"github.com/satimoto/go-ocpi-api/internal/image"
)

func (r *LocationResolver) ReplaceLocation(ctx context.Context, credential db.Credential, uid string, dto *LocationDto) *db.Location {
	if dto != nil {
		countryCode, partyID := evse.GetEvsesIdentity(dto.Evses)

		return r.ReplaceLocationByIdentifier(ctx, credential, countryCode, partyID, uid, dto)
	}

	return nil
}

func (r *LocationResolver) ReplaceLocationByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, dto *LocationDto) *db.Location {
	if dto != nil {
		location, err := r.Repository.GetLocationByUid(ctx, uid)
		geoLocationID := util.SqlNullInt64(util.NilInt64(location.GeoLocationID))
		energyMixID := util.NilInt64(location.EnergyMixID)
		openingTimeID := util.NilInt64(location.OpeningTimeID)
		operatorID := util.NilInt64(location.OperatorID)
		ownerID := util.NilInt64(location.OwnerID)
		suboperatorID := util.NilInt64(location.SuboperatorID)
		geomPoint := &location.Geom

		if dto.Coordinates != nil {
			geomPoint = r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, dto.Coordinates)
		}

		if !geoLocationID.Valid || geomPoint == nil {
			return nil
		}

		if dto.EnergyMix != nil {
			r.EnergyMixResolver.ReplaceEnergyMix(ctx, energyMixID, dto.EnergyMix)
		}

		if dto.OpeningTimes != nil {
			r.OpeningTimeResolver.ReplaceOpeningTime(ctx, openingTimeID, dto.OpeningTimes)
		}

		if dto.Operator != nil {
			r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, operatorID, dto.Operator)
		}

		if dto.Owner != nil {
			r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, ownerID, dto.Owner)
		}

		if dto.Suboperator != nil {
			r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, suboperatorID, dto.Suboperator)
		}

		if err == nil {
			locationParams := NewUpdateLocationByUidParams(location)
			locationParams.CountryCode = util.SqlNullString(countryCode)
			locationParams.PartyID = util.SqlNullString(partyID)
			locationParams.Geom = *geomPoint
			locationParams.GeoLocationID = geoLocationID.Int64
			locationParams.EnergyMixID = util.SqlNullInt64(energyMixID)
			locationParams.OpeningTimeID = util.SqlNullInt64(openingTimeID)
			locationParams.OperatorID = util.SqlNullInt64(operatorID)
			locationParams.OwnerID = util.SqlNullInt64(ownerID)
			locationParams.SuboperatorID = util.SqlNullInt64(suboperatorID)

			if dto.Address != nil {
				locationParams.Address = *dto.Address
			}

			if dto.City != nil {
				locationParams.City = *dto.City
			}

			if dto.ChargingWhenClosed != nil {
				locationParams.ChargingWhenClosed = *dto.ChargingWhenClosed
			}

			if dto.Country != nil {
				locationParams.Country = *dto.Country
			}

			if dto.LastUpdated != nil {
				locationParams.LastUpdated = *dto.LastUpdated
			}

			if dto.PostalCode != nil {
				locationParams.PostalCode = *dto.PostalCode
			}

			if dto.Name != nil {
				locationParams.Name = util.SqlNullString(dto.Name)
			}

			if dto.TimeZone != nil {
				locationParams.TimeZone = util.SqlNullString(dto.TimeZone)
			}

			if dto.Type != nil {
				locationParams.Type = *dto.Type
			}

			location, err = r.Repository.UpdateLocationByUid(ctx, locationParams)
		} else {
			locationParams := NewCreateLocationParams(dto)
			locationParams.CredentialID = credential.ID
			locationParams.CountryCode = util.SqlNullString(countryCode)
			locationParams.PartyID = util.SqlNullString(partyID)
			locationParams.Geom = *geomPoint
			locationParams.GeoLocationID = geoLocationID.Int64
			locationParams.EnergyMixID = util.SqlNullInt64(energyMixID)
			locationParams.OpeningTimeID = util.SqlNullInt64(openingTimeID)
			locationParams.OperatorID = util.SqlNullInt64(operatorID)
			locationParams.OwnerID = util.SqlNullInt64(ownerID)
			locationParams.SuboperatorID = util.SqlNullInt64(suboperatorID)

			location, err = r.Repository.CreateLocation(ctx, locationParams)
		}

		if dto.Directions != nil {
			r.replaceDirections(ctx, location.ID, dto)
		}

		if dto.Facilities != nil {
			r.replaceFacilities(ctx, location.ID, dto)
		}

		if dto.Evses != nil {
			r.replaceEvses(ctx, location.ID, dto)
		}

		if dto.Images != nil {
			r.replaceImages(ctx, location.ID, dto)
		}

		if dto.RelatedLocations != nil {
			r.replaceRelatedLocations(ctx, location.ID, dto)
		}

		return &location
	}

	return nil
}

func (r *LocationResolver) ReplaceLocationsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, dto []*LocationDto) {
	for _, locationDto := range dto {
		r.ReplaceLocationByIdentifier(ctx, credential, countryCode, partyID, *locationDto.ID, locationDto)
	}
}

func (r *LocationResolver) replaceDirections(ctx context.Context, locationID int64, dto *LocationDto) {
	r.Repository.DeleteLocationDirections(ctx, locationID)

	for _, directionDto := range dto.Directions {
		displayTextParams := displaytext.NewCreateDisplayTextParams(directionDto)

		if displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams); err == nil {
			r.Repository.SetLocationDirection(ctx, db.SetLocationDirectionParams{
				LocationID:    locationID,
				DisplayTextID: displayText.ID,
			})
		}
	}
}

func (r *LocationResolver) replaceEvses(ctx context.Context, locationID int64, dto *LocationDto) {
	r.EvseResolver.ReplaceEvses(ctx, locationID, dto.Evses)
}

func (r *LocationResolver) replaceFacilities(ctx context.Context, locationID int64, dto *LocationDto) {
	r.Repository.UnsetLocationFacilities(ctx, locationID)

	if facilities, err := r.Repository.ListFacilities(ctx); err == nil {
		filteredFacilities := []*db.Facility{}

		for _, facility := range facilities {
			if util.StringsContainString(dto.Facilities, facility.Text) {
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

func (r *LocationResolver) replaceImages(ctx context.Context, locationID int64, dto *LocationDto) {
	r.Repository.DeleteLocationImages(ctx, locationID)

	for _, imageDto := range dto.Images {
		imageParams := image.NewCreateImageParams(imageDto)

		if image, err := r.ImageResolver.Repository.CreateImage(ctx, imageParams); err == nil {
			r.Repository.SetLocationImage(ctx, db.SetLocationImageParams{
				LocationID: locationID,
				ImageID:    image.ID,
			})
		}
	}
}

func (r *LocationResolver) replaceRelatedLocations(ctx context.Context, locationID int64, dto *LocationDto) {
	r.Repository.DeleteRelatedLocations(ctx, locationID)

	for _, relatedLocation := range dto.RelatedLocations {
		geoLocationParams := geolocation.NewCreateGeoLocationParams(relatedLocation)

		if geoLocation, err := r.GeoLocationResolver.Repository.CreateGeoLocation(ctx, geoLocationParams); err == nil {
			r.Repository.SetRelatedLocation(ctx, db.SetRelatedLocationParams{
				LocationID:    locationID,
				GeoLocationID: geoLocation.ID,
			})
		}
	}
}
