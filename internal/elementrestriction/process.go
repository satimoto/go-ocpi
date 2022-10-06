package elementrestriction

import (
	"context"
	"database/sql"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *ElementRestrictionResolver) ReplaceElementRestriction(ctx context.Context, id *sql.NullInt64, elementRestrictionDto *coreDto.ElementRestrictionDto) {
	if elementRestrictionDto != nil {
		if id.Valid {
			elementRestrictionParams := NewUpdateElementRestrictionParams(id.Int64, elementRestrictionDto)
			_, err := r.Repository.UpdateElementRestriction(ctx, elementRestrictionParams)

			if err != nil {
				util.LogOnError("OCPI093", "Error updating element restriction", err)
				log.Printf("OCPI093: Params=%#v", elementRestrictionParams)
			}
		} else {
			elementRestrictionParams := NewCreateElementRestrictionParams(elementRestrictionDto)
			elementRestriction, err := r.Repository.CreateElementRestriction(ctx, elementRestrictionParams)

			if err != nil {
				util.LogOnError("OCPI092", "Error creating element restriction", err)
				log.Printf("OCPI092: Params=%#v", elementRestrictionParams)
				return
			}

			id.Scan(elementRestriction.ID)
		}

		if elementRestrictionDto.DayOfWeek != nil {
			r.replaceWeekdays(ctx, id.Int64, elementRestrictionDto)
		}
	}
}

func (r *ElementRestrictionResolver) replaceWeekdays(ctx context.Context, elementRestrictionID int64, elementRestrictionDto *coreDto.ElementRestrictionDto) {
	r.Repository.UnsetElementRestrictionWeekdays(ctx, elementRestrictionID)

	if weekdays, err := r.Repository.ListWeekdays(ctx); err == nil {
		filteredWeekdays := []db.Weekday{}

		for _, weekday := range weekdays {
			if util.StringsContainString(elementRestrictionDto.DayOfWeek, weekday.Text) {
				filteredWeekdays = append(filteredWeekdays, weekday)
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
