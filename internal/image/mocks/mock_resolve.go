package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	imageMocks "github.com/satimoto/go-datastore/pkg/image/mocks"
	"github.com/satimoto/go-ocpi/internal/image"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *image.ImageResolver {
	return &image.ImageResolver{
		Repository: imageMocks.NewRepository(repositoryService),
	}
}
