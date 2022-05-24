package element

import (
	"context"
	"log"
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
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

	priceComponents, err := r.PriceComponentResolver.Repository.ListPriceComponents(ctx, element.ID)

	if err != nil {
		util.LogOnError("OCPI226", "Error listing price components", err)
		log.Printf("OCPI226: ElementID=%v", element.ID)
	} else {
		response.PriceComponents = r.PriceComponentResolver.CreatePriceComponentListDto(ctx, priceComponents)
	}

	if element.ElementRestrictionID.Valid {
		restriction, err := r.ElementRestrictionResolver.Repository.GetElementRestriction(ctx, element.ElementRestrictionID.Int64)

		if err != nil {
			util.LogOnError("OCPI227", "Error retrieving element restriction", err)
			log.Printf("OCPI227: ElementRestrictionID=%#v", element.ElementRestrictionID)
		} else {
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
