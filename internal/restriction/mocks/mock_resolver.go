package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/restriction"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *restriction.RestrictionResolver {
	repo := restriction.RestrictionRepository(repositoryService)

	return &restriction.RestrictionResolver{
		Repository: repo,
	}
}
