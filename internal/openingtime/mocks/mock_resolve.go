package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi-api/internal/openingtime"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *openingtime.OpeningTimeResolver {
	repo := openingtime.OpeningTimeRepository(repositoryService)

	return &openingtime.OpeningTimeResolver{
		Repository: repo,
	}
}
