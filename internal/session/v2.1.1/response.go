package session

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type SessionPayload struct {
	ID              *string                                 `json:"id"`
	StartDatetime   *time.Time                              `json:"start_datetime"`
	EndDatetime     *time.Time                              `json:"end_datetime,omitempty"`
	Kwh             *float64                                `json:"kwh"`
	AuthID          *string                                 `json:"auth_id"`
	AuthMethod      *db.AuthMethodType                      `json:"auth_method"`
	Location        *location.LocationPayload               `json:"location"`
	MeterID         *string                                 `json:"meter_id,omitempty"`
	Currency        *string                                 `json:"currency"`
	ChargingPeriods []*chargingperiod.ChargingPeriodPayload `json:"charging_periods"`
	TotalCost       *float64                                `json:"total_cost,omitempty"`
	Status          *db.SessionStatusType                   `json:"status"`
	LastUpdated     *time.Time                              `json:"last_updated"`
}

func (r *SessionPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewSessionPayload(session db.Session) *SessionPayload {
	return &SessionPayload{
		ID:            &session.Uid,
		StartDatetime: &session.StartDatetime,
		EndDatetime:   util.NilTime(session.EndDatetime.Time),
		Kwh:           &session.Kwh,
		AuthID:        &session.AuthID,
		AuthMethod:    &session.AuthMethod,
		MeterID:       util.NilString(session.MeterID.String),
		Currency:      &session.Currency,
		TotalCost:     util.NilFloat64(session.TotalCost.Float64),
		Status:        &session.Status,
		LastUpdated:   &session.LastUpdated,
	}
}

func NewCreateSessionParams(payload *SessionPayload) db.CreateSessionParams {
	return db.CreateSessionParams{
		Uid:           *payload.ID,
		StartDatetime: *payload.StartDatetime,
		EndDatetime:   util.SqlNullTime(payload.EndDatetime),
		Kwh:           *payload.Kwh,
		AuthID:        *payload.AuthID,
		AuthMethod:    *payload.AuthMethod,
		MeterID:       util.SqlNullString(payload.MeterID),
		Currency:      *payload.Currency,
		TotalCost:     util.SqlNullFloat64(payload.TotalCost),
		Status:        *payload.Status,
		LastUpdated:   *payload.LastUpdated,
	}
}

func NewUpdateSessionByUidParams(session db.Session) db.UpdateSessionByUidParams {
	return db.UpdateSessionByUidParams{
		Uid:           session.Uid,
		StartDatetime: session.StartDatetime,
		EndDatetime:   session.EndDatetime,
		Kwh:           session.Kwh,
		AuthID:        session.AuthID,
		AuthMethod:    session.AuthMethod,
		LocationID:    session.LocationID,
		MeterID:       session.MeterID,
		Currency:      session.Currency,
		TotalCost:     session.TotalCost,
		Status:        session.Status,
		LastUpdated:   session.LastUpdated,
	}
}

func (r *SessionResolver) CreateSessionPayload(ctx context.Context, session db.Session) *SessionPayload {
	response := NewSessionPayload(session)

	if chargingPeriods, err := r.Repository.ListSessionChargingPeriods(ctx, session.ID); err == nil {
		response.ChargingPeriods = r.ChargingPeriodResolver.CreateChargingPeriodListPayload(ctx, chargingPeriods)
	}

	if location, err := r.LocationResolver.Repository.GetLocation(ctx, session.LocationID); err == nil {
		response.Location = r.LocationResolver.CreateLocationPayload(ctx, location)
	}

	return response
}

func (r *SessionResolver) CreateSessionListPayload(ctx context.Context, sessions []db.Session) []render.Renderer {
	list := []render.Renderer{}
	for _, session := range sessions {
		list = append(list, r.CreateSessionPayload(ctx, session))
	}
	return list
}
