package mocks

import (
	businessdetailMocks "github.com/satimoto/go-datastore/pkg/businessdetail/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi/internal/businessdetail"
	image "github.com/satimoto/go-ocpi/internal/image/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *businessdetail.BusinessDetailResolver {
	return &businessdetail.BusinessDetailResolver{
		Repository:    businessdetailMocks.NewRepository(repositoryService),
		ImageResolver: image.NewResolver(repositoryService),
	}
}
