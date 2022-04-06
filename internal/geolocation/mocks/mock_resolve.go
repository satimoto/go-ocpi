package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *geolocation.GeoLocationResolver {
	repo := geolocation.GeoLocationRepository(repositoryService)

	return &geolocation.GeoLocationResolver{
		Repository: repo,
	}
}
