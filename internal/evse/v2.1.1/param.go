package evse

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateEvseParams(locationID int64, dto *EvseDto) db.CreateEvseParams {
	return db.CreateEvseParams{
		Uid:               *dto.Uid,
		EvseID:            util.SqlNullString(dto.EvseID),
		LocationID:        locationID,
		Status:            *dto.Status,
		IsRemoteCapable:   false,
		IsRfidCapable:     false,
		FloorLevel:        util.SqlNullString(dto.FloorLevel),
		PhysicalReference: util.SqlNullString(dto.PhysicalReference),
		LastUpdated:       *dto.LastUpdated,
	}
}

func NewCreateStatusScheduleParams(evseID int64, dto *StatusScheduleDto) db.CreateStatusScheduleParams {
	return db.CreateStatusScheduleParams{
		EvseID:      evseID,
		PeriodBegin: *dto.PeriodBegin,
		PeriodEnd:   util.SqlNullTime(dto.PeriodEnd),
	}
}
