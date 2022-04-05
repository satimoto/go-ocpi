package tariffrestriction

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type TariffRestrictionDto struct {
	StartTime  *string   `json:"start_time,omitempty"`
	EndTime    *string   `json:"end_time,omitempty"`
	StartTime2 *string   `json:"start_time_2,omitempty"`
	EndTime2   *string   `json:"end_time_2,omitempty"`
	DayOfWeek  []*string `json:"day_of_week,omitempty"`
}

func (r *TariffRestrictionDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewTariffRestrictionDto(tariffRestriction db.TariffRestriction) *TariffRestrictionDto {
	return &TariffRestrictionDto{
		StartTime:  &tariffRestriction.StartTime,
		EndTime:    &tariffRestriction.EndTime,
		StartTime2: util.NilString(tariffRestriction.StartTime2.String),
		EndTime2:   util.NilString(tariffRestriction.EndTime2.String),
	}
}

func NewCreateTariffRestrictionParams(dto *TariffRestrictionDto) db.CreateTariffRestrictionParams {
	return db.CreateTariffRestrictionParams{
		StartTime:  *dto.StartTime,
		EndTime:    *dto.EndTime,
		StartTime2: util.SqlNullString(dto.StartTime2),
		EndTime2:   util.SqlNullString(dto.EndTime2),
	}
}

func NewUpdateTariffRestrictionParams(id int64, dto *TariffRestrictionDto) db.UpdateTariffRestrictionParams {
	return db.UpdateTariffRestrictionParams{
		ID:         id,
		StartTime:  *dto.StartTime,
		EndTime:    *dto.EndTime,
		StartTime2: util.SqlNullString(dto.StartTime2),
		EndTime2:   util.SqlNullString(dto.EndTime2),
	}
}

func (r *TariffRestrictionResolver) CreateTariffRestrictionDto(ctx context.Context, tariffRestriction db.TariffRestriction) *TariffRestrictionDto {
	response := NewTariffRestrictionDto(tariffRestriction)

	if weekdays, err := r.Repository.ListTariffRestrictionWeekdays(ctx, tariffRestriction.ID); err == nil && len(weekdays) > 0 {
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
