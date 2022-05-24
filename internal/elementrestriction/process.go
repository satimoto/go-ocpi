package elementrestriction

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *ElementRestrictionResolver) ReplaceElementRestriction(ctx context.Context, id *int64, dto *ElementRestrictionDto) {
	if dto != nil {
		if id == nil {
			elementRestrictionParams := NewCreateElementRestrictionParams(dto)
			elementRestriction, err := r.Repository.CreateElementRestriction(ctx, elementRestrictionParams)
			
			if err != nil {
				util.LogOnError("OCPI092", "Error creating element restriction", err)
				log.Printf("OCPI092: Params=%#v", elementRestrictionParams)
				return
			}

			id = &elementRestriction.ID
		} else {
			elementRestrictionParams := NewUpdateElementRestrictionParams(*id, dto)
			_, err := r.Repository.UpdateElementRestriction(ctx, elementRestrictionParams)

			if err != nil {
				util.LogOnError("OCPI093", "Error updating element restriction", err)
				log.Printf("OCPI093: Params=%#v", elementRestrictionParams)
			}
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
			setElementRestrictionWeekdayParams := db.SetElementRestrictionWeekdayParams{
				ElementRestrictionID: elementRestrictionID,
				WeekdayID:            weekday.ID,
			}
			err := r.Repository.SetElementRestrictionWeekday(ctx, setElementRestrictionWeekdayParams)

			if err != nil {
				util.LogOnError("OCPI094", "Error setting element restriction weekends", err)
				log.Printf("OCPI094: Params=%#v", setElementRestrictionWeekdayParams)
				return
			}

		}
	}
}
