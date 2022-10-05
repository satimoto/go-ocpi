package pricecomponent

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreatePriceComponentParams(priceComponentDto *coreDto.PriceComponentDto) db.CreatePriceComponentParams {
	return db.CreatePriceComponentParams{
		Type:                priceComponentDto.Type,
		Price:               priceComponentDto.Price,
		StepSize:            priceComponentDto.StepSize,
		ExactPriceComponent: util.SqlNullBool(priceComponentDto.ExactPriceComponent),
	}
}

func NewCreatePriceComponentRoundingParams(priceComponentRoundingDto *coreDto.PriceComponentRoundingDto) db.CreatePriceComponentRoundingParams {
	return db.CreatePriceComponentRoundingParams{
		Granularity: priceComponentRoundingDto.Granularity,
		Rule:        priceComponentRoundingDto.Rule,
	}
}
