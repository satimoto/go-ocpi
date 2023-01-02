package party

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/party"
)

type PartyResolver struct {
	Repository party.PartyRepository
}

func NewResolver(repositoryService *db.RepositoryService) *PartyResolver {
	return &PartyResolver{
		Repository: party.NewRepository(repositoryService),
	}
}
