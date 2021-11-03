package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	connector "github.com/satimoto/go-ocpi-api/connector/v2.1.1/mocks"
	displaytext "github.com/satimoto/go-ocpi-api/displaytext/mocks"
	evse "github.com/satimoto/go-ocpi-api/evse/v2.1.1"
	geolocation "github.com/satimoto/go-ocpi-api/geolocation/mocks"
	image "github.com/satimoto/go-ocpi-api/image/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *evse.EvseResolver {
	repo := evse.EvseRepository(repositoryService)

	return &evse.EvseResolver{
		Repository:          repo,
		ConnectorResolver:   connector.NewResolver(repositoryService),
		DisplayTextResolver: displaytext.NewResolver(repositoryService),
		GeoLocationResolver: geolocation.NewResolver(repositoryService),
		ImageResolver:       image.NewResolver(repositoryService),
	}
}
