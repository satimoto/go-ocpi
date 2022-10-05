package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

type ElementRestrictionDto struct {
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

func (r *ElementRestrictionDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewElementRestrictionDto(elementRestriction db.ElementRestriction) *ElementRestrictionDto {
	return &ElementRestrictionDto{
		StartTime:   util.NilString(elementRestriction.StartTime),
		EndTime:     util.NilString(elementRestriction.EndTime),
		StartDate:   util.NilString(elementRestriction.StartDate),
		EndDate:     util.NilString(elementRestriction.EndDate),
		MinKwh:      util.NilFloat64(elementRestriction.MinKwh.Float64),
		MaxKwh:      util.NilFloat64(elementRestriction.MaxKwh.Float64),
		MinPower:    util.NilFloat64(elementRestriction.MinPower.Float64),
		MaxPower:    util.NilFloat64(elementRestriction.MaxPower.Float64),
		MinDuration: util.NilInt32(elementRestriction.MinDuration.Int32),
		MaxDuration: util.NilInt32(elementRestriction.MaxDuration.Int32),
	}
}
