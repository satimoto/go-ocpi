package pricecomponent

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
)

type PriceComponentDto struct {
	Type     db.TariffDimension `json:"type"`
	Price    float64            `json:"price"`
	StepSize int32              `json:"step_size"`
}

func (r *PriceComponentDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewPriceComponentDto(priceComponent db.PriceComponent) *PriceComponentDto {
	return &PriceComponentDto{
		Type:     priceComponent.Type,
		Price:    priceComponent.Price,
		StepSize: priceComponent.StepSize,
	}
}

func NewCreatePriceComponentParams(dto *PriceComponentDto) db.CreatePriceComponentParams {
	return db.CreatePriceComponentParams{
		Type:     dto.Type,
		Price:    dto.Price,
		StepSize: dto.StepSize,
	}
}

func (r *PriceComponentResolver) CreatePriceComponentDto(ctx context.Context, priceComponent db.PriceComponent) *PriceComponentDto {
	return NewPriceComponentDto(priceComponent)
}

func (r *PriceComponentResolver) CreatePriceComponentListDto(ctx context.Context, priceComponents []db.PriceComponent) []*PriceComponentDto {
	list := []*PriceComponentDto{}
	for _, priceComponent := range priceComponents {
		list = append(list, r.CreatePriceComponentDto(ctx, priceComponent))
	}
	return list
}
