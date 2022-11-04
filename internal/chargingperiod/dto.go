package chargingperiod

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/metric"
)

func (r *ChargingPeriodResolver) CreateChargingPeriodDto(ctx context.Context, chargingPeriod db.ChargingPeriod) *coreDto.ChargingPeriodDto {
	response := coreDto.NewChargingPeriodDto(chargingPeriod)

	chargingPeriodDimensions, err := r.Repository.ListChargingPeriodDimensions(ctx, chargingPeriod.ID)

	if err != nil {
		metrics.RecordError("OCPI223", "Error listing charging period dimensions", err)
		log.Printf("OCPI223: CalibrationID=%v", chargingPeriod.ID)
		return response
	}

	response.Dimensions = r.CreateChargingPeriodDimensionListDto(ctx, chargingPeriodDimensions)

	return response
}

func (r *ChargingPeriodResolver) CreateChargingPeriodListDto(ctx context.Context, chargingPeriods []db.ChargingPeriod) []*coreDto.ChargingPeriodDto {
	list := []*coreDto.ChargingPeriodDto{}

	for _, chargingPeriod := range chargingPeriods {
		list = append(list, r.CreateChargingPeriodDto(ctx, chargingPeriod))
	}

	return list
}

func (r *ChargingPeriodResolver) CreateChargingPeriodDimensionDto(ctx context.Context, chargingPeriodDimension db.ChargingPeriodDimension) *coreDto.ChargingPeriodDimensionDto {
	return coreDto.NewChargingPeriodDimensionDto(chargingPeriodDimension)
}

func (r *ChargingPeriodResolver) CreateChargingPeriodDimensionListDto(ctx context.Context, chargingPeriodDimensions []db.ChargingPeriodDimension) []*coreDto.ChargingPeriodDimensionDto {
	list := []*coreDto.ChargingPeriodDimensionDto{}

	for _, chargingPeriodDimension := range chargingPeriodDimensions {
		list = append(list, r.CreateChargingPeriodDimensionDto(ctx, chargingPeriodDimension))
	}

	return list
}
