package element

import (
	"context"

	"github.com/satimoto/go-datastore/util"
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

			if element, err := r.Repository.CreateElement(ctx, elementParams); err == nil {
				r.PriceComponentResolver.CreatePriceComponents(ctx, element.ID, elementDto.PriceComponents)
			}
		}
	}
}
