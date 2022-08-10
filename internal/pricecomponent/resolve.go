package pricecomponent

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/pricecomponent"
)

type PriceComponentResolver struct {
	Repository pricecomponent.PriceComponentRepository
}

func NewResolver(repositoryService *db.RepositoryService) *PriceComponentResolver {
	return &PriceComponentResolver{
		Repository: pricecomponent.NewRepository(repositoryService),
	}
}
