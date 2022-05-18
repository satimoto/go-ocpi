package chargingperiod

import "github.com/satimoto/go-datastore/pkg/db"

func NewCreateChargingPeriodDimensionParams(id int64, dto *ChargingPeriodDimensionDto) db.CreateChargingPeriodDimensionParams {
	return db.CreateChargingPeriodDimensionParams{
		ChargingPeriodID: id,
		Type:             dto.Type,
		Volume:           dto.Volume,
	}
}
