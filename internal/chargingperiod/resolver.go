package chargingperiod

import (
	"context"
	"time"

	"github.com/satimoto/go-datastore/db"
)

type ChargingPeriodRepository interface {
	CreateChargingPeriod(ctx context.Context, startDateTime time.Time) (db.ChargingPeriod, error)
	CreateChargingPeriodDimension(ctx context.Context, arg db.CreateChargingPeriodDimensionParams) (db.ChargingPeriodDimension, error)
	DeleteChargingPeriodDimensions(ctx context.Context, chargingPeriodID int64) error
	ListChargingPeriodDimensions(ctx context.Context, chargingPeriodID int64) ([]db.ChargingPeriodDimension, error)
}

type ChargingPeriodResolver struct {
	Repository ChargingPeriodRepository
}

func NewResolver(repositoryService *db.RepositoryService) *ChargingPeriodResolver {
	repo := ChargingPeriodRepository(repositoryService)
	return &ChargingPeriodResolver{repo}
}

func (r *ChargingPeriodResolver) ReplaceChargingPeriod(ctx context.Context, dto *ChargingPeriodDto) *db.ChargingPeriod {
	if dto != nil {
		chargingPeriod, err := r.Repository.CreateChargingPeriod(ctx, *dto.StartDateTime)

		if err == nil {
			r.ReplaceChargingPeriodDimensions(ctx, &chargingPeriod.ID, *dto)

			return &chargingPeriod
		}
	}

	return nil
}

func (r *ChargingPeriodResolver) ReplaceChargingPeriodDimensions(ctx context.Context, chargingPeriodID *int64, dto ChargingPeriodDto) {
	if chargingPeriodID != nil {
		r.Repository.DeleteChargingPeriodDimensions(ctx, *chargingPeriodID)

		for _, dimension := range dto.Dimensions {
			dimensionParams := NewCreateChargingPeriodDimensionParams(*chargingPeriodID, dimension)
			r.Repository.CreateChargingPeriodDimension(ctx, dimensionParams)
		}
	}
}
