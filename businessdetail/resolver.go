package businessdetail

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/image"
	"github.com/satimoto/go-ocpi-api/util"
)

type BusinessDetailRepository interface {
	CreateBusinessDetail(ctx context.Context, arg db.CreateBusinessDetailParams) (db.BusinessDetail, error)
	DeleteBusinessDetail(ctx context.Context, id int64) error
	DeleteBusinessDetailLogo(ctx context.Context, id int64) error
	GetBusinessDetail(ctx context.Context, id int64) (db.BusinessDetail, error)
	UpdateBusinessDetail(ctx context.Context, arg db.UpdateBusinessDetailParams) (db.BusinessDetail, error)
}

type BusinessDetailResolver struct {
	Repository BusinessDetailRepository
	*image.ImageResolver
}

func NewResolver(repositoryService *db.RepositoryService) *BusinessDetailResolver {
	repo := BusinessDetailRepository(repositoryService)
	return &BusinessDetailResolver{
		Repository:    repo,
		ImageResolver: image.NewResolver(repositoryService),
	}
}

func NewCreateBusinessDetailParams(payload *BusinessDetailPayload) db.CreateBusinessDetailParams {
	return db.CreateBusinessDetailParams{
		Name:    payload.Name,
		Website: util.SqlNullString(payload.Website),
	}
}

func NewUpdateBusinessDetailParams(id int64, payload *BusinessDetailPayload) db.UpdateBusinessDetailParams {
	return db.UpdateBusinessDetailParams{
		ID:      id,
		Name:    payload.Name,
		Website: util.SqlNullString(payload.Website),
	}
}

func (r *BusinessDetailResolver) ReplaceBusinessDetail(ctx context.Context, id *int64, payload *BusinessDetailPayload) {
	if payload != nil {
		logoID := r.ImageResolver.CreateImage(ctx, payload.Logo)

		if id == nil {
			businessDetailParams := NewCreateBusinessDetailParams(payload)
			businessDetailParams.LogoID = util.SqlNullInt64(logoID)

			if businessDetail, err := r.Repository.CreateBusinessDetail(ctx, businessDetailParams); err == nil {
				id = &businessDetail.ID
			}
		} else {
			businessDetailParams := NewUpdateBusinessDetailParams(*id, payload)
			businessDetailParams.LogoID = util.SqlNullInt64(logoID)

			r.Repository.UpdateBusinessDetail(ctx, businessDetailParams)
		}
	}
}
