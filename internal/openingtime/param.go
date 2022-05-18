package openingtime

import "github.com/satimoto/go-datastore/pkg/db"

func NewCreateExceptionalPeriodParams(id int64, periodType db.PeriodType, dto *ExceptionalPeriodDto) db.CreateExceptionalPeriodParams {
	return db.CreateExceptionalPeriodParams{
		OpeningTimeID: id,
		PeriodType:    periodType,
		PeriodBegin:   *dto.PeriodBegin,
		PeriodEnd:     *dto.PeriodEnd,
	}
}

func NewUpdateOpeningTimeParams(id int64, dto *OpeningTimeDto) db.UpdateOpeningTimeParams {
	return db.UpdateOpeningTimeParams{
		ID:              id,
		Twentyfourseven: dto.Twentyfourseven,
	}
}

func NewCreateRegularHourParams(id int64, dto *RegularHourDto) db.CreateRegularHourParams {
	return db.CreateRegularHourParams{
		OpeningTimeID: id,
		Weekday:       dto.Weekday,
		PeriodBegin:   dto.PeriodBegin,
		PeriodEnd:     dto.PeriodEnd,
	}
}
