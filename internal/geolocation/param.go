package geolocation

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewCreateGeoLocationParams(dto *GeoLocationDto) db.CreateGeoLocationParams {
	return db.CreateGeoLocationParams{
		Latitude:  dto.Latitude,
		Longitude: dto.Longitude,
		Name:      util.SqlNullString(dto.Name),
	}
}

func NewUpdateGeoLocationParams(id int64, dto *GeoLocationDto) db.UpdateGeoLocationParams {
	return db.UpdateGeoLocationParams{
		ID:        id,
		Latitude:  dto.Latitude,
		Longitude: dto.Longitude,
		Name:      util.SqlNullString(dto.Name),
	}
}