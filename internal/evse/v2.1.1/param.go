package evse

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
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

func NewUpdateEvseByUidParams(evse db.Evse) db.UpdateEvseByUidParams {
	return db.UpdateEvseByUidParams{
		Uid:               evse.Uid,
		EvseID:            evse.EvseID,
		Status:            evse.Status,
		FloorLevel:        evse.FloorLevel,
		Geom:              evse.Geom,
		GeoLocationID:     evse.GeoLocationID,
		IsRemoteCapable:   evse.IsRemoteCapable,
		IsRfidCapable:     evse.IsRfidCapable,
		PhysicalReference: evse.PhysicalReference,
		LastUpdated:       evse.LastUpdated,
	}
}

func NewUpdateLocationAvailabilityParams(locationID int64) db.UpdateLocationAvailabilityParams {
	return db.UpdateLocationAvailabilityParams{
		ID:              locationID,
		AvailableEvses:  0,
		TotalEvses:      0,
		IsRemoteCapable: false,
		IsRfidCapable:   false,
	}
}
