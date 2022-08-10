package mocks

import (
	chargingperiodMocks "github.com/satimoto/go-datastore/pkg/chargingperiod/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi/internal/chargingperiod"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *chargingperiod.ChargingPeriodResolver {
	return &chargingperiod.ChargingPeriodResolver{
		Repository: chargingperiodMocks.NewRepository(repositoryService),
	}
}
