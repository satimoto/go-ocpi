package evse

import (
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/util"
)

func NewCreateEvseParams(locationID int64, dto *EvseDto) db.CreateEvseParams {
	return db.CreateEvseParams{
		Uid:               *dto.Uid,
		EvseID:            dbUtil.SqlNullString(util.ReplaceAllString(dto.EvseID, "", "[^a-zA-Z0-9]+")),
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
