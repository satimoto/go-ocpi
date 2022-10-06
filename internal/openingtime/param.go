package openingtime

import (
	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func NewCreateExceptionalPeriodParams(id int64, periodType db.PeriodType, exceptionalPeriodDto *coreDto.ExceptionalPeriodDto) db.CreateExceptionalPeriodParams {
	return db.CreateExceptionalPeriodParams{
		OpeningTimeID: id,
		PeriodType:    periodType,
		PeriodBegin:   exceptionalPeriodDto.PeriodBegin.Time(),
		PeriodEnd:     exceptionalPeriodDto.PeriodEnd.Time(),
	}
}

func NewUpdateOpeningTimeParams(id int64, openingTimeDto *coreDto.OpeningTimeDto) db.UpdateOpeningTimeParams {
	return db.UpdateOpeningTimeParams{
		ID:              id,
		Twentyfourseven: openingTimeDto.Twentyfourseven,
	}
}

func NewCreateRegularHourParams(id int64, regularHourDto *coreDto.RegularHourDto) db.CreateRegularHourParams {
	return db.CreateRegularHourParams{
		OpeningTimeID: id,
		Weekday:       regularHourDto.Weekday,
		PeriodBegin:   regularHourDto.PeriodBegin,
		PeriodEnd:     regularHourDto.PeriodEnd,
	}
}
