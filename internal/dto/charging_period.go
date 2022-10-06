package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type ChargingPeriodDto struct {
	StartDateTime *ocpitype.Time                `json:"start_date_time"`
	Dimensions    []*ChargingPeriodDimensionDto `json:"dimensions"`
}

func (r *ChargingPeriodDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewChargingPeriodDto(chargingPeriod db.ChargingPeriod) *ChargingPeriodDto {
	return &ChargingPeriodDto{
		StartDateTime: ocpitype.NilOcpiTime(&chargingPeriod.StartDateTime),
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
