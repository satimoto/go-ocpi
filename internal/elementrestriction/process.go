package elementrestriction

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *ElementRestrictionResolver) ReplaceElementRestriction(ctx context.Context, id *int64, dto *ElementRestrictionDto) {
	if dto != nil {
		if id == nil {
			elementRestrictionParams := NewCreateElementRestrictionParams(dto)

			if elementRestriction, err := r.Repository.CreateElementRestriction(ctx, elementRestrictionParams); err == nil {
				id = &elementRestriction.ID
			}
		} else {
			elementRestrictionParams := NewUpdateElementRestrictionParams(*id, dto)

			r.Repository.UpdateElementRestriction(ctx, elementRestrictionParams)
		}

		if dto.DayOfWeek != nil {
			r.replaceWeekdays(ctx, *id, dto)
		}
	}
}

func (r *ElementRestrictionResolver) replaceWeekdays(ctx context.Context, elementRestrictionID int64, dto *ElementRestrictionDto) {
	r.Repository.UnsetElementRestrictionWeekdays(ctx, elementRestrictionID)

	if weekdays, err := r.Repository.ListWeekdays(ctx); err == nil {
		filteredWeekdays := []*db.Weekday{}

		for _, weekday := range weekdays {
			if util.StringsContainString(dto.DayOfWeek, weekday.Text) {
				filteredWeekdays = append(filteredWeekdays, &weekday)
			}
		}

		for _, weekday := range filteredWeekdays {
			r.Repository.SetElementRestrictionWeekday(ctx, db.SetElementRestrictionWeekdayParams{
				ElementRestrictionID: elementRestrictionID,
				WeekdayID:            weekday.ID,
			})
		}
	}
}
