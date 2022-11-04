package openingtime

import (
	"context"
	"database/sql"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *OpeningTimeResolver) ReplaceOpeningTime(ctx context.Context, id *sql.NullInt64, openingTimeDto *coreDto.OpeningTimeDto) {
	if openingTimeDto != nil {
		if id.Valid {
			openingTimeParams := NewUpdateOpeningTimeParams(id.Int64, openingTimeDto)
			_, err := r.Repository.UpdateOpeningTime(ctx, openingTimeParams)

			if err != nil {
				metrics.RecordError("OCPI134", "Error updating opening time", err)
				log.Printf("OCPI134: Params=%#v", openingTimeParams)
				return
			}
		} else {
			openingTime, err := r.Repository.CreateOpeningTime(ctx, openingTimeDto.Twentyfourseven)

			if err != nil {
				metrics.RecordError("OCPI133", "Error creating opening time", err)
				log.Printf("OCPI133: Dto=%#v", openingTimeDto)
				return
			}

			id.Scan(openingTime.ID)
		}

		r.ReplaceRegularHours(ctx, id.Int64, *openingTimeDto)
		r.ReplaceExceptionalClosingPeriods(ctx, id.Int64, *openingTimeDto)
		r.ReplaceExceptionalOpeningPeriods(ctx, id.Int64, *openingTimeDto)
	}
}

func (r *OpeningTimeResolver) ReplaceExceptionalClosingPeriods(ctx context.Context, openingTimeID int64, openingTimeDto coreDto.OpeningTimeDto) {
	r.Repository.DeleteExceptionalClosingPeriods(ctx, openingTimeID)

	for _, exceptionalClosing := range openingTimeDto.ExceptionalClosings {
		exceptionalClosingParams := NewCreateExceptionalPeriodParams(openingTimeID, db.PeriodTypeCLOSING, exceptionalClosing)
		_, err := r.Repository.CreateExceptionalPeriod(ctx, exceptionalClosingParams)

		if err != nil {
			metrics.RecordError("OCPI135", "Error creating exceptional closing period", err)
			log.Printf("OCPI135: Params=%#v", exceptionalClosingParams)
		}
	}
}

func (r *OpeningTimeResolver) ReplaceExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64, openingTimeDto coreDto.OpeningTimeDto) {
	r.Repository.DeleteExceptionalOpeningPeriods(ctx, openingTimeID)

	for _, exceptionalOpening := range openingTimeDto.ExceptionalOpenings {
		exceptionalOpeningParams := NewCreateExceptionalPeriodParams(openingTimeID, db.PeriodTypeOPENING, exceptionalOpening)
		_, err := r.Repository.CreateExceptionalPeriod(ctx, exceptionalOpeningParams)

		if err != nil {
			metrics.RecordError("OCPI136", "Error creating exceptional opening period", err)
			log.Printf("OCPI136: Params=%#v", exceptionalOpeningParams)
		}
	}
}

func (r *OpeningTimeResolver) ReplaceRegularHours(ctx context.Context, openingTimeID int64, openingTimeDto coreDto.OpeningTimeDto) {
	r.Repository.DeleteRegularHours(ctx, openingTimeID)

	for _, regularHour := range openingTimeDto.RegularHours {
		regularHourParams := NewCreateRegularHourParams(openingTimeID, regularHour)
		_, err := r.Repository.CreateRegularHour(ctx, regularHourParams)

		if err != nil {
			metrics.RecordError("OCPI137", "Error creating exceptional opening period", err)
			log.Printf("OCPI137: Params=%#v", regularHourParams)
		}
	}
}
