package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi-api/internal/image"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *image.ImageResolver {
	repo := image.ImageRepository(repositoryService)

	return &image.ImageResolver{
		Repository: repo,
	}
}
