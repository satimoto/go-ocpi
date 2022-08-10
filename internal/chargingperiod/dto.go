package chargingperiod

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

type ChargingPeriodDto struct {
	StartDateTime *time.Time                    `json:"start_date_time"`
	Dimensions    []*ChargingPeriodDimensionDto `json:"dimensions"`
}

func (r *ChargingPeriodDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewChargingPeriodDto(chargingPeriod db.ChargingPeriod) *ChargingPeriodDto {
	return &ChargingPeriodDto{
		StartDateTime: &chargingPeriod.StartDateTime,
	}
}

func (r *ChargingPeriodResolver) CreateChargingPeriodDto(ctx context.Context, chargingPeriod db.ChargingPeriod) *ChargingPeriodDto {
	response := NewChargingPeriodDto(chargingPeriod)

	chargingPeriodDimensions, err := r.Repository.ListChargingPeriodDimensions(ctx, chargingPeriod.ID)
	
	if err != nil {
		util.LogOnError("OCPI223", "Error listing charging period dimensions", err)
		log.Printf("OCPI223: CalibrationID=%v", chargingPeriod.ID)
		return response
	}
	
	response.Dimensions = r.CreateChargingPeriodDimensionListDto(ctx, chargingPeriodDimensions)

	return response
}

func (r *ChargingPeriodResolver) CreateChargingPeriodListDto(ctx context.Context, chargingPeriods []db.ChargingPeriod) []*ChargingPeriodDto {
	list := []*ChargingPeriodDto{}

	for _, chargingPeriod := range chargingPeriods {
		list = append(list, r.CreateChargingPeriodDto(ctx, chargingPeriod))
	}
	
	return list
}

type ChargingPeriodDimensionDto struct {
	Type   db.ChargingPeriodDimensionType `json:"type"`
	Volume float64                        `json:"volume"`
}

func (r *ChargingPeriodDimensionDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewChargingPeriodDimensionDto(chargingPeriodDimension db.ChargingPeriodDimension) *ChargingPeriodDimensionDto {
	return &ChargingPeriodDimensionDto{
		Type:   chargingPeriodDimension.Type,
		Volume: chargingPeriodDimension.Volume,
	}
}

func (r *ChargingPeriodResolver) CreateChargingPeriodDimensionDto(ctx context.Context, chargingPeriodDimension db.ChargingPeriodDimension) *ChargingPeriodDimensionDto {
	return NewChargingPeriodDimensionDto(chargingPeriodDimension)
}

func (r *ChargingPeriodResolver) CreateChargingPeriodDimensionListDto(ctx context.Context, chargingPeriodDimensions []db.ChargingPeriodDimension) []*ChargingPeriodDimensionDto {
	list := []*ChargingPeriodDimensionDto{}

	for _, chargingPeriodDimension := range chargingPeriodDimensions {
		list = append(list, r.CreateChargingPeriodDimensionDto(ctx, chargingPeriodDimension))
	}

	return list
}