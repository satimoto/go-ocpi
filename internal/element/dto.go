package element

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/internal/elementrestriction"
	"github.com/satimoto/go-ocpi-api/internal/pricecomponent"
)

type ElementDto struct {
	PriceComponents []*pricecomponent.PriceComponentDto       `json:"price_components"`
	Restrictions    *elementrestriction.ElementRestrictionDto `json:"restrictions,omitempty"`
}

func (r *ElementDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewElementDto(element db.Element) *ElementDto {
	return &ElementDto{}
}

func (r *ElementResolver) CreateElementDto(ctx context.Context, element db.Element) *ElementDto {
	response := NewElementDto(element)

	if priceComponents, err := r.PriceComponentResolver.Repository.ListPriceComponents(ctx, element.ID); err == nil {
		response.PriceComponents = r.PriceComponentResolver.CreatePriceComponentListDto(ctx, priceComponents)
	}

	if element.ElementRestrictionID.Valid {
		if restriction, err := r.ElementRestrictionResolver.Repository.GetElementRestriction(ctx, element.ElementRestrictionID.Int64); err == nil {
			response.Restrictions = r.ElementRestrictionResolver.CreateElementRestrictionDto(ctx, restriction)
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
