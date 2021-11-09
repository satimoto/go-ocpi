package version

import (
	"github.com/satimoto/go-datastore/db"
)

type VersionDetailRepository interface {
}

type VersionDetailResolver struct {
	Repository VersionDetailRepository
}

func NewResolver(repositoryService *db.RepositoryService) *VersionDetailResolver {
	repo := VersionDetailRepository(repositoryService)
	return &VersionDetailResolver{
		Repository: repo,
	}
}
