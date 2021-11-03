package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/image"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *image.ImageResolver {
	repo := image.ImageRepository(repositoryService)

	return &image.ImageResolver{
		Repository: repo,
	}
}