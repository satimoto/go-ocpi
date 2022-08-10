package geolocation

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/geolocation"
)

type GeoLocationResolver struct {
	Repository geolocation.GeoLocationRepository
}

func NewResolver(repositoryService *db.RepositoryService) *GeoLocationResolver {
	return &GeoLocationResolver{
		Repository: geolocation.NewRepository(repositoryService),
	}
}
