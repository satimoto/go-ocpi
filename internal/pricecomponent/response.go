package pricecomponent

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
)

type PriceComponentPayload struct {
	Type     db.TariffDimension `json:"type"`
	Price    float64            `json:"price"`
	StepSize int32              `json:"step_size"`
}

func (r *PriceComponentPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewPriceComponentPayload(priceComponent db.PriceComponent) *PriceComponentPayload {
	return &PriceComponentPayload{
		Type:     priceComponent.Type,
		Price:    priceComponent.Price,
		StepSize: priceComponent.StepSize,
	}
}

func NewCreatePriceComponentParams(payload *PriceComponentPayload) db.CreatePriceComponentParams {
	return db.CreatePriceComponentParams{
		Type:     payload.Type,
		Price:    payload.Price,
		StepSize: payload.StepSize,
	}
}

func (r *PriceComponentResolver) CreatePriceComponentPayload(ctx context.Context, priceComponent db.PriceComponent) *PriceComponentPayload {
	return NewPriceComponentPayload(priceComponent)
}

func (r *PriceComponentResolver) CreatePriceComponentListPayload(ctx context.Context, priceComponents []db.PriceComponent) []*PriceComponentPayload {
	list := []*PriceComponentPayload{}
	for _, priceComponent := range priceComponents {
		list = append(list, r.CreatePriceComponentPayload(ctx, priceComponent))
	}
	return list
}
