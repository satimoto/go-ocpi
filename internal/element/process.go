package element

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *ElementResolver) ReplaceElements(ctx context.Context, tariffID int64, dto []*ElementDto) {
	if dto != nil {
		r.PriceComponentResolver.Repository.DeletePriceComponents(ctx, tariffID)
		r.Repository.DeleteElements(ctx, tariffID)
		r.ElementRestrictionResolver.Repository.DeleteElementRestrictions(ctx, tariffID)

		for _, elementDto := range dto {
			elementParams := NewCreateElementParams(elementDto)
			elementParams.TariffID = tariffID

			if elementDto.Restrictions != nil {
				restrictionID := util.NilInt64(nil)
				r.ElementRestrictionResolver.ReplaceElementRestriction(ctx, restrictionID, elementDto.Restrictions)
				elementParams.ElementRestrictionID = util.SqlNullInt64(restrictionID)
			}

			element, err := r.Repository.CreateElement(ctx, elementParams)

			if err != nil {
				util.LogOnError("OCPI091", "Error creating element", err)
				log.Printf("OCPI091: Params=%#v", elementParams)
				continue
			}
	
			r.PriceComponentResolver.CreatePriceComponents(ctx, element.ID, elementDto.PriceComponents)
		}
	}
}
