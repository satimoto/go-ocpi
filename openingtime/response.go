package openingtime

import (
	"context"
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/db"
)

type ExceptionalPeriodPayload struct {
	PeriodBegin *time.Time `json:"period_begin"`
	PeriodEnd   *time.Time `json:"period_end"`
}

func (r *ExceptionalPeriodPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewExceptionalPeriodPayload(exceptionalPeriod db.ExceptionalPeriod) *ExceptionalPeriodPayload {
	return &ExceptionalPeriodPayload{
		PeriodBegin: &exceptionalPeriod.PeriodBegin,
		PeriodEnd:   &exceptionalPeriod.PeriodEnd,
	}
}

func NewCreateExceptionalPeriodParams(id int64, periodType db.PeriodType, payload *ExceptionalPeriodPayload) db.CreateExceptionalPeriodParams {
	return db.CreateExceptionalPeriodParams{
		OpeningTimeID: id,
		PeriodType: periodType,
		PeriodBegin:   *payload.PeriodBegin,
		PeriodEnd:     *payload.PeriodEnd,
	}
}
func (r *OpeningTimeResolver) CreateExceptionalPeriodPayload(ctx context.Context, exceptionalPeriod db.ExceptionalPeriod) *ExceptionalPeriodPayload {
	return NewExceptionalPeriodPayload(exceptionalPeriod)
}

func (r *OpeningTimeResolver) CreateExceptionalPeriodListPayload(ctx context.Context, exceptionalPeriods []db.ExceptionalPeriod) []*ExceptionalPeriodPayload {
	list := []*ExceptionalPeriodPayload{}
	for _, exceptionalPeriod := range exceptionalPeriods {
		list = append(list, r.CreateExceptionalPeriodPayload(ctx, exceptionalPeriod))
	}
	return list
}

type OpeningTimePayload struct {
	RegularHours        []*RegularHourPayload       `json:"regular_hours"`
	Twentyfourseven     bool                        `json:"twentyfourseven"`
	ExceptionalOpenings []*ExceptionalPeriodPayload `json:"exceptional_openings"`
	ExceptionalClosings []*ExceptionalPeriodPayload `json:"exceptional_closings"`
}

func (r *OpeningTimePayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewOpeningTimePayload(openingTime db.OpeningTime) *OpeningTimePayload {
	return &OpeningTimePayload{
		Twentyfourseven: openingTime.Twentyfourseven,
	}
}

func NewUpdateOpeningTimeParams(id int64, payload *OpeningTimePayload) db.UpdateOpeningTimeParams {
	return db.UpdateOpeningTimeParams{
		ID:              id,
		Twentyfourseven: payload.Twentyfourseven,
	}
}

func (r *OpeningTimeResolver) CreateOpeningTimePayload(ctx context.Context, openingTime db.OpeningTime) *OpeningTimePayload {
	response := NewOpeningTimePayload(openingTime)

	if regularHours, err := r.Repository.ListRegularHours(ctx, openingTime.ID); err == nil {
		response.RegularHours = r.CreateRegularHourListPayload(ctx, regularHours)
	}

	if exceptionalOpenings, err := r.Repository.ListExceptionalOpeningPeriods(ctx, openingTime.ID); err == nil {
		response.ExceptionalOpenings = r.CreateExceptionalPeriodListPayload(ctx, exceptionalOpenings)
	}

	if exceptionalClosings, err := r.Repository.ListExceptionalClosingPeriods(ctx, openingTime.ID); err == nil {
		response.ExceptionalClosings = r.CreateExceptionalPeriodListPayload(ctx, exceptionalClosings)
	}

	return response
}

type RegularHourPayload struct {
	Weekday     int16  `json:"weekday"`
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
}

func (r *RegularHourPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewRegularHourPayload(regularHour db.RegularHour) *RegularHourPayload {
	return &RegularHourPayload{
		Weekday:     regularHour.Weekday,
		PeriodBegin: regularHour.PeriodBegin,
		PeriodEnd:   regularHour.PeriodEnd,
	}
}

func NewCreateRegularHourParams(id int64, payload *RegularHourPayload) db.CreateRegularHourParams {
	return db.CreateRegularHourParams{
		OpeningTimeID: id,
		Weekday:       payload.Weekday,
		PeriodBegin:   payload.PeriodBegin,
		PeriodEnd:     payload.PeriodEnd,
	}
}

func (r *OpeningTimeResolver) CreateRegularHourPayload(ctx context.Context, regularHour db.RegularHour) *RegularHourPayload {
	return NewRegularHourPayload(regularHour)
}

func (r *OpeningTimeResolver) CreateRegularHourListPayload(ctx context.Context, regularHours []db.RegularHour) []*RegularHourPayload {
	list := []*RegularHourPayload{}
	for _, regularHour := range regularHours {
		list = append(list, r.CreateRegularHourPayload(ctx, regularHour))
	}
	return list
}
