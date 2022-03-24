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

func (r *PriceComponentResolver) ReplacePriceComponents(ctx context.Context, elementID int64, payload []*PriceComponentPayload) {
	if payload != nil {
		for _, priceComponentPayload := range payload {
			priceComponentParams := NewCreatePriceComponentParams(priceComponentPayload)
			priceComponentParams.ElementID = elementID
	
			r.Repository.CreatePriceComponent(ctx, priceComponentParams)
		}
	}
}
