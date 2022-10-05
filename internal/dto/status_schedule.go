package dto

import (
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

type StatusScheduleDto struct {
	PeriodBegin *time.Time    `json:"period_begin"`
	PeriodEnd   *time.Time    `json:"period_end,omitempty"`
	Status      db.EvseStatus `json:"status"`
}

func (r *StatusScheduleDto) Render(writer http.ResponseWriter, request *http.Request) error {
	if r.PeriodEnd.IsZero() {
		r.PeriodEnd = nil
	}

	return nil
}

func NewStatusScheduleDto(statusSchedule db.StatusSchedule) *StatusScheduleDto {
	return &StatusScheduleDto{
		PeriodBegin: &statusSchedule.PeriodBegin,
		PeriodEnd:   util.NilTime(statusSchedule.PeriodEnd.Time),
		Status:      statusSchedule.Status,
	}
}
