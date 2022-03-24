package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/element"
	pricecomponent "github.com/satimoto/go-ocpi-api/internal/pricecomponent/mocks"
	restriction "github.com/satimoto/go-ocpi-api/internal/restriction/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *element.ElementResolver {
	repo := element.ElementRepository(repositoryService)

	return &element.ElementResolver{
		Repository:             repo,
		PriceComponentResolver: pricecomponent.NewResolver(repositoryService),
		RestrictionResolver:    restriction.NewResolver(repositoryService),
	}
}
