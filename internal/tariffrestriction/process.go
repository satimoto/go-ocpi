package tariffrestriction

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *TariffRestrictionResolver) ReplaceTariffByIdentifierRestriction(ctx context.Context, id *int64, dto *TariffRestrictionDto) {
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
