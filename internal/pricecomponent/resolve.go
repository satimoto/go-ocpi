package pricecomponent

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

type PriceComponentRepository interface {
	CreatePriceComponent(ctx context.Context, arg db.CreatePriceComponentParams) (db.PriceComponent, error)
	CreatePriceComponentRounding(ctx context.Context, arg db.CreatePriceComponentRoundingParams) (db.PriceComponentRounding, error)
	DeletePriceComponents(ctx context.Context, tariffID int64) error
	DeletePriceComponentRoundings(ctx context.Context, tariffID int64) error
	GetPriceComponentRounding(ctx context.Context, id int64) (db.PriceComponentRounding, error)
	ListPriceComponents(ctx context.Context, elementID int64) ([]db.PriceComponent, error)
}

type PriceComponentResolver struct {
	Repository PriceComponentRepository
}

func NewResolver(repositoryService *db.RepositoryService) *PriceComponentResolver {
	repo := PriceComponentRepository(repositoryService)

	return &PriceComponentResolver{repo}
}
