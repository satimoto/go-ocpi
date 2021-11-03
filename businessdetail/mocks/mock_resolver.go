package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/businessdetail"
	image "github.com/satimoto/go-ocpi-api/image/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *businessdetail.BusinessDetailResolver {
	repo := businessdetail.BusinessDetailRepository(repositoryService)

	return &businessdetail.BusinessDetailResolver{
		Repository: repo,
		ImageResolver: image.NewResolver(repositoryService),
	}
}
