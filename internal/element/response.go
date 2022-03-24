package element

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/pricecomponent"
	"github.com/satimoto/go-ocpi-api/internal/restriction"
)

type ElementPayload struct {
	PriceComponents []*pricecomponent.PriceComponentPayload `json:"price_components"`
	Restrictions    *restriction.RestrictionPayload         `json:"restrictions,omitempty"`
}

func (r *ElementPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewElementPayload(element db.Element) *ElementPayload {
	return &ElementPayload{}
}

func NewCreateElementParams(payload *ElementPayload) db.CreateElementParams {
	return db.CreateElementParams{}
}

func (r *ElementResolver) CreateElementPayload(ctx context.Context, element db.Element) *ElementPayload {
	response := NewElementPayload(element)

	if priceComponents, err := r.PriceComponentResolver.Repository.ListPriceComponents(ctx, element.ID); err == nil {
		response.PriceComponents = r.PriceComponentResolver.CreatePriceComponentListPayload(ctx, priceComponents)
	}

	if element.RestrictionID.Valid {
		if restriction, err := r.RestrictionResolver.Repository.GetRestriction(ctx, element.RestrictionID.Int64); err == nil {
			response.Restrictions = r.RestrictionResolver.CreateRestrictionPayload(ctx, restriction)
		}
	}

	return response
}

func (r *ElementResolver) CreateElementListPayload(ctx context.Context, elements []db.Element) []*ElementPayload {
	list := []*ElementPayload{}
	for _, element := range elements {
		list = append(list, r.CreateElementPayload(ctx, element))
	}
	return list
}
