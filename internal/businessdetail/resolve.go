package businessdetail

import (
	"github.com/satimoto/go-datastore/pkg/businessdetail"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/image"
)

type BusinessDetailResolver struct {
	Repository    businessdetail.BusinessDetailRepository
	ImageResolver *image.ImageResolver
}

func NewResolver(repositoryService *db.RepositoryService) *BusinessDetailResolver {
	return &BusinessDetailResolver{
		Repository:    businessdetail.NewRepository(repositoryService),
		ImageResolver: image.NewResolver(repositoryService),
	}
}
