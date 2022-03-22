package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/energymix"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *energymix.EnergyMixResolver {
	repo := energymix.EnergyMixRepository(repositoryService)

	return &energymix.EnergyMixResolver{
		Repository: repo,
	}
}
