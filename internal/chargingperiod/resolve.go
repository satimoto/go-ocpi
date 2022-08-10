package chargingperiod

import (
	"github.com/satimoto/go-datastore/pkg/chargingperiod"
	"github.com/satimoto/go-datastore/pkg/db"
)

type ChargingPeriodResolver struct {
	Repository chargingperiod.ChargingPeriodRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ChargingPeriodResolver {
	return &ChargingPeriodResolver{
		Repository: chargingperiod.NewRepository(repositoryService),
	}
}
