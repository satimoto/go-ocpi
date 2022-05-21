package displaytext

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/displaytext"
)

type DisplayTextResolver struct {
	Repository displaytext.DisplayTextRepository
}

func NewResolver(repositoryService *db.RepositoryService) *DisplayTextResolver {
	return &DisplayTextResolver{
		Repository: displaytext.NewRepository(repositoryService),
	}
}
