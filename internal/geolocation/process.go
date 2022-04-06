package geolocation

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
	"github.com/twpayne/go-geom"
)

func (r *GeoLocationResolver) ReplaceGeoLocation(ctx context.Context, id *sql.NullInt64, dto *GeoLocationDto) *geom.Point {
	var geomPoint *geom.Point

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

			if point, err := util.NewGeomPoint(geoLocation.Latitude, geoLocation.Longitude); err == nil {
				geomPoint = geom.NewPointFlat(geom.XY, point)
			}
		}
	}

	return geomPoint
}
