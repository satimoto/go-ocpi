package dto

import (
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
)

type ExceptionalPeriodDto struct {
	PeriodBegin *time.Time `json:"period_begin"`
	PeriodEnd   *time.Time `json:"period_end"`
}

func (r *ExceptionalPeriodDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewExceptionalPeriodDto(exceptionalPeriod db.ExceptionalPeriod) *ExceptionalPeriodDto {
	return &ExceptionalPeriodDto{
		PeriodBegin: &exceptionalPeriod.PeriodBegin,
		PeriodEnd:   &exceptionalPeriod.PeriodEnd,
	}
}

type OpeningTimeDto struct {
	RegularHours        []*RegularHourDto       `json:"regular_hours"`
	Twentyfourseven     bool                    `json:"twentyfourseven"`
	ExceptionalOpenings []*ExceptionalPeriodDto `json:"exceptional_openings"`
	ExceptionalClosings []*ExceptionalPeriodDto `json:"exceptional_closings"`
}

func (r *OpeningTimeDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewOpeningTimeDto(openingTime db.OpeningTime) *OpeningTimeDto {
	return &OpeningTimeDto{
		Twentyfourseven: openingTime.Twentyfourseven,
	}
}

type RegularHourDto struct {
	Weekday     int16  `json:"weekday"`
	PeriodBegin string `json:"period_begin"`
	PeriodEnd   string `json:"period_end"`
}

func (r *RegularHourDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewRegularHourDto(regularHour db.RegularHour) *RegularHourDto {
	return &RegularHourDto{
		Weekday:     regularHour.Weekday,
		PeriodBegin: regularHour.PeriodBegin,
		PeriodEnd:   regularHour.PeriodEnd,
	}
}
