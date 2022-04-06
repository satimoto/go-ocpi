package openingtime

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

func (r *OpeningTimeResolver) ReplaceOpeningTime(ctx context.Context, id *int64, dto *OpeningTimeDto) {
	if dto != nil {
		if id == nil {
			if openingTime, err := r.Repository.CreateOpeningTime(ctx, dto.Twentyfourseven); err == nil {
				id = &openingTime.ID
			}
		} else {
			openingTimeParams := NewUpdateOpeningTimeParams(*id, dto)

			r.Repository.UpdateOpeningTime(ctx, openingTimeParams)
		}

		r.ReplaceRegularHours(ctx, id, *dto)
		r.ReplaceExceptionalClosingPeriods(ctx, id, *dto)
		r.ReplaceExceptionalOpeningPeriods(ctx, id, *dto)
	}
}

func (r *OpeningTimeResolver) ReplaceExceptionalClosingPeriods(ctx context.Context, openingTimeID *int64, dto OpeningTimeDto) {
	if openingTimeID != nil {
		r.Repository.DeleteExceptionalClosingPeriods(ctx, *openingTimeID)

		for _, exceptionalClosing := range dto.ExceptionalClosings {
			exceptionalClosingParams := NewCreateExceptionalPeriodParams(*openingTimeID, db.PeriodTypeCLOSING, exceptionalClosing)
			r.Repository.CreateExceptionalPeriod(ctx, exceptionalClosingParams)
		}
	}
}

func (r *OpeningTimeResolver) ReplaceExceptionalOpeningPeriods(ctx context.Context, openingTimeID *int64, dto OpeningTimeDto) {
	if openingTimeID != nil {
		r.Repository.DeleteExceptionalOpeningPeriods(ctx, *openingTimeID)

		for _, exceptionalOpening := range dto.ExceptionalOpenings {
			exceptionalOpeningParams := NewCreateExceptionalPeriodParams(*openingTimeID, db.PeriodTypeOPENING, exceptionalOpening)
			r.Repository.CreateExceptionalPeriod(ctx, exceptionalOpeningParams)
		}
	}
}

func (r *OpeningTimeResolver) ReplaceRegularHours(ctx context.Context, openingTimeID *int64, dto OpeningTimeDto) {
	if openingTimeID != nil {
		r.Repository.DeleteRegularHours(ctx, *openingTimeID)

		for _, regularHour := range dto.RegularHours {
			regularHourParams := NewCreateRegularHourParams(*openingTimeID, regularHour)
			r.Repository.CreateRegularHour(ctx, regularHourParams)
		}
	}
}
