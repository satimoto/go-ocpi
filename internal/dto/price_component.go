package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
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
