package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi-api/internal/pricecomponent"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *pricecomponent.PriceComponentResolver {
	repo := pricecomponent.PriceComponentRepository(repositoryService)

	return &pricecomponent.PriceComponentResolver{
		Repository: repo,
	}
}
