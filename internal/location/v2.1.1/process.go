package location

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/displaytext"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/image"
)

func (r *LocationResolver) ReplaceLocation(ctx context.Context, credential db.Credential, uid string, locationDto *dto.LocationDto) *db.Location {
	if locationDto != nil {
		countryCode, partyID := evse.GetEvsesIdentity(locationDto, locationDto.Evses)

		return r.ReplaceLocationByIdentifier(ctx, credential, countryCode, partyID, uid, locationDto)
	}

	return nil
}

func (r *LocationResolver) ReplaceLocationByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, locationDto *dto.LocationDto) *db.Location {
	if locationDto != nil {
		location, err := r.Repository.GetLocationByUid(ctx, uid)
		geoLocationID := util.SqlNullInt64(util.NilInt64(location.GeoLocationID))
		energyMixID := location.EnergyMixID
		openingTimeID := location.OpeningTimeID
		operatorID := location.OperatorID
		ownerID := location.OwnerID
		suboperatorID := location.SuboperatorID
		geomPoint := &location.Geom

		if locationDto.Coordinates != nil {
			geomPoint = r.GeoLocationResolver.ReplaceGeoLocation(ctx, &geoLocationID, locationDto.Coordinates)
		}

		if !geoLocationID.Valid || geomPoint == nil {
			return nil
		}

		if locationDto.EnergyMix != nil {
			r.EnergyMixResolver.ReplaceEnergyMix(ctx, &energyMixID, locationDto.EnergyMix)
		}

		if locationDto.OpeningTimes != nil {
			r.OpeningTimeResolver.ReplaceOpeningTime(ctx, &openingTimeID, locationDto.OpeningTimes)
		}

		if locationDto.Operator != nil {
			r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, &operatorID, locationDto.Operator)
		}

		if locationDto.Owner != nil {
			r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, &ownerID, locationDto.Owner)
		}

		if locationDto.Suboperator != nil {
			r.BusinessDetailResolver.ReplaceBusinessDetail(ctx, &suboperatorID, locationDto.Suboperator)
		}

		if err == nil {
			locationParams := param.NewUpdateLocationByUidParams(location)
			locationParams.CountryCode = util.SqlNullString(countryCode)
			locationParams.PartyID = util.SqlNullString(partyID)
			locationParams.Geom = *geomPoint
			locationParams.GeoLocationID = geoLocationID.Int64
			locationParams.EnergyMixID = energyMixID
			locationParams.OpeningTimeID = openingTimeID
			locationParams.OperatorID = operatorID
			locationParams.OwnerID = ownerID
			locationParams.SuboperatorID = suboperatorID

			if locationDto.Address != nil {
				locationParams.Address = *locationDto.Address
			}

			if locationDto.City != nil {
				locationParams.City = *locationDto.City
			}

			if locationDto.ChargingWhenClosed != nil {
				locationParams.ChargingWhenClosed = *locationDto.ChargingWhenClosed
			}

			if locationDto.Country != nil {
				locationParams.Country = *locationDto.Country
			}

			if locationDto.LastUpdated != nil {
				locationParams.LastUpdated = *locationDto.LastUpdated
			}

			if locationDto.PostalCode != nil {
				locationParams.PostalCode = *locationDto.PostalCode
			}

			if locationDto.Name != nil {
				locationParams.Name = util.SqlNullString(locationDto.Name)
			}

			if locationDto.TimeZone != nil {
				locationParams.TimeZone = util.SqlNullString(locationDto.TimeZone)
			}

			if locationDto.Type != nil {
				locationParams.Type = *locationDto.Type
			}

			updatedLocation, err := r.Repository.UpdateLocationByUid(ctx, locationParams)

			if err != nil {
				util.LogOnError("OCPI117", "Error updating location", err)
				log.Printf("OCPI117: Params=%#v", locationParams)
				return nil
			}

			location = updatedLocation
		} else {
			locationParams := NewCreateLocationParams(locationDto)
			locationParams.CredentialID = credential.ID
			locationParams.CountryCode = util.SqlNullString(countryCode)
			locationParams.PartyID = util.SqlNullString(partyID)
			locationParams.Geom = *geomPoint
			locationParams.GeoLocationID = geoLocationID.Int64
			locationParams.EnergyMixID = energyMixID
			locationParams.OpeningTimeID = openingTimeID
			locationParams.OperatorID = operatorID
			locationParams.OwnerID = ownerID
			locationParams.SuboperatorID = suboperatorID

			location, err = r.Repository.CreateLocation(ctx, locationParams)

			if err != nil {
				util.LogOnError("OCPI118", "Error creating location", err)
				log.Printf("OCPI118: Params=%#v", locationParams)
				return nil
			}
		}

		if locationDto.Directions != nil {
			r.replaceDirections(ctx, location.ID, locationDto)
		}

		if locationDto.Facilities != nil {
			r.replaceFacilities(ctx, location.ID, locationDto)
		}

		if locationDto.Evses != nil {
			r.replaceEvses(ctx, location.ID, locationDto)
		}

		if locationDto.Images != nil {
			r.replaceImages(ctx, location.ID, locationDto)
		}

		if locationDto.RelatedLocations != nil {
			r.replaceRelatedLocations(ctx, location.ID, locationDto)
		}

		return &location
	}

	return nil
}

func (r *LocationResolver) ReplaceLocations(ctx context.Context, credential db.Credential, locationsDto []*dto.LocationDto) {
	for _, locationDto := range locationsDto {
		r.ReplaceLocation(ctx, credential, *locationDto.ID, locationDto)
	}
}

func (r *LocationResolver) ReplaceLocationsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, locationsDto []*dto.LocationDto) {
	for _, locationDto := range locationsDto {
		r.ReplaceLocationByIdentifier(ctx, credential, countryCode, partyID, *locationDto.ID, locationDto)
	}
}

func (r *LocationResolver) replaceDirections(ctx context.Context, locationID int64, locationDto *dto.LocationDto) {
	r.Repository.DeleteLocationDirections(ctx, locationID)

	for _, directionDto := range locationDto.Directions {
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

func (r *LocationResolver) replaceEvses(ctx context.Context, locationID int64, locationDto *dto.LocationDto) {
	r.EvseResolver.ReplaceEvses(ctx, locationID, locationDto.Evses)
}

func (r *LocationResolver) replaceFacilities(ctx context.Context, locationID int64, locationDto *dto.LocationDto) {
	r.Repository.UnsetLocationFacilities(ctx, locationID)

	if facilities, err := r.Repository.ListFacilities(ctx); err == nil {
		filteredFacilities := []db.Facility{}

		for _, facility := range facilities {
			if util.StringsContainString(locationDto.Facilities, facility.Text) {
				filteredFacilities = append(filteredFacilities, facility)
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

func (r *LocationResolver) replaceImages(ctx context.Context, locationID int64, locationDto *dto.LocationDto) {
	r.Repository.DeleteLocationImages(ctx, locationID)

	for _, imageDto := range locationDto.Images {
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

func (r *LocationResolver) replaceRelatedLocations(ctx context.Context, locationID int64, locationDto *dto.LocationDto) {
	r.Repository.DeleteAdditionalGeoLocations(ctx, locationID)

	for _, relatedLocation := range locationDto.RelatedLocations {
		additionalGeoLocationParams := NewCreateAdditionalGeoLocationParams(relatedLocation, locationID)

		if relatedLocation.Name != nil {
			displayTextParams := displaytext.NewCreateDisplayTextParams(relatedLocation.Name)
			displayText, err := r.DisplayTextResolver.Repository.CreateDisplayText(ctx, displayTextParams)

			if err != nil {
				util.LogOnError("OCPI123", "Error creating display text", err)
				log.Printf("OCPI123: LocationID=%v, Params=%#v", locationID, displayTextParams)
				continue
			}

			additionalGeoLocationParams.DisplayTextID = util.SqlNullInt64(displayText.ID)
		}

		_, err := r.Repository.CreateAdditionalGeoLocation(ctx, additionalGeoLocationParams)

		if err != nil {
			util.LogOnError("OCPI124", "Error creating additional geo location", err)
			log.Printf("OCPI124: Params=%#v", additionalGeoLocationParams)
		}
	}
}
