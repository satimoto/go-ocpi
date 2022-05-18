package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi-api/internal/tariffrestriction"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *tariffrestriction.TariffRestrictionResolver {
	repo := tariffrestriction.TariffRestrictionRepository(repositoryService)

	return &tariffrestriction.TariffRestrictionResolver{
		Repository: repo,
	}
}
