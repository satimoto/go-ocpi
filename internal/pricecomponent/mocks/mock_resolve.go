package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	pricecomponentMocks "github.com/satimoto/go-datastore/pkg/pricecomponent/mocks"
	"github.com/satimoto/go-ocpi-api/internal/pricecomponent"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *pricecomponent.PriceComponentResolver {
	return &pricecomponent.PriceComponentResolver{
		Repository: pricecomponentMocks.NewRepository(repositoryService),
	}
}
