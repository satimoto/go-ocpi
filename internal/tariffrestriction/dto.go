package tariffrestriction

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *TariffRestrictionResolver) CreateTariffRestrictionDto(ctx context.Context, tariffRestriction db.TariffRestriction) *coreDto.TariffRestrictionDto {
	response := coreDto.NewTariffRestrictionDto(tariffRestriction)

	weekdays, err := r.Repository.ListTariffRestrictionWeekdays(ctx, tariffRestriction.ID)

	if err != nil {
		metrics.RecordError("OCPI260", "Error listing tariff restriction weekdays", err)
		log.Printf("OCPI260: TariffRestrictionID=%v", tariffRestriction.ID)
	}

	if len(weekdays) > 0 {
		response.DayOfWeek = r.CreateWeekdayListDto(ctx, weekdays)
	}

	return response
}

func (r *TariffRestrictionResolver) CreateWeekdayListDto(ctx context.Context, weekdays []db.Weekday) []*string {
	list := []*string{}

	for _, weekday := range weekdays {
		text := weekday.Text
		list = append(list, &text)
	}

	return list
}
