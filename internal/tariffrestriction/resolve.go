package tariffrestriction

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/tariffrestriction"
)

type TariffRestrictionResolver struct {
	Repository tariffrestriction.TariffRestrictionRepository
}

func NewResolver(repositoryService *db.RepositoryService) *TariffRestrictionResolver {
	return &TariffRestrictionResolver{
		Repository: tariffrestriction.NewRepository(repositoryService),
	}
}
