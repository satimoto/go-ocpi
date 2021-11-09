package version

import (
	"github.com/satimoto/go-datastore/db"
)

type VersionRepository interface {
}

type VersionResolver struct {
	Repository VersionRepository
}

func NewResolver(repositoryService *db.RepositoryService) *VersionResolver {
	repo := VersionRepository(repositoryService)
	return &VersionResolver{
		Repository: repo,
	}
}
