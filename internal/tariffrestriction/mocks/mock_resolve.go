package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	tariffrestrictionMocks "github.com/satimoto/go-datastore/pkg/tariffrestriction/mocks"
	"github.com/satimoto/go-ocpi-api/internal/tariffrestriction"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *tariffrestriction.TariffRestrictionResolver {
	return &tariffrestriction.TariffRestrictionResolver{
		Repository: tariffrestrictionMocks.NewRepository(repositoryService),
	}
}
