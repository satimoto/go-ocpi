package openingtime

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type OpeningTimeRepository interface {
	CreateExceptionalPeriod(ctx context.Context, arg db.CreateExceptionalPeriodParams) (db.ExceptionalPeriod, error)
	CreateOpeningTime(ctx context.Context, twentyfourseven bool) (db.OpeningTime, error)
	CreateRegularHour(ctx context.Context, arg db.CreateRegularHourParams) (db.RegularHour, error)
	DeleteExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) error
	DeleteExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) error
	DeleteOpeningTime(ctx context.Context, id int64) error
	DeleteRegularHours(ctx context.Context, openingTimeID int64) error
	GetOpeningTime(ctx context.Context, id int64) (db.OpeningTime, error)
	ListExceptionalOpeningPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error)
	ListExceptionalClosingPeriods(ctx context.Context, openingTimeID int64) ([]db.ExceptionalPeriod, error)
	ListRegularHours(ctx context.Context, openingTimeID int64) ([]db.RegularHour, error)
	UpdateOpeningTime(ctx context.Context, arg db.UpdateOpeningTimeParams) (db.OpeningTime, error)
}

type OpeningTimeResolver struct {
	Repository OpeningTimeRepository
}

func NewResolver(repositoryService *db.RepositoryService) *OpeningTimeResolver {
	repo := OpeningTimeRepository(repositoryService)
	return &OpeningTimeResolver{repo}
}

func (r *OpeningTimeResolver) ReplaceOpeningTime(ctx context.Context, id *int64, payload *OpeningTimePayload) {
	if payload != nil {
		if id == nil {
			if openingTime, err := r.Repository.CreateOpeningTime(ctx, payload.Twentyfourseven); err == nil {
				id = &openingTime.ID
			}
		} else {
			openingTimeParams := NewUpdateOpeningTimeParams(*id, payload)

			r.Repository.UpdateOpeningTime(ctx, openingTimeParams)	
		}

		r.ReplaceRegularHours(ctx, id, *payload)
		r.ReplaceExceptionalClosingPeriods(ctx, id, *payload)
		r.ReplaceExceptionalOpeningPeriods(ctx, id, *payload)
	}
}

func (r *OpeningTimeResolver) ReplaceExceptionalClosingPeriods(ctx context.Context, openingTimeID *int64, payload OpeningTimePayload) {
	if openingTimeID != nil {
		r.Repository.DeleteExceptionalClosingPeriods(ctx, *openingTimeID)

		for _, exceptionalClosing := range payload.ExceptionalClosings {
			exceptionalClosingParams := NewCreateExceptionalPeriodParams(*openingTimeID, db.PeriodTypeCLOSING, exceptionalClosing)
			r.Repository.CreateExceptionalPeriod(ctx, exceptionalClosingParams)
		}
	}
}

func (r *OpeningTimeResolver) ReplaceExceptionalOpeningPeriods(ctx context.Context, openingTimeID *int64, payload OpeningTimePayload) {
	if openingTimeID != nil {
		r.Repository.DeleteExceptionalOpeningPeriods(ctx, *openingTimeID)

		for _, exceptionalOpening := range payload.ExceptionalOpenings {
			exceptionalOpeningParams := NewCreateExceptionalPeriodParams(*openingTimeID, db.PeriodTypeOPENING, exceptionalOpening)
			r.Repository.CreateExceptionalPeriod(ctx, exceptionalOpeningParams)
		}
	}
}

func (r *OpeningTimeResolver) ReplaceRegularHours(ctx context.Context, openingTimeID *int64, payload OpeningTimePayload) {
	if openingTimeID != nil {
		r.Repository.DeleteRegularHours(ctx, *openingTimeID)

		for _, regularHour := range payload.RegularHours {
			regularHourParams := NewCreateRegularHourParams(*openingTimeID, regularHour)
			r.Repository.CreateRegularHour(ctx, regularHourParams)
		}
	}
}
