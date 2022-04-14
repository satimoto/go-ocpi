package pricecomponent

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func NewCreatePriceComponentParams(dto *PriceComponentDto) db.CreatePriceComponentParams {
	return db.CreatePriceComponentParams{
		Type:                dto.Type,
		Price:               dto.Price,
		StepSize:            dto.StepSize,
		ExactPriceComponent: util.SqlNullBool(dto.ExactPriceComponent),
	}
}

func NewCreatePriceComponentRoundingParams(dto *PriceComponentRoundingDto) db.CreatePriceComponentRoundingParams {
	return db.CreatePriceComponentRoundingParams{
		Granularity: dto.Granularity,
		Rule:        dto.Rule,
	}
}
