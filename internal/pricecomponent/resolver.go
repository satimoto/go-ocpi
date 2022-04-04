package pricecomponent

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type PriceComponentRepository interface {
	CreatePriceComponent(ctx context.Context, arg db.CreatePriceComponentParams) (db.PriceComponent, error)
	ListPriceComponents(ctx context.Context, elementID int64) ([]db.PriceComponent, error)
}

type PriceComponentResolver struct {
	Repository PriceComponentRepository
}

func NewResolver(repositoryService *db.RepositoryService) *PriceComponentResolver {
	repo := PriceComponentRepository(repositoryService)
	return &PriceComponentResolver{repo}
}

func (r *PriceComponentResolver) ReplacePriceComponents(ctx context.Context, elementID int64, dto []*PriceComponentDto) {
	if dto != nil {
		for _, priceComponentDto := range dto {
			priceComponentParams := NewCreatePriceComponentParams(priceComponentDto)
			priceComponentParams.ElementID = elementID

			r.Repository.CreatePriceComponent(ctx, priceComponentParams)
		}
	}
}
