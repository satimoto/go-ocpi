package dto

import (
	"net/http"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type OcpiSessionsDto struct {
	Data          []*SessionDto `json:"data,omitempty"`
	StatusCode    int16         `json:"status_code"`
	StatusMessage string        `json:"status_message"`
	Timestamp     ocpitype.Time `json:"timestamp"`
}

type SessionDto struct {
	ID              *string                      `json:"id"`
	AuthorizationID *string                      `json:"authorization_id,omitempty"`
	StartDatetime   *time.Time                   `json:"start_datetime"`
	EndDatetime     *time.Time                   `json:"end_datetime,omitempty"`
	Kwh             *float64                     `json:"kwh"`
	AuthID          *string                      `json:"auth_id"`
	AuthMethod      *db.AuthMethodType           `json:"auth_method"`
	Location        *LocationDto                 `json:"location"`
	MeterID         *string                      `json:"meter_id,omitempty"`
	Currency        *string                      `json:"currency"`
	ChargingPeriods []*coreDto.ChargingPeriodDto `json:"charging_periods"`
	TotalCost       *float64                     `json:"total_cost,omitempty"`
	Status          *db.SessionStatusType        `json:"status"`
	LastUpdated     *time.Time                   `json:"last_updated"`
}

func (r *SessionDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewSessionDto(session db.Session) *SessionDto {
	return &SessionDto{
		ID:              &session.Uid,
		AuthorizationID: util.NilString(session.AuthorizationID),
		StartDatetime:   &session.StartDatetime,
		EndDatetime:     util.NilTime(session.EndDatetime.Time),
		Kwh:             &session.Kwh,
		AuthID:          &session.AuthID,
		AuthMethod:      &session.AuthMethod,
		MeterID:         util.NilString(session.MeterID),
		Currency:        &session.Currency,
		TotalCost:       util.NilFloat64(session.TotalCost.Float64),
		Status:          &session.Status,
		LastUpdated:     &session.LastUpdated,
	}
}
