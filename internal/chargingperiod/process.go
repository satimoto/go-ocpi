package chargingperiod

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *ChargingPeriodResolver) ReplaceChargingPeriod(ctx context.Context, dto *ChargingPeriodDto) *db.ChargingPeriod {
	if dto != nil {
		chargingPeriod, err := r.Repository.CreateChargingPeriod(ctx, *dto.StartDateTime)

		if err != nil {
			util.LogOnError("OCPI036", "Error creating charging period", err)
			log.Printf("OCPI036: StartDateTime=%v", *dto.StartDateTime)
			return nil
		}

		r.ReplaceChargingPeriodDimensions(ctx, &chargingPeriod.ID, *dto)

		return &chargingPeriod
	}

	return nil
}

func (r *ChargingPeriodResolver) ReplaceChargingPeriodDimensions(ctx context.Context, chargingPeriodID *int64, dto ChargingPeriodDto) {
	if chargingPeriodID != nil {
		r.Repository.DeleteChargingPeriodDimensions(ctx, *chargingPeriodID)

		for _, dimension := range dto.Dimensions {
			dimensionParams := NewCreateChargingPeriodDimensionParams(*chargingPeriodID, dimension)
			_, err := r.Repository.CreateChargingPeriodDimension(ctx, dimensionParams)

			if err != nil {
				util.LogOnError("OCPI037", "Error creating charging period dimension", err)
				log.Printf("OCPI037: Params=%#v", dimensionParams)
			}
		}
	}
}
