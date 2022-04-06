package geolocation

import (
	"context"

	"github.com/satimoto/go-datastore/db"
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
