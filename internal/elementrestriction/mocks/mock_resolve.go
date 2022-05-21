package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	elementrestrictionMocks "github.com/satimoto/go-datastore/pkg/elementrestriction/mocks"
	"github.com/satimoto/go-ocpi-api/internal/elementrestriction"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *elementrestriction.ElementRestrictionResolver {
	return &elementrestriction.ElementRestrictionResolver{
		Repository: elementrestrictionMocks.NewRepository(repositoryService),
	}
}
