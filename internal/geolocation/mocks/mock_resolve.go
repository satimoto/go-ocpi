package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	geolocationMocks "github.com/satimoto/go-datastore/pkg/geolocation/mocks"
	"github.com/satimoto/go-ocpi-api/internal/geolocation"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *geolocation.GeoLocationResolver {
	return &geolocation.GeoLocationResolver{
		Repository: geolocationMocks.NewRepository(repositoryService),
	}
}
