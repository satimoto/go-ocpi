package pricecomponent

import (
	"context"
	"database/sql"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *PriceComponentResolver) CreatePriceComponents(ctx context.Context, elementID int64, tariff db.Tariff, priceComponentsDto []*coreDto.PriceComponentDto) {
	for _, priceComponentDto := range priceComponentsDto {
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

func (r *PriceComponentResolver) createPriceComponentRounding(ctx context.Context, id *sql.NullInt64, priceComponentRoundingDto *coreDto.PriceComponentRoundingDto) {
	if priceComponentRoundingDto != nil {
		priceComponentRoundingParams := NewCreatePriceComponentRoundingParams(priceComponentRoundingDto)
		priceComponentRounding, err := r.Repository.CreatePriceComponentRounding(ctx, priceComponentRoundingParams)

		if err != nil {
			util.LogOnError("OCPI139", "Error creating price component rounding", err)
			log.Printf("OCPI139: Params=%#v", priceComponentRoundingParams)
			return
		}

		id.Scan(priceComponentRounding.ID)
	}
}
