package pricecomponent

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *PriceComponentResolver) CreatePriceComponents(ctx context.Context, elementID int64, dto []*PriceComponentDto) {
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
		_, err := r.Repository.CreatePriceComponent(ctx, priceComponentParams)

		if err != nil {
			util.LogOnError("OCPI138", "Error creating price component", err)
			log.Printf("OCPI138: Params=%#v", priceComponentParams)
		}
	}
}

func (r *PriceComponentResolver) createPriceComponentRounding(ctx context.Context, id *int64, dto *PriceComponentRoundingDto) {
	if dto != nil {
		priceComponentRoundingParams := NewCreatePriceComponentRoundingParams(dto)
		priceComponentRounding, err := r.Repository.CreatePriceComponentRounding(ctx, priceComponentRoundingParams)
		
		if err != nil {
			util.LogOnError("OCPI139", "Error creating price component rounding", err)
			log.Printf("OCPI139: Params=%#v", priceComponentRoundingParams)
			return
		}
		
		id = &priceComponentRounding.ID
	}
}
