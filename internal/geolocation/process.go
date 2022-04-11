package geolocation

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/postgis"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func (r *GeoLocationResolver) ReplaceGeoLocation(ctx context.Context, id *sql.NullInt64, dto *GeoLocationDto) *postgis.Geometry4326 {
	if dto != nil {
		var geoLocation db.GeoLocation
		var err error

		if !id.Valid {
			geoLocationParams := NewCreateGeoLocationParams(dto)
			geoLocation, err = r.Repository.CreateGeoLocation(ctx, geoLocationParams)
		} else {
			geoLocationParams := NewUpdateGeoLocationParams(id.Int64, dto)
			geoLocation, err = r.Repository.UpdateGeoLocation(ctx, geoLocationParams)
		}

		if err == nil {
			id.Scan(geoLocation.ID)

			if point, err := util.NewPoint(geoLocation.Latitude, geoLocation.Longitude); err == nil {
				return &postgis.Geometry4326{
					Coordinates: point,
					Type: point.GeoJSONType(),
				}
			}
		}
	}

	return nil
}
