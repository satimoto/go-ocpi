package geolocation

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateGeoLocationParams(dto *GeoLocationDto) db.CreateGeoLocationParams {
	latitudeStr := dto.Latitude.String()
	longitudeStr := dto.Longitude.String()

	return db.CreateGeoLocationParams{
		Latitude:       latitudeStr,
		LatitudeFloat:  util.ParseFloat64(latitudeStr, 0),
		Longitude:      longitudeStr,
		LongitudeFloat: util.ParseFloat64(longitudeStr, 0),
		Name:           util.SqlNullString(dto.Name),
	}
}

func NewUpdateGeoLocationParams(id int64, dto *GeoLocationDto) db.UpdateGeoLocationParams {
	latitudeStr := dto.Latitude.String()
	longitudeStr := dto.Longitude.String()

	return db.UpdateGeoLocationParams{
		ID:             id,
		Latitude:       latitudeStr,
		LatitudeFloat:  util.ParseFloat64(latitudeStr, 0),
		Longitude:      longitudeStr,
		LongitudeFloat: util.ParseFloat64(longitudeStr, 0),
		Name:           util.SqlNullString(dto.Name),
	}
}
