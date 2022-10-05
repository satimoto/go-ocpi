package tariffrestriction

import (
	"context"
	"database/sql"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *TariffRestrictionResolver) ReplaceTariffByIdentifierRestriction(ctx context.Context, id *sql.NullInt64, tariffRestrictionDto *coreDto.TariffRestrictionDto) {
	if tariffRestrictionDto != nil {
		if id.Valid {
			tariffRestrictionParams := NewUpdateTariffRestrictionParams(id.Int64, tariffRestrictionDto)
			_, err := r.Repository.UpdateTariffRestriction(ctx, tariffRestrictionParams)

			if err != nil {
				util.LogOnError("OCPI192", "Error updating tariff restriction", err)
				log.Printf("OCPI192: Params=%#v", tariffRestrictionParams)
				return
			}
		} else {
			tariffRestrictionParams := NewCreateTariffRestrictionParams(tariffRestrictionDto)
			tariffRestriction, err := r.Repository.CreateTariffRestriction(ctx, tariffRestrictionParams)
				
			if err != nil {
				util.LogOnError("OCPI191", "Error creating tariff restriction", err)
				log.Printf("OCPI191: Params=%#v", tariffRestrictionParams)
				return
			}
	
			id.Scan(tariffRestriction.ID)
		}

		if tariffRestrictionDto.DayOfWeek != nil {
			r.replaceWeekdays(ctx, id.Int64, tariffRestrictionDto)
		}
	}
}

func (r *TariffRestrictionResolver) replaceWeekdays(ctx context.Context, tariffRestrictionID int64, tariffRestrictionDto *coreDto.TariffRestrictionDto) {
	r.Repository.UnsetTariffRestrictionWeekdays(ctx, tariffRestrictionID)

	if weekdays, err := r.Repository.ListWeekdays(ctx); err == nil {
		filteredWeekdays := []db.Weekday{}

		for _, weekday := range weekdays {
			if util.StringsContainString(tariffRestrictionDto.DayOfWeek, weekday.Text) {
				filteredWeekdays = append(filteredWeekdays, weekday)
			}
		}

		for _, weekday := range filteredWeekdays {
			setTariffRestrictionWeekdayParams := db.SetTariffRestrictionWeekdayParams{
				TariffRestrictionID: tariffRestrictionID,
				WeekdayID:           weekday.ID,
			}
			err := r.Repository.SetTariffRestrictionWeekday(ctx, setTariffRestrictionWeekdayParams)

			if err != nil {
				util.LogOnError("OCPI193", "Error setting tariff restriction weekday", err)
				log.Printf("OCPI193: Params=%#v", setTariffRestrictionWeekdayParams)
			}
		}
	}
}
