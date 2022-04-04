package element

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/pricecomponent"
	"github.com/satimoto/go-ocpi-api/internal/restriction"
)

type ElementDto struct {
	PriceComponents []*pricecomponent.PriceComponentDto `json:"price_components"`
	Restrictions    *restriction.RestrictionDto         `json:"restrictions,omitempty"`
}

func (r *ElementDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewElementDto(element db.Element) *ElementDto {
	return &ElementDto{}
}

func NewCreateElementParams(dto *ElementDto) db.CreateElementParams {
	return db.CreateElementParams{}
}

func (r *ElementResolver) CreateElementDto(ctx context.Context, element db.Element) *ElementDto {
	response := NewElementDto(element)

	if priceComponents, err := r.PriceComponentResolver.Repository.ListPriceComponents(ctx, element.ID); err == nil {
		response.PriceComponents = r.PriceComponentResolver.CreatePriceComponentListDto(ctx, priceComponents)
	}

	if element.RestrictionID.Valid {
		if restriction, err := r.RestrictionResolver.Repository.GetRestriction(ctx, element.RestrictionID.Int64); err == nil {
			response.Restrictions = r.RestrictionResolver.CreateRestrictionDto(ctx, restriction)
		}
	}

	return response
}

func (r *ElementResolver) CreateElementListDto(ctx context.Context, elements []db.Element) []*ElementDto {
	list := []*ElementDto{}
	for _, element := range elements {
		list = append(list, r.CreateElementDto(ctx, element))
	}
	return list
}
