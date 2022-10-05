package dto

import (
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
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
