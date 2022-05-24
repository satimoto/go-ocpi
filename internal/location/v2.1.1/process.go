package location

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
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
			locationParams := param.NewUpdateLocationByUidParams(location)
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

			updatedLocation, err := r.Repository.UpdateLocationByUid(ctx, locationParams)

			if err != nil {
				util.LogOnError("OCPI117", "Error updating location", err)
				log.Printf("OCPI117: Params=%#v", locationParams)
				return nil
			}

			location = updatedLocation
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

			if err != nil {
				util.LogOnError("OCPI118", "Error creating location", err)
				log.Printf("OCPI118: Params=%#v", locationParams)
				return nil
			}
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
			setLocationDirectionParams := db.SetLocationDirectionParams{
				LocationID:    locationID,
				DisplayTextID: displayText.ID,
			}
			err := r.Repository.SetLocationDirection(ctx, setLocationDirectionParams)

			if err != nil {
				util.LogOnError("OCPI119", "Error setting location direction", err)
				log.Printf("OCPI119: Params=%#v", setLocationDirectionParams)
			}
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
			setLocationFacilityParams := db.SetLocationFacilityParams{
				LocationID: locationID,
				FacilityID: facility.ID,
			}
			err := r.Repository.SetLocationFacility(ctx, setLocationFacilityParams)

			if err != nil {
				util.LogOnError("OCPI120", "Error setting location facility", err)
				log.Printf("OCPI120: Params=%#v", setLocationFacilityParams)
			}
		}
	}
}

func (r *LocationResolver) replaceImages(ctx context.Context, locationID int64, dto *LocationDto) {
	r.Repository.DeleteLocationImages(ctx, locationID)

	for _, imageDto := range dto.Images {
		imageParams := image.NewCreateImageParams(imageDto)
		image, err := r.ImageResolver.Repository.CreateImage(ctx, imageParams)

		if err != nil {
			util.LogOnError("OCPI121", "Error creating image", err)
			log.Printf("OCPI121: Params=%#v", imageParams)
			continue
		}

		setLocationImageParams := db.SetLocationImageParams{
			LocationID: locationID,
			ImageID:    image.ID,
		}
		err = r.Repository.SetLocationImage(ctx, setLocationImageParams)

		if err != nil {
			util.LogOnError("OCPI122", "Error setting location image", err)
			log.Printf("OCPI122: Params=%#v", setLocationImageParams)
		}

	}
}

func (r *LocationResolver) replaceRelatedLocations(ctx context.Context, locationID int64, dto *LocationDto) {
	r.Repository.DeleteRelatedLocations(ctx, locationID)

	for _, relatedLocation := range dto.RelatedLocations {
		geoLocationParams := geolocation.NewCreateGeoLocationParams(relatedLocation)
		geoLocation, err := r.GeoLocationResolver.Repository.CreateGeoLocation(ctx, geoLocationParams)

		if err != nil {
			util.LogOnError("OCPI123", "Error creating geolocation", err)
			log.Printf("OCPI123: Params=%#v", geoLocationParams)
			continue
		}

		setRelatedLocationParams := db.SetRelatedLocationParams{
			LocationID:    locationID,
			GeoLocationID: geoLocation.ID,
		}
		r.Repository.SetRelatedLocation(ctx, setRelatedLocationParams)

		if err != nil {
			util.LogOnError("OCPI124", "Error setting relation location", err)
			log.Printf("OCPI124: Params=%#v", setRelatedLocationParams)
		}
	}
}
