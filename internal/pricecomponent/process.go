package pricecomponent

import (
	"context"
	"database/sql"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *PriceComponentResolver) CreatePriceComponents(ctx context.Context, elementID int64, tariff db.Tariff, dto []*PriceComponentDto) {
	for _, priceComponentDto := range dto {
		priceRoundingID := util.SqlNullInt64(nil)
		stepRoundingID := util.SqlNullInt64(nil)

		if priceComponentDto.PriceRound != nil {
			r.createPriceComponentRounding(ctx, &priceRoundingID, priceComponentDto.PriceRound)
		}

		if priceComponentDto.StepRound != nil {
			r.createPriceComponentRounding(ctx, &stepRoundingID, priceComponentDto.StepRound)
		}

		priceComponentParams := NewCreatePriceComponentParams(priceComponentDto)
		priceComponentParams.Currency = tariff.Currency
		priceComponentParams.ElementID = elementID
		priceComponentParams.PriceRoundingID = priceRoundingID
		priceComponentParams.StepRoundingID = stepRoundingID
		_, err := r.Repository.CreatePriceComponent(ctx, priceComponentParams)

		if err != nil {
			util.LogOnError("OCPI138", "Error creating price component", err)
			log.Printf("OCPI138: Params=%#v", priceComponentParams)
		}
	}
}

func (r *PriceComponentResolver) createPriceComponentRounding(ctx context.Context, id *sql.NullInt64, dto *PriceComponentRoundingDto) {
	if dto != nil {
		priceComponentRoundingParams := NewCreatePriceComponentRoundingParams(dto)
		priceComponentRounding, err := r.Repository.CreatePriceComponentRounding(ctx, priceComponentRoundingParams)
		
		if err != nil {
			util.LogOnError("OCPI139", "Error creating price component rounding", err)
			log.Printf("OCPI139: Params=%#v", priceComponentRoundingParams)
			return
		}
		
		id.Scan(priceComponentRounding.ID)
	}
}
