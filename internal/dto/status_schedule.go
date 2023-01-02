package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type StatusScheduleDto struct {
	PeriodBegin *ocpitype.Time `json:"period_begin"`
	PeriodEnd   *ocpitype.Time `json:"period_end,omitempty"`
	Status      db.EvseStatus  `json:"status"`
}

func (r *StatusScheduleDto) Render(writer http.ResponseWriter, request *http.Request) error {
	if r.PeriodEnd != nil && r.PeriodEnd.Time().IsZero() {
		r.PeriodEnd = nil
	}

	return nil
}

func NewStatusScheduleDto(statusSchedule db.StatusSchedule) *StatusScheduleDto {
	return &StatusScheduleDto{
		PeriodBegin: ocpitype.NilOcpiTime(&statusSchedule.PeriodBegin),
		PeriodEnd:   ocpitype.NilOcpiTime(&statusSchedule.PeriodEnd.Time),
		Status:      statusSchedule.Status,
	}
}
