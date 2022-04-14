package pricecomponent

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

type PriceComponentDto struct {
	Type                db.TariffDimension         `json:"type"`
	Price               float64                    `json:"price"`
	StepSize            int32                      `json:"step_size"`
	PriceRound          *PriceComponentRoundingDto `json:"price_round,omitempty"`
	StepRound           *PriceComponentRoundingDto `json:"step_round,omitempty"`
	ExactPriceComponent *bool                      `json:"exact_price_component,omitempty"`
}

func (r *PriceComponentDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewPriceComponentDto(priceComponent db.PriceComponent) *PriceComponentDto {
	return &PriceComponentDto{
		Type:                priceComponent.Type,
		Price:               priceComponent.Price,
		StepSize:            priceComponent.StepSize,
		ExactPriceComponent: util.NilBool(priceComponent.ExactPriceComponent),
	}
}

func (r *PriceComponentResolver) CreatePriceComponentDto(ctx context.Context, priceComponent db.PriceComponent) *PriceComponentDto {
	response := NewPriceComponentDto(priceComponent)

	if priceComponent.PriceRoundingID.Valid {
		if priceComponentRounding, err := r.Repository.GetPriceComponentRounding(ctx, priceComponent.PriceRoundingID.Int64); err == nil {
			response.PriceRound = r.CreatePriceComponentRoundingDto(ctx, priceComponentRounding)
		}
	}

	if priceComponent.StepRoundingID.Valid {
		if priceComponentRounding, err := r.Repository.GetPriceComponentRounding(ctx, priceComponent.StepRoundingID.Int64); err == nil {
			response.StepRound = r.CreatePriceComponentRoundingDto(ctx, priceComponentRounding)
		}
	}

	return response
}

func (r *PriceComponentResolver) CreatePriceComponentListDto(ctx context.Context, priceComponents []db.PriceComponent) []*PriceComponentDto {
	list := []*PriceComponentDto{}
	for _, priceComponent := range priceComponents {
		list = append(list, r.CreatePriceComponentDto(ctx, priceComponent))
	}
	return list
}

type PriceComponentRoundingDto struct {
	Granularity db.RoundingGranularity `json:"round_granularity"`
	Rule        db.RoundingRule        `json:"round_rule"`
}

func (r *PriceComponentRoundingDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewPriceComponentRoundingDto(priceComponentRounding db.PriceComponentRounding) *PriceComponentRoundingDto {
	return &PriceComponentRoundingDto{
		Granularity: priceComponentRounding.Granularity,
		Rule:        priceComponentRounding.Rule,
	}
}

func (r *PriceComponentResolver) CreatePriceComponentRoundingDto(ctx context.Context, priceComponentRounding db.PriceComponentRounding) *PriceComponentRoundingDto {
	return NewPriceComponentRoundingDto(priceComponentRounding)
}
