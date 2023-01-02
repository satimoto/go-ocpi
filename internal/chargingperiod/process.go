package chargingperiod

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *ChargingPeriodResolver) ReplaceChargingPeriod(ctx context.Context, chargingPeriodDto *coreDto.ChargingPeriodDto) *db.ChargingPeriod {
	if chargingPeriodDto != nil {
		chargingPeriod, err := r.Repository.CreateChargingPeriod(ctx, chargingPeriodDto.StartDateTime.Time())

		if err != nil {
			metrics.RecordError("OCPI036", "Error creating charging period", err)
			log.Printf("OCPI036: StartDateTime=%v", *chargingPeriodDto.StartDateTime)
			return nil
		}

		r.ReplaceChargingPeriodDimensions(ctx, &chargingPeriod.ID, *chargingPeriodDto)

		return &chargingPeriod
	}

	return nil
}

func (r *ChargingPeriodResolver) ReplaceChargingPeriodDimensions(ctx context.Context, chargingPeriodID *int64, chargingPeriodDto coreDto.ChargingPeriodDto) {
	if chargingPeriodID != nil {
		r.Repository.DeleteChargingPeriodDimensions(ctx, *chargingPeriodID)

		for _, dimension := range chargingPeriodDto.Dimensions {
			dimensionParams := NewCreateChargingPeriodDimensionParams(*chargingPeriodID, dimension)
			_, err := r.Repository.CreateChargingPeriodDimension(ctx, dimensionParams)

			if err != nil {
				metrics.RecordError("OCPI037", "Error creating charging period dimension", err)
				log.Printf("OCPI037: Params=%#v", dimensionParams)
			}
		}
	}
}
