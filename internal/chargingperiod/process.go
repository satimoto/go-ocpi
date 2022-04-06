package chargingperiod

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

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
