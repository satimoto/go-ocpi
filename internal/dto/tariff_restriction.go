package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
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
		StartTime2: util.NilString(tariffRestriction.StartTime2),
		EndTime2:   util.NilString(tariffRestriction.EndTime2),
	}
}
