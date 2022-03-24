package restriction

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type RestrictionRepository interface {
	CreateRestriction(ctx context.Context, arg db.CreateRestrictionParams) (db.Restriction, error)
	DeleteRestrictions(ctx context.Context, tariffID int64) error
	GetRestriction(ctx context.Context, id int64) (db.Restriction, error)
	ListRestrictionWeekdays(ctx context.Context, restrictionID int64) ([]db.Weekday, error)
	ListWeekdays(ctx context.Context) ([]db.Weekday, error)
	SetRestrictionWeekday(ctx context.Context, arg db.SetRestrictionWeekdayParams) error
	UnsetRestrictionWeekdays(ctx context.Context, restrictionID int64) error
	UpdateRestriction(ctx context.Context, arg db.UpdateRestrictionParams) (db.Restriction, error)
}

type RestrictionResolver struct {
	Repository RestrictionRepository
}

func NewResolver(repositoryService *db.RepositoryService) *RestrictionResolver {
	repo := RestrictionRepository(repositoryService)
	return &RestrictionResolver{repo}
}

func (r *RestrictionResolver) ReplaceRestriction(ctx context.Context, id *int64, payload *RestrictionPayload) {
	if payload != nil {
		if id == nil {
			restrictionParams := NewCreateRestrictionParams(payload)

			if restriction, err := r.Repository.CreateRestriction(ctx, restrictionParams); err == nil {
				id = &restriction.ID
			}
		} else {
			restrictionParams := NewUpdateRestrictionParams(*id, payload)

			r.Repository.UpdateRestriction(ctx, restrictionParams)
		}

		if payload.DayOfWeek != nil {
			r.replaceWeekdays(ctx, *id, payload)
		}
	}
}

func (r *RestrictionResolver) replaceWeekdays(ctx context.Context, restrictionID int64, payload *RestrictionPayload) {
	r.Repository.UnsetRestrictionWeekdays(ctx, restrictionID)

	if weekdays, err := r.Repository.ListWeekdays(ctx); err == nil {
		filteredWeekdays := []*db.Weekday{}

		for _, weekday := range weekdays {
			if util.StringsContainString(payload.DayOfWeek, weekday.Text) {
				filteredWeekdays = append(filteredWeekdays, &weekday)
			}
		}

		for _, weekday := range filteredWeekdays {
			r.Repository.SetRestrictionWeekday(ctx, db.SetRestrictionWeekdayParams{
				RestrictionID: restrictionID,
				WeekdayID:     weekday.ID,
			})
		}
	}
}
