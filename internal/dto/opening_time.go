package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type ExceptionalPeriodDto struct {
	PeriodBegin *ocpitype.Time `json:"period_begin"`
	PeriodEnd   *ocpitype.Time `json:"period_end"`
}

func (r *ExceptionalPeriodDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewExceptionalPeriodDto(exceptionalPeriod db.ExceptionalPeriod) *ExceptionalPeriodDto {
	return &ExceptionalPeriodDto{
		PeriodBegin: ocpitype.NilOcpiTime(&exceptionalPeriod.PeriodBegin),
		PeriodEnd:   ocpitype.NilOcpiTime(&exceptionalPeriod.PeriodEnd),
	}
}

type OpeningTimeDto struct {
	RegularHours        []*RegularHourDto       `json:"regular_hours,omitempty"`
	Twentyfourseven     bool                    `json:"twentyfourseven"`
	ExceptionalOpenings []*ExceptionalPeriodDto `json:"exceptional_openings,omitempty"`
	ExceptionalClosings []*ExceptionalPeriodDto `json:"exceptional_closings,omitempty"`
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
