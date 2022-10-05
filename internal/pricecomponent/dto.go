package pricecomponent

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *PriceComponentResolver) CreatePriceComponentDto(ctx context.Context, priceComponent db.PriceComponent) *coreDto.PriceComponentDto {
	response := coreDto.NewPriceComponentDto(priceComponent)

	if priceComponent.PriceRoundingID.Valid {
		priceComponentRounding, err := r.Repository.GetPriceComponentRounding(ctx, priceComponent.PriceRoundingID.Int64)

		if err != nil {
			util.LogOnError("OCPI252", "Error retrieving price rounding", err)
			log.Printf("OCPI252: PriceRoundingID=%#v", priceComponent.PriceRoundingID)
		} else {	
			response.PriceRound = r.CreatePriceComponentRoundingDto(ctx, priceComponentRounding)
		}
	}

	if priceComponent.StepRoundingID.Valid {
		priceComponentRounding, err := r.Repository.GetPriceComponentRounding(ctx, priceComponent.StepRoundingID.Int64)

		if err != nil {
			util.LogOnError("OCPI253", "Error retrieving step rounding", err)
			log.Printf("OCPI253: StepRoundingID=%#v", priceComponent.StepRoundingID)
		} else {	
			response.StepRound = r.CreatePriceComponentRoundingDto(ctx, priceComponentRounding)
		}
	}

	return response
}

func (r *PriceComponentResolver) CreatePriceComponentListDto(ctx context.Context, priceComponents []db.PriceComponent) []*coreDto.PriceComponentDto {
	list := []*coreDto.PriceComponentDto{}
	
	for _, priceComponent := range priceComponents {
		list = append(list, r.CreatePriceComponentDto(ctx, priceComponent))
	}

	return list
}

func (r *PriceComponentResolver) CreatePriceComponentRoundingDto(ctx context.Context, priceComponentRounding db.PriceComponentRounding) *coreDto.PriceComponentRoundingDto {
	return coreDto.NewPriceComponentRoundingDto(priceComponentRounding)
}
