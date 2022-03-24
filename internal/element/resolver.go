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

func (r *ElementResolver) ReplaceElements(ctx context.Context, tariffID int64, payload []*ElementPayload) {
	if payload != nil {
		r.Repository.DeleteElements(ctx, tariffID)
		r.RestrictionResolver.Repository.DeleteRestrictions(ctx, tariffID)

		for _, elementPayload := range payload {
			elementParams := NewCreateElementParams(elementPayload)
			elementParams.TariffID = tariffID

			if elementPayload.Restrictions != nil {
				restrictionID := util.NilInt64(nil)
				r.RestrictionResolver.ReplaceRestriction(ctx, restrictionID, elementPayload.Restrictions)
				elementParams.RestrictionID = util.SqlNullInt64(restrictionID)
			}

			if element, err := r.Repository.CreateElement(ctx, elementParams); err == nil {
				r.PriceComponentResolver.ReplacePriceComponents(ctx, element.ID, elementPayload.PriceComponents)
			}
		}
	}
}

