package evse

import (
	"github.com/satimoto/go-datastore/pkg/db"
	dbUtil "github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
)

func NewCreateEvseParams(locationID int64, evseDto *dto.EvseDto) db.CreateEvseParams {
	return db.CreateEvseParams{
		Uid:               *evseDto.Uid,
		EvseID:            dbUtil.SqlNullString(evseDto.EvseID),
		Identifier:        dbUtil.SqlNullString(GetEvseIdentifier(evseDto)),
		LocationID:        locationID,
		Status:            *evseDto.Status,
		IsRemoteCapable:   false,
		IsRfidCapable:     false,
		FloorLevel:        dbUtil.SqlNullString(evseDto.FloorLevel),
		PhysicalReference: dbUtil.SqlNullString(evseDto.PhysicalReference),
		LastUpdated:       *evseDto.LastUpdated,
	}
}

func NewCreateStatusScheduleParams(evseID int64, dto *coreDto.StatusScheduleDto) db.CreateStatusScheduleParams {
	return db.CreateStatusScheduleParams{
		EvseID:      evseID,
		PeriodBegin: *dto.PeriodBegin,
		PeriodEnd:   dbUtil.SqlNullTime(dto.PeriodEnd),
	}
}
