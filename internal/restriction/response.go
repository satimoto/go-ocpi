package restriction

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type RestrictionDto struct {
	StartTime   *string   `json:"start_time,omitempty"`
	EndTime     *string   `json:"end_time,omitempty"`
	StartDate   *string   `json:"start_date,omitempty"`
	EndDate     *string   `json:"end_date,omitempty"`
	MinKwh      *float64  `json:"min_kwh,omitempty"`
	MaxKwh      *float64  `json:"max_kwh,omitempty"`
	MinPower    *float64  `json:"min_power,omitempty"`
	MaxPower    *float64  `json:"max_power,omitempty"`
	MinDuration *int32    `json:"min_duration,omitempty"`
	MaxDuration *int32    `json:"max_duration,omitempty"`
	DayOfWeek   []*string `json:"day_of_week,omitempty"`
}

func (r *RestrictionDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewRestrictionDto(restriction db.Restriction) *RestrictionDto {
	return &RestrictionDto{
		StartTime:   util.NilString(restriction.StartTime.String),
		EndTime:     util.NilString(restriction.EndTime.String),
		StartDate:   util.NilString(restriction.StartDate.String),
		EndDate:     util.NilString(restriction.EndDate.String),
		MinKwh:      util.NilFloat64(restriction.MinKwh.Float64),
		MaxKwh:      util.NilFloat64(restriction.MaxKwh.Float64),
		MinPower:    util.NilFloat64(restriction.MinPower.Float64),
		MaxPower:    util.NilFloat64(restriction.MaxPower.Float64),
		MinDuration: util.NilInt32(restriction.MinDuration.Int32),
		MaxDuration: util.NilInt32(restriction.MaxDuration.Int32),
	}
}

func NewCreateRestrictionParams(dto *RestrictionDto) db.CreateRestrictionParams {
	return db.CreateRestrictionParams{
		StartTime:   util.SqlNullString(dto.StartTime),
		EndTime:     util.SqlNullString(dto.EndTime),
		StartDate:   util.SqlNullString(dto.StartDate),
		EndDate:     util.SqlNullString(dto.EndDate),
		MinKwh:      util.SqlNullFloat64(dto.MinKwh),
		MaxKwh:      util.SqlNullFloat64(dto.MaxKwh),
		MinPower:    util.SqlNullFloat64(dto.MinPower),
		MaxPower:    util.SqlNullFloat64(dto.MaxPower),
		MinDuration: util.SqlNullInt32(dto.MinDuration),
		MaxDuration: util.SqlNullInt32(dto.MaxDuration),
	}
}

func NewUpdateRestrictionParams(id int64, dto *RestrictionDto) db.UpdateRestrictionParams {
	return db.UpdateRestrictionParams{
		ID:          id,
		StartTime:   util.SqlNullString(dto.StartTime),
		EndTime:     util.SqlNullString(dto.EndTime),
		StartDate:   util.SqlNullString(dto.StartDate),
		EndDate:     util.SqlNullString(dto.EndDate),
		MinKwh:      util.SqlNullFloat64(dto.MinKwh),
		MaxKwh:      util.SqlNullFloat64(dto.MaxKwh),
		MinPower:    util.SqlNullFloat64(dto.MinPower),
		MaxPower:    util.SqlNullFloat64(dto.MaxPower),
		MinDuration: util.SqlNullInt32(dto.MinDuration),
		MaxDuration: util.SqlNullInt32(dto.MaxDuration),
	}
}

func (r *RestrictionResolver) CreateRestrictionDto(ctx context.Context, restriction db.Restriction) *RestrictionDto {
	response := NewRestrictionDto(restriction)

	if weekdays, err := r.Repository.ListRestrictionWeekdays(ctx, restriction.ID); err == nil && len(weekdays) > 0 {
		response.DayOfWeek = r.CreateWeekdayListDto(ctx, weekdays)
	}

	return response
}

func (r *RestrictionResolver) CreateWeekdayListDto(ctx context.Context, weekdays []db.Weekday) []*string {
	list := []*string{}
	for _, weekday := range weekdays {
		list = append(list, &weekday.Text)
	}
	return list
}
