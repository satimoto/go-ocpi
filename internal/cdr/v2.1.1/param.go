package cdr

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func NewCreateCdrParams(cdrDto *dto.CdrDto) db.CreateCdrParams {
	return db.CreateCdrParams{
		Uid:              *cdrDto.ID,
		AuthorizationID:  util.SqlNullString(cdrDto.AuthorizationID),
		StartDateTime:    *cdrDto.StartDateTime,
		StopDateTime:     util.SqlNullTime(cdrDto.StopDateTime),
		AuthID:           *cdrDto.AuthID,
		AuthMethod:       *cdrDto.AuthMethod,
		MeterID:          util.SqlNullString(cdrDto.MeterID),
		Currency:         *cdrDto.Currency,
		TotalCost:        *cdrDto.TotalCost,
		TotalEnergy:      *cdrDto.TotalEnergy,
		TotalTime:        *cdrDto.TotalTime,
		TotalParkingTime: util.SqlNullFloat64(cdrDto.TotalParkingTime),
		Remark:           util.SqlNullString(cdrDto.Remark),
		LastUpdated:      *cdrDto.LastUpdated,
	}
}
