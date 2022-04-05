package element

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/elementrestriction"
	"github.com/satimoto/go-ocpi-api/internal/pricecomponent"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type ElementRepository interface {
	CreateElement(ctx context.Context, arg db.CreateElementParams) (db.Element, error)
	DeleteElements(ctx context.Context, tariffID int64) error
	ListElements(ctx context.Context, tariffID int64) ([]db.Element, error)
}

type ElementResolver struct {
	Repository ElementRepository
	*elementrestriction.ElementRestrictionResolver
	*pricecomponent.PriceComponentResolver
}

func NewResolver(repositoryService *db.RepositoryService) *ElementResolver {
	repo := ElementRepository(repositoryService)
	return &ElementResolver{
		Repository:                 repo,
		ElementRestrictionResolver: elementrestriction.NewResolver(repositoryService),
		PriceComponentResolver:     pricecomponent.NewResolver(repositoryService),
	}
}

func (r *ElementResolver) ReplaceElements(ctx context.Context, tariffID int64, dto []*ElementDto) {
	if dto != nil {
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
				r.PriceComponentResolver.ReplacePriceComponents(ctx, element.ID, elementDto.PriceComponents)
			}
		}
	}
}
