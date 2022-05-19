package image

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/image"
)

type ImageResolver struct {
	Repository image.ImageRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ImageResolver {
	return &ImageResolver{
		Repository: image.NewRepository(repositoryService),
	}
}
