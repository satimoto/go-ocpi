package element

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/elementrestriction"
	"github.com/satimoto/go-ocpi-api/internal/pricecomponent"
)

type ElementRepository interface {
	CreateElement(ctx context.Context, arg db.CreateElementParams) (db.Element, error)
	DeleteElements(ctx context.Context, tariffID int64) error
	ListElements(ctx context.Context, tariffID int64) ([]db.Element, error)
}

type ElementResolver struct {
	Repository                 ElementRepository
	ElementRestrictionResolver *elementrestriction.ElementRestrictionResolver
	PriceComponentResolver     *pricecomponent.PriceComponentResolver
}

func NewResolver(repositoryService *db.RepositoryService) *ElementResolver {
	repo := ElementRepository(repositoryService)
	
	return &ElementResolver{
		Repository:                 repo,
		ElementRestrictionResolver: elementrestriction.NewResolver(repositoryService),
		PriceComponentResolver:     pricecomponent.NewResolver(repositoryService),
	}
}
