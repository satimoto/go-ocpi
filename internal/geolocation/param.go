package geolocation

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateGeoLocationParams(dto *GeoLocationDto) db.CreateGeoLocationParams {
	return db.CreateGeoLocationParams{
		Latitude:       dto.Latitude,
		LatitudeFloat:  util.ParseFloat64(dto.Latitude, 0),
		Longitude:      dto.Longitude,
		LongitudeFloat: util.ParseFloat64(dto.Longitude, 0),
		Name:           util.SqlNullString(dto.Name),
	}
}

func NewUpdateGeoLocationParams(id int64, dto *GeoLocationDto) db.UpdateGeoLocationParams {
	return db.UpdateGeoLocationParams{
		ID:             id,
		Latitude:       dto.Latitude,
		LatitudeFloat:  util.ParseFloat64(dto.Latitude, 0),
		Longitude:      dto.Longitude,
		LongitudeFloat: util.ParseFloat64(dto.Longitude, 0),
		Name:           util.SqlNullString(dto.Name),
	}
}
