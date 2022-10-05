package elementrestriction

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
)

func (r *ElementRestrictionResolver) CreateElementRestrictionDto(ctx context.Context, elementRestriction db.ElementRestriction) *coreDto.ElementRestrictionDto {
	response := coreDto.NewElementRestrictionDto(elementRestriction)

	weekdays, err := r.Repository.ListElementRestrictionWeekdays(ctx, elementRestriction.ID)

	if err != nil {
		util.LogOnError("OCPI228", "Error listing element restriction weekdays", err)
		log.Printf("OCPI228: ElementRestrictionID=%v", elementRestriction.ID)
		return response
	}

	if len(weekdays) > 0 {
		response.DayOfWeek = r.CreateWeekdayListDto(ctx, weekdays)
	}

	return response
}

func (r *ElementRestrictionResolver) CreateWeekdayListDto(ctx context.Context, weekdays []db.Weekday) []*string {
	list := []*string{}

	for _, weekday := range weekdays {
		text := weekday.Text
		list = append(list, &text)
	}

	return list
}
