package elementrestriction

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/elementrestriction"
)


type ElementRestrictionResolver struct {
	Repository elementrestriction.ElementRestrictionRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ElementRestrictionResolver {
	return &ElementRestrictionResolver{
		Repository: elementrestriction.NewRepository(repositoryService),
	}
}
