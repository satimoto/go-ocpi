package tariffrestriction

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type TariffRestrictionRepository interface {
	CreateTariffRestriction(ctx context.Context, arg db.CreateTariffRestrictionParams) (db.TariffRestriction, error)
	DeleteTariffRestriction(ctx context.Context, tariffID int64) error
	GetTariffRestriction(ctx context.Context, id int64) (db.TariffRestriction, error)
	ListTariffRestrictionWeekdays(ctx context.Context, tariffRestrictionID int64) ([]db.Weekday, error)
	ListWeekdays(ctx context.Context) ([]db.Weekday, error)
	SetTariffRestrictionWeekday(ctx context.Context, arg db.SetTariffRestrictionWeekdayParams) error
	UnsetTariffRestrictionWeekdays(ctx context.Context, tariffRestrictionID int64) error
	UpdateTariffRestriction(ctx context.Context, arg db.UpdateTariffRestrictionParams) (db.TariffRestriction, error)
}

type TariffRestrictionResolver struct {
	Repository TariffRestrictionRepository
}

func NewResolver(repositoryService *db.RepositoryService) *TariffRestrictionResolver {
	repo := TariffRestrictionRepository(repositoryService)
	return &TariffRestrictionResolver{repo}
}

func (r *TariffRestrictionResolver) ReplaceTariffRestriction(ctx context.Context, id *int64, dto *TariffRestrictionDto) {
	if dto != nil {
		if id == nil {
			tariffRestrictionParams := NewCreateTariffRestrictionParams(dto)

			if tariffRestriction, err := r.Repository.CreateTariffRestriction(ctx, tariffRestrictionParams); err == nil {
				id = &tariffRestriction.ID
			}
		} else {
			tariffRestrictionParams := NewUpdateTariffRestrictionParams(*id, dto)

			r.Repository.UpdateTariffRestriction(ctx, tariffRestrictionParams)
		}

		if dto.DayOfWeek != nil {
			r.replaceWeekdays(ctx, *id, dto)
		}
	}
}

func (r *TariffRestrictionResolver) replaceWeekdays(ctx context.Context, tariffRestrictionID int64, dto *TariffRestrictionDto) {
	r.Repository.UnsetTariffRestrictionWeekdays(ctx, tariffRestrictionID)

	if weekdays, err := r.Repository.ListWeekdays(ctx); err == nil {
		filteredWeekdays := []*db.Weekday{}

		for _, weekday := range weekdays {
			if util.StringsContainString(dto.DayOfWeek, weekday.Text) {
				filteredWeekdays = append(filteredWeekdays, &weekday)
			}
		}

		for _, weekday := range filteredWeekdays {
			r.Repository.SetTariffRestrictionWeekday(ctx, db.SetTariffRestrictionWeekdayParams{
				TariffRestrictionID: tariffRestrictionID,
				WeekdayID:           weekday.ID,
			})
		}
	}
}
