package cdr

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateCdrParams(dto *CdrDto) db.CreateCdrParams {
	return db.CreateCdrParams{
		Uid:              *dto.ID,
		AuthorizationID:  util.SqlNullString(dto.AuthorizationID),
		StartDateTime:    *dto.StartDateTime,
		StopDateTime:     util.SqlNullTime(dto.StopDateTime),
		AuthID:           *dto.AuthID,
		AuthMethod:       *dto.AuthMethod,
		MeterID:          util.SqlNullString(dto.MeterID),
		Currency:         *dto.Currency,
		TotalCost:        *dto.TotalCost,
		TotalEnergy:      *dto.TotalEnergy,
		TotalTime:        *dto.TotalTime,
		TotalParkingTime: util.SqlNullFloat64(dto.TotalParkingTime),
		Remark:           util.SqlNullString(dto.Remark),
		LastUpdated:      *dto.LastUpdated,
	}
}
