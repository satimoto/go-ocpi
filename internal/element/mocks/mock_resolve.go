package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	elementMocks "github.com/satimoto/go-datastore/pkg/element/mocks"
	"github.com/satimoto/go-ocpi-api/internal/element"
	elementrestriction "github.com/satimoto/go-ocpi-api/internal/elementrestriction/mocks"
	pricecomponent "github.com/satimoto/go-ocpi-api/internal/pricecomponent/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *element.ElementResolver {
	return &element.ElementResolver{
		Repository:                 elementMocks.NewRepository(repositoryService),
		ElementRestrictionResolver: elementrestriction.NewResolver(repositoryService),
		PriceComponentResolver:     pricecomponent.NewResolver(repositoryService),
	}
}
