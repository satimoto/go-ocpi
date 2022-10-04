package location

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateLocationParams(dto *LocationDto) db.CreateLocationParams {
	return db.CreateLocationParams{
		Uid:                *dto.ID,
		Type:               *dto.Type,
		Name:               util.SqlNullString(dto.Name),
		Address:            *dto.Address,
		City:               *dto.City,
		PostalCode:         *dto.PostalCode,
		Country:            *dto.Country,
		AvailableEvses:     0,
		TotalEvses:         0,
		IsRemoteCapable:    false,
		IsRfidCapable:      false,
		TimeZone:           util.SqlNullString(dto.TimeZone),
		ChargingWhenClosed: util.DefaultBool(dto.ChargingWhenClosed, true),
		LastUpdated:        *dto.LastUpdated,
	}
}

func NewCreateAdditionalGeoLocationParams(dto *AdditionalGeoLocationDto, locationID int64) db.CreateAdditionalGeoLocationParams {
	latitudeStr := dto.Latitude.String()
	longitudeStr := dto.Longitude.String()

	return db.CreateAdditionalGeoLocationParams{
		LocationID:     locationID,
		Latitude:       latitudeStr,
		LatitudeFloat:  util.ParseFloat64(latitudeStr, 0),
		Longitude:      longitudeStr,
		LongitudeFloat: util.ParseFloat64(longitudeStr, 0),
	}
}
