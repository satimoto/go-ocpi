package openingtime

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *OpeningTimeResolver) ReplaceOpeningTime(ctx context.Context, id *int64, dto *OpeningTimeDto) {
	if dto != nil {
		if id == nil {
			openingTime, err := r.Repository.CreateOpeningTime(ctx, dto.Twentyfourseven)
			
			if err != nil {
				util.LogOnError("OCPI133", "Error creating opening time", err)
				log.Printf("OCPI133: Dto=%#v", dto)
				return
			}

			id = &openingTime.ID
		} else {
			openingTimeParams := NewUpdateOpeningTimeParams(*id, dto)
			_, err := r.Repository.UpdateOpeningTime(ctx, openingTimeParams)

			if err != nil {
				util.LogOnError("OCPI134", "Error updating opening time", err)
				log.Printf("OCPI134: Params=%#v", openingTimeParams)
				return
			}
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
			_, err := r.Repository.CreateExceptionalPeriod(ctx, exceptionalClosingParams)

			if err != nil {
				util.LogOnError("OCPI135", "Error creating exceptional closing period", err)
				log.Printf("OCPI135: Params=%#v", exceptionalClosingParams)
			}
		}
	}
}

func (r *OpeningTimeResolver) ReplaceExceptionalOpeningPeriods(ctx context.Context, openingTimeID *int64, dto OpeningTimeDto) {
	if openingTimeID != nil {
		r.Repository.DeleteExceptionalOpeningPeriods(ctx, *openingTimeID)

		for _, exceptionalOpening := range dto.ExceptionalOpenings {
			exceptionalOpeningParams := NewCreateExceptionalPeriodParams(*openingTimeID, db.PeriodTypeOPENING, exceptionalOpening)
			_, err := r.Repository.CreateExceptionalPeriod(ctx, exceptionalOpeningParams)

			if err != nil {
				util.LogOnError("OCPI136", "Error creating exceptional opening period", err)
				log.Printf("OCPI136: Params=%#v", exceptionalOpeningParams)
			}
		}
	}
}

func (r *OpeningTimeResolver) ReplaceRegularHours(ctx context.Context, openingTimeID *int64, dto OpeningTimeDto) {
	if openingTimeID != nil {
		r.Repository.DeleteRegularHours(ctx, *openingTimeID)

		for _, regularHour := range dto.RegularHours {
			regularHourParams := NewCreateRegularHourParams(*openingTimeID, regularHour)
			_, err := r.Repository.CreateRegularHour(ctx, regularHourParams)

			if err != nil {
				util.LogOnError("OCPI137", "Error creating exceptional opening period", err)
				log.Printf("OCPI137: Params=%#v", regularHourParams)
			}
		}
	}
}
