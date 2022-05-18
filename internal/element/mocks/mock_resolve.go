package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi-api/internal/element"
	elementrestriction "github.com/satimoto/go-ocpi-api/internal/elementrestriction/mocks"
	pricecomponent "github.com/satimoto/go-ocpi-api/internal/pricecomponent/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *element.ElementResolver {
	repo := element.ElementRepository(repositoryService)

	return &element.ElementResolver{
		Repository:                 repo,
		ElementRestrictionResolver: elementrestriction.NewResolver(repositoryService),
		PriceComponentResolver:     pricecomponent.NewResolver(repositoryService),
	}
}
