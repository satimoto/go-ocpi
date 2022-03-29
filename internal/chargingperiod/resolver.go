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

func (r *ChargingPeriodResolver) ReplaceChargingPeriod(ctx context.Context, payload *ChargingPeriodPayload) *db.ChargingPeriod {
	if payload != nil {
		chargingPeriod, err := r.Repository.CreateChargingPeriod(ctx, *payload.StartDateTime)

		if err == nil {
			r.ReplaceChargingPeriodDimensions(ctx, &chargingPeriod.ID, *payload)

			return &chargingPeriod
		}
	}

	return nil
}

func (r *ChargingPeriodResolver) ReplaceChargingPeriodDimensions(ctx context.Context, chargingPeriodID *int64, payload ChargingPeriodPayload) {
	if chargingPeriodID != nil {
		r.Repository.DeleteChargingPeriodDimensions(ctx, *chargingPeriodID)

		for _, dimension := range payload.Dimensions {
			dimensionParams := NewCreateChargingPeriodDimensionParams(*chargingPeriodID, dimension)
			r.Repository.CreateChargingPeriodDimension(ctx, dimensionParams)
		}
	}
}
