package chargingperiod

import (
	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateChargingPeriodDimensionParams(id int64, dto *coreDto.ChargingPeriodDimensionDto) db.CreateChargingPeriodDimensionParams {
	return db.CreateChargingPeriodDimensionParams{
		ChargingPeriodID: id,
		Type:             dto.Type,
		Volume:           dto.Volume,
	}
}
