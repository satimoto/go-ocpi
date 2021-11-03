package geolocation

import (
	"context"
	"database/sql"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/util"
	"github.com/twpayne/go-geom"
)

type GeoLocationRepository interface {
	CreateGeoLocation(ctx context.Context, arg db.CreateGeoLocationParams) (db.GeoLocation, error)
	DeleteGeoLocation(ctx context.Context, id int64) error
	GetGeoLocation(ctx context.Context, id int64) (db.GeoLocation, error)
	UpdateGeoLocation(ctx context.Context, arg db.UpdateGeoLocationParams) (db.GeoLocation, error)
}

type GeoLocationResolver struct {
	Repository GeoLocationRepository
}

func NewResolver(repositoryService *db.RepositoryService) *GeoLocationResolver {
	repo := GeoLocationRepository(repositoryService)
	return &GeoLocationResolver{repo}
}

func (r *GeoLocationResolver) ReplaceGeoLocation(ctx context.Context, id *sql.NullInt64, payload *GeoLocationPayload) *geom.Point {
	var geomPoint *geom.Point

	if payload != nil {
		var geoLocation db.GeoLocation
		var err error

		if !id.Valid {
			geoLocationParams := NewCreateGeoLocationParams(payload)
			geoLocation, err = r.Repository.CreateGeoLocation(ctx, geoLocationParams)
		} else {
			geoLocationParams := NewUpdateGeoLocationParams(id.Int64, payload)
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

