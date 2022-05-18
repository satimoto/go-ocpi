package session

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
)

type OcpiSessionsDto struct {
	Data          []*SessionDto `json:"data,omitempty"`
	StatusCode    int16         `json:"status_code"`
	StatusMessage string        `json:"status_message"`
	Timestamp     time.Time     `json:"timestamp"`
}

type SessionDto struct {
	ID              *string                             `json:"id"`
	AuthorizationID *string                             `json:"authorization_id,omitempty"`
	StartDatetime   *time.Time                          `json:"start_datetime"`
	EndDatetime     *time.Time                          `json:"end_datetime,omitempty"`
	Kwh             *float64                            `json:"kwh"`
	AuthID          *string                             `json:"auth_id"`
	AuthMethod      *db.AuthMethodType                  `json:"auth_method"`
	Location        *location.LocationDto               `json:"location"`
	MeterID         *string                             `json:"meter_id,omitempty"`
	Currency        *string                             `json:"currency"`
	ChargingPeriods []*chargingperiod.ChargingPeriodDto `json:"charging_periods"`
	TotalCost       *float64                            `json:"total_cost,omitempty"`
	Status          *db.SessionStatusType               `json:"status"`
	LastUpdated     *time.Time                          `json:"last_updated"`
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

func (r *SessionResolver) CreateSessionDto(ctx context.Context, session db.Session) *SessionDto {
	response := NewSessionDto(session)

	if chargingPeriods, err := r.Repository.ListSessionChargingPeriods(ctx, session.ID); err == nil {
		response.ChargingPeriods = r.ChargingPeriodResolver.CreateChargingPeriodListDto(ctx, chargingPeriods)
	}

	if location, err := r.LocationResolver.Repository.GetLocation(ctx, session.LocationID); err == nil {
		response.Location = r.LocationResolver.CreateLocationDto(ctx, location)
	}

	return response
}

func (r *SessionResolver) CreateSessionListDto(ctx context.Context, sessions []db.Session) []render.Renderer {
	list := []render.Renderer{}
	for _, session := range sessions {
		list = append(list, r.CreateSessionDto(ctx, session))
	}
	return list
}
