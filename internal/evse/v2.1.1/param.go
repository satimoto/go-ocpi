package evse

import (
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
)

func NewCreateEvseParams(locationID int64, dto *EvseDto) db.CreateEvseParams {
	return db.CreateEvseParams{
		Uid:               *dto.Uid,
		EvseID:            dbUtil.SqlNullString(dto.EvseID),
		Identifier:        dbUtil.SqlNullString(GetEvseIdentifier(dto)),
		LocationID:        locationID,
		Status:            *dto.Status,
		IsRemoteCapable:   false,
		IsRfidCapable:     false,
		FloorLevel:        dbUtil.SqlNullString(dto.FloorLevel),
		PhysicalReference: dbUtil.SqlNullString(dto.PhysicalReference),
		LastUpdated:       *dto.LastUpdated,
	}
}

func NewCreateStatusScheduleParams(evseID int64, dto *StatusScheduleDto) db.CreateStatusScheduleParams {
	return db.CreateStatusScheduleParams{
		EvseID:      evseID,
		PeriodBegin: *dto.PeriodBegin,
		PeriodEnd:   dbUtil.SqlNullTime(dto.PeriodEnd),
	}
}
