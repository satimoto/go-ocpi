package element

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/metric"
)

func (r *ElementResolver) CreateElementDto(ctx context.Context, element db.Element) *coreDto.ElementDto {
	response := coreDto.NewElementDto(element)

	priceComponents, err := r.PriceComponentResolver.Repository.ListPriceComponents(ctx, element.ID)

	if err != nil {
		metrics.RecordError("OCPI226", "Error listing price components", err)
		log.Printf("OCPI226: ElementID=%v", element.ID)
	} else {
		response.PriceComponents = r.PriceComponentResolver.CreatePriceComponentListDto(ctx, priceComponents)
	}

	if element.ElementRestrictionID.Valid {
		restriction, err := r.ElementRestrictionResolver.Repository.GetElementRestriction(ctx, element.ElementRestrictionID.Int64)

		if err != nil {
			metrics.RecordError("OCPI227", "Error retrieving element restriction", err)
			log.Printf("OCPI227: ElementRestrictionID=%#v", element.ElementRestrictionID)
		} else {
			response.Restrictions = r.ElementRestrictionResolver.CreateElementRestrictionDto(ctx, restriction)
		}
	}

	return response
}

func (r *ElementResolver) CreateElementListDto(ctx context.Context, elements []db.Element) []*coreDto.ElementDto {
	list := []*coreDto.ElementDto{}

	for _, element := range elements {
		list = append(list, r.CreateElementDto(ctx, element))
	}

	return list
}
