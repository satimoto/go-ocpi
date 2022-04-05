package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/elementrestriction"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *elementrestriction.ElementRestrictionResolver {
	repo := elementrestriction.ElementRestrictionRepository(repositoryService)

	return &elementrestriction.ElementRestrictionResolver{
		Repository: repo,
	}
}
