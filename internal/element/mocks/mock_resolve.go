package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	elementMocks "github.com/satimoto/go-datastore/pkg/element/mocks"
	"github.com/satimoto/go-ocpi/internal/element"
	elementrestriction "github.com/satimoto/go-ocpi/internal/elementrestriction/mocks"
	pricecomponent "github.com/satimoto/go-ocpi/internal/pricecomponent/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *element.ElementResolver {
	return &element.ElementResolver{
		Repository:                 elementMocks.NewRepository(repositoryService),
		ElementRestrictionResolver: elementrestriction.NewResolver(repositoryService),
		PriceComponentResolver:     pricecomponent.NewResolver(repositoryService),
	}
}
