package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	energymixMocks "github.com/satimoto/go-datastore/pkg/energymix/mocks"
	"github.com/satimoto/go-ocpi/internal/energymix"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *energymix.EnergyMixResolver {
	return &energymix.EnergyMixResolver{
		Repository: energymixMocks.NewRepository(repositoryService),
	}
}
