package elementrestriction

import (
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
)

func NewCreateElementRestrictionParams(dto *ElementRestrictionDto) db.CreateElementRestrictionParams {
	return db.CreateElementRestrictionParams{
		StartTime:   util.SqlNullString(dto.StartTime),
		EndTime:     util.SqlNullString(dto.EndTime),
		StartDate:   util.SqlNullString(dto.StartDate),
		EndDate:     util.SqlNullString(dto.EndDate),
		MinKwh:      util.SqlNullFloat64(dto.MinKwh),
		MaxKwh:      util.SqlNullFloat64(dto.MaxKwh),
		MinPower:    util.SqlNullFloat64(dto.MinPower),
		MaxPower:    util.SqlNullFloat64(dto.MaxPower),
		MinDuration: util.SqlNullInt32(dto.MinDuration),
		MaxDuration: util.SqlNullInt32(dto.MaxDuration),
	}
}

func NewUpdateElementRestrictionParams(id int64, dto *ElementRestrictionDto) db.UpdateElementRestrictionParams {
	return db.UpdateElementRestrictionParams{
		ID:          id,
		StartTime:   util.SqlNullString(dto.StartTime),
		EndTime:     util.SqlNullString(dto.EndTime),
		StartDate:   util.SqlNullString(dto.StartDate),
		EndDate:     util.SqlNullString(dto.EndDate),
		MinKwh:      util.SqlNullFloat64(dto.MinKwh),
		MaxKwh:      util.SqlNullFloat64(dto.MaxKwh),
		MinPower:    util.SqlNullFloat64(dto.MinPower),
		MaxPower:    util.SqlNullFloat64(dto.MaxPower),
		MinDuration: util.SqlNullInt32(dto.MinDuration),
		MaxDuration: util.SqlNullInt32(dto.MaxDuration),
	}
}
