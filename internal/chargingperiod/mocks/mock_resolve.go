package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *chargingperiod.ChargingPeriodResolver {
	repo := chargingperiod.ChargingPeriodRepository(repositoryService)

	return &chargingperiod.ChargingPeriodResolver{
		Repository: repo,
	}
}
