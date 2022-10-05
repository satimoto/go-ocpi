package geolocation

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateGeoLocationParams(geoLocationDto *coreDto.GeoLocationDto) db.CreateGeoLocationParams {
	latitudeStr := geoLocationDto.Latitude.String()
	longitudeStr := geoLocationDto.Longitude.String()

	return db.CreateGeoLocationParams{
		Latitude:       latitudeStr,
		LatitudeFloat:  util.ParseFloat64(latitudeStr, 0),
		Longitude:      longitudeStr,
		LongitudeFloat: util.ParseFloat64(longitudeStr, 0),
	}
}

func NewUpdateGeoLocationParams(id int64, geoLocationDto *coreDto.GeoLocationDto) db.UpdateGeoLocationParams {
	latitudeStr := geoLocationDto.Latitude.String()
	longitudeStr := geoLocationDto.Longitude.String()

	return db.UpdateGeoLocationParams{
		ID:             id,
		Latitude:       latitudeStr,
		LatitudeFloat:  util.ParseFloat64(latitudeStr, 0),
		Longitude:      longitudeStr,
		LongitudeFloat: util.ParseFloat64(longitudeStr, 0),
	}
}
