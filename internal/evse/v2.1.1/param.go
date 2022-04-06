package evse

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewCreateEvseParams(locationID int64, dto *EvseDto) db.CreateEvseParams {
	return db.CreateEvseParams{
		Uid:               *dto.Uid,
		EvseID:            util.SqlNullString(dto.EvseID),
		LocationID:        locationID,
		Status:            *dto.Status,
		FloorLevel:        util.SqlNullString(dto.FloorLevel),
		PhysicalReference: util.SqlNullString(dto.PhysicalReference),
		LastUpdated:       *dto.LastUpdated,
	}
}

func NewUpdateEvseByUidParams(evse db.Evse) db.UpdateEvseByUidParams {
	return db.UpdateEvseByUidParams{
		Uid:               evse.Uid,
		EvseID:            evse.EvseID,
		Status:            evse.Status,
		FloorLevel:        evse.FloorLevel,
		Geom:              evse.Geom,
		GeoLocationID:     evse.GeoLocationID,
		PhysicalReference: evse.PhysicalReference,
		LastUpdated:       evse.LastUpdated,
	}
}

func NewCreateStatusScheduleParams(evseID int64, dto *StatusScheduleDto) db.CreateStatusScheduleParams {
	return db.CreateStatusScheduleParams{
		EvseID:      evseID,
		PeriodBegin: *dto.PeriodBegin,
		PeriodEnd:   util.SqlNullTime(dto.PeriodEnd),
	}
}
