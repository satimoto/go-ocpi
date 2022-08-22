package element

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *ElementResolver) ReplaceElements(ctx context.Context, tariff db.Tariff, dto []*ElementDto) {
	if dto != nil {
		r.PriceComponentResolver.Repository.DeletePriceComponents(ctx, tariff.ID)
		r.Repository.DeleteElements(ctx, tariff.ID)
		r.ElementRestrictionResolver.Repository.DeleteElementRestrictions(ctx, tariff.ID)

		for _, elementDto := range dto {
			elementParams := NewCreateElementParams(elementDto)
			elementParams.TariffID = tariff.ID

			if elementDto.Restrictions != nil {
				restrictionID := util.SqlNullInt64(nil)
				r.ElementRestrictionResolver.ReplaceElementRestriction(ctx, &restrictionID, elementDto.Restrictions)
				elementParams.ElementRestrictionID = restrictionID
			}

			element, err := r.Repository.CreateElement(ctx, elementParams)

			if err != nil {
				util.LogOnError("OCPI091", "Error creating element", err)
				log.Printf("OCPI091: Params=%#v", elementParams)
				continue
			}
	
			r.PriceComponentResolver.CreatePriceComponents(ctx, element.ID, tariff, elementDto.PriceComponents)
		}
	}
}
