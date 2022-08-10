package openingtime

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/openingtime"
)

type OpeningTimeResolver struct {
	Repository openingtime.OpeningTimeRepository
}

func NewResolver(repositoryService *db.RepositoryService) *OpeningTimeResolver {
	return &OpeningTimeResolver{
		Repository: openingtime.NewRepository(repositoryService),
	}
}
