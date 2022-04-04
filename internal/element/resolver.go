package element

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/pricecomponent"
	"github.com/satimoto/go-ocpi-api/internal/restriction"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type ElementRepository interface {
	CreateElement(ctx context.Context, arg db.CreateElementParams) (db.Element, error)
	DeleteElements(ctx context.Context, tariffID int64) error
	ListElements(ctx context.Context, tariffID int64) ([]db.Element, error)
}

type ElementResolver struct {
	Repository ElementRepository
	*pricecomponent.PriceComponentResolver
	*restriction.RestrictionResolver
}

func NewResolver(repositoryService *db.RepositoryService) *ElementResolver {
	repo := ElementRepository(repositoryService)
	return &ElementResolver{
		Repository:             repo,
		PriceComponentResolver: pricecomponent.NewResolver(repositoryService),
		RestrictionResolver:    restriction.NewResolver(repositoryService),
	}
}

func (r *ElementResolver) ReplaceElements(ctx context.Context, tariffID int64, dto []*ElementDto) {
	if dto != nil {
		r.Repository.DeleteElements(ctx, tariffID)
		r.RestrictionResolver.Repository.DeleteRestrictions(ctx, tariffID)

		for _, elementDto := range dto {
			elementParams := NewCreateElementParams(elementDto)
			elementParams.TariffID = tariffID

			if elementDto.Restrictions != nil {
				restrictionID := util.NilInt64(nil)
				r.RestrictionResolver.ReplaceRestriction(ctx, restrictionID, elementDto.Restrictions)
				elementParams.RestrictionID = util.SqlNullInt64(restrictionID)
			}

			if element, err := r.Repository.CreateElement(ctx, elementParams); err == nil {
				r.PriceComponentResolver.ReplacePriceComponents(ctx, element.ID, elementDto.PriceComponents)
			}
		}
	}
}
