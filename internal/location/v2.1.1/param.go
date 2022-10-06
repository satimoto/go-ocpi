package location

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func NewCreateLocationParams(locationDto *dto.LocationDto) db.CreateLocationParams {
	return db.CreateLocationParams{
		Uid:                *locationDto.ID,
		Type:               *locationDto.Type,
		Name:               util.SqlNullString(locationDto.Name),
		Address:            *locationDto.Address,
		City:               *locationDto.City,
		PostalCode:         *locationDto.PostalCode,
		Country:            *locationDto.Country,
		AvailableEvses:     0,
		TotalEvses:         0,
		IsRemoteCapable:    false,
		IsRfidCapable:      false,
		TimeZone:           util.SqlNullString(locationDto.TimeZone),
		ChargingWhenClosed: util.DefaultBool(locationDto.ChargingWhenClosed, true),
		LastUpdated:        locationDto.LastUpdated.Time(),
	}
}

func NewCreateAdditionalGeoLocationParams(additionalGeoLocationDto *coreDto.AdditionalGeoLocationDto, locationID int64) db.CreateAdditionalGeoLocationParams {
	latitudeStr := additionalGeoLocationDto.Latitude.String()
	longitudeStr := additionalGeoLocationDto.Longitude.String()

	return db.CreateAdditionalGeoLocationParams{
		LocationID:     locationID,
		Latitude:       latitudeStr,
		LatitudeFloat:  util.ParseFloat64(latitudeStr, 0),
		Longitude:      longitudeStr,
		LongitudeFloat: util.ParseFloat64(longitudeStr, 0),
	}
}
