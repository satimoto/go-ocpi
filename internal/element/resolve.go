package element

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/element"
	"github.com/satimoto/go-ocpi/internal/elementrestriction"
	"github.com/satimoto/go-ocpi/internal/pricecomponent"
)

type ElementResolver struct {
	Repository                 element.ElementRepository
	ElementRestrictionResolver *elementrestriction.ElementRestrictionResolver
	PriceComponentResolver     *pricecomponent.PriceComponentResolver
}

func NewResolver(repositoryService *db.RepositoryService) *ElementResolver {
	return &ElementResolver{
		Repository:                 element.NewRepository(repositoryService),
		ElementRestrictionResolver: elementrestriction.NewResolver(repositoryService),
		PriceComponentResolver:     pricecomponent.NewResolver(repositoryService),
	}
}
