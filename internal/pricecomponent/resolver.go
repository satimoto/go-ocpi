package pricecomponent

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
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

func (r *PriceComponentResolver) CreatePriceComponents(ctx context.Context, elementID int64, dto []*PriceComponentDto) {
	if dto != nil {
		for _, priceComponentDto := range dto {
			priceRoundingID := util.NilInt64(nil)
			stepRoundingID := util.NilInt64(nil)

			if priceComponentDto.PriceRound != nil {
				r.createPriceComponentRounding(ctx, priceRoundingID, priceComponentDto.PriceRound)
			}

			if priceComponentDto.StepRound != nil {
				r.createPriceComponentRounding(ctx, stepRoundingID, priceComponentDto.StepRound)
			}

			priceComponentParams := NewCreatePriceComponentParams(priceComponentDto)
			priceComponentParams.ElementID = elementID
			priceComponentParams.PriceRoundingID = util.SqlNullInt64(priceRoundingID)
			priceComponentParams.StepRoundingID = util.SqlNullInt64(stepRoundingID)

			r.Repository.CreatePriceComponent(ctx, priceComponentParams)
		}
	}
}

func (r *PriceComponentResolver) createPriceComponentRounding(ctx context.Context, id *int64, dto *PriceComponentRoundingDto) {
	if dto != nil {
		priceComponentRoundingParams := NewCreatePriceComponentRoundingParams(dto)

		if priceComponentRounding, err := r.Repository.CreatePriceComponentRounding(ctx, priceComponentRoundingParams); err == nil {
			id = &priceComponentRounding.ID
		}
	}
}
