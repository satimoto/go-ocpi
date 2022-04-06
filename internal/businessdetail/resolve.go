package businessdetail

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/image"
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
