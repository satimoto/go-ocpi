package elementrestriction

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateElementRestrictionParams(elementRestrictionDto *coreDto.ElementRestrictionDto) db.CreateElementRestrictionParams {
	return db.CreateElementRestrictionParams{
		StartTime:   util.SqlNullString(elementRestrictionDto.StartTime),
		EndTime:     util.SqlNullString(elementRestrictionDto.EndTime),
		StartDate:   util.SqlNullString(elementRestrictionDto.StartDate),
		EndDate:     util.SqlNullString(elementRestrictionDto.EndDate),
		MinKwh:      util.SqlNullFloat64(elementRestrictionDto.MinKwh),
		MaxKwh:      util.SqlNullFloat64(elementRestrictionDto.MaxKwh),
		MinPower:    util.SqlNullFloat64(elementRestrictionDto.MinPower),
		MaxPower:    util.SqlNullFloat64(elementRestrictionDto.MaxPower),
		MinDuration: util.SqlNullInt32(elementRestrictionDto.MinDuration),
		MaxDuration: util.SqlNullInt32(elementRestrictionDto.MaxDuration),
	}
}

func NewUpdateElementRestrictionParams(id int64, elementRestrictionDto *coreDto.ElementRestrictionDto) db.UpdateElementRestrictionParams {
	return db.UpdateElementRestrictionParams{
		ID:          id,
		StartTime:   util.SqlNullString(elementRestrictionDto.StartTime),
		EndTime:     util.SqlNullString(elementRestrictionDto.EndTime),
		StartDate:   util.SqlNullString(elementRestrictionDto.StartDate),
		EndDate:     util.SqlNullString(elementRestrictionDto.EndDate),
		MinKwh:      util.SqlNullFloat64(elementRestrictionDto.MinKwh),
		MaxKwh:      util.SqlNullFloat64(elementRestrictionDto.MaxKwh),
		MinPower:    util.SqlNullFloat64(elementRestrictionDto.MinPower),
		MaxPower:    util.SqlNullFloat64(elementRestrictionDto.MaxPower),
		MinDuration: util.SqlNullInt32(elementRestrictionDto.MinDuration),
		MaxDuration: util.SqlNullInt32(elementRestrictionDto.MaxDuration),
	}
}
