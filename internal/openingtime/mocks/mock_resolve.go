package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	openingtimeMocks "github.com/satimoto/go-datastore/pkg/openingtime/mocks"
	"github.com/satimoto/go-ocpi/internal/openingtime"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *openingtime.OpeningTimeResolver {
	return &openingtime.OpeningTimeResolver{
		Repository: openingtimeMocks.NewRepository(repositoryService),
	}
}
