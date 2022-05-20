package energymix

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/energymix"
)

type EnergyMixResolver struct {
	Repository energymix.EnergyMixRepository
}

func NewResolver(repositoryService *db.RepositoryService) *EnergyMixResolver {
	return &EnergyMixResolver{
		Repository: energymix.NewRepository(repositoryService),
	}
}
