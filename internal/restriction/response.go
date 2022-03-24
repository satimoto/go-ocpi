package restriction

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type RestrictionPayload struct {
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

func (r *RestrictionPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewRestrictionPayload(restriction db.Restriction) *RestrictionPayload {
	return &RestrictionPayload{
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

func NewCreateRestrictionParams(payload *RestrictionPayload) db.CreateRestrictionParams {
	return db.CreateRestrictionParams{
		StartTime:   util.SqlNullString(payload.StartTime),
		EndTime:     util.SqlNullString(payload.EndTime),
		StartDate:   util.SqlNullString(payload.StartDate),
		EndDate:     util.SqlNullString(payload.EndDate),
		MinKwh:      util.SqlNullFloat64(payload.MinKwh),
		MaxKwh:      util.SqlNullFloat64(payload.MaxKwh),
		MinPower:    util.SqlNullFloat64(payload.MinPower),
		MaxPower:    util.SqlNullFloat64(payload.MaxPower),
		MinDuration: util.SqlNullInt32(payload.MinDuration),
		MaxDuration: util.SqlNullInt32(payload.MaxDuration),
	}
}

func NewUpdateRestrictionParams(id int64, payload *RestrictionPayload) db.UpdateRestrictionParams {
	return db.UpdateRestrictionParams{
		ID:          id,
		StartTime:   util.SqlNullString(payload.StartTime),
		EndTime:     util.SqlNullString(payload.EndTime),
		StartDate:   util.SqlNullString(payload.StartDate),
		EndDate:     util.SqlNullString(payload.EndDate),
		MinKwh:      util.SqlNullFloat64(payload.MinKwh),
		MaxKwh:      util.SqlNullFloat64(payload.MaxKwh),
		MinPower:    util.SqlNullFloat64(payload.MinPower),
		MaxPower:    util.SqlNullFloat64(payload.MaxPower),
		MinDuration: util.SqlNullInt32(payload.MinDuration),
		MaxDuration: util.SqlNullInt32(payload.MaxDuration),
	}
}

func (r *RestrictionResolver) CreateRestrictionPayload(ctx context.Context, restriction db.Restriction) *RestrictionPayload {
	response := NewRestrictionPayload(restriction)

	if weekdays, err := r.Repository.ListRestrictionWeekdays(ctx, restriction.ID); err == nil && len(weekdays) > 0 {
		response.DayOfWeek = r.CreateWeekdayListPayload(ctx, weekdays)
	}

	return response
}

func (r *RestrictionResolver) CreateWeekdayListPayload(ctx context.Context, weekdays []db.Weekday) []*string {
	list := []*string{}
	for _, weekday := range weekdays {
		list = append(list, &weekday.Text)
	}
	return list
}
