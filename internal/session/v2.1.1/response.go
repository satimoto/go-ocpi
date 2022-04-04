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
		AuthorizationID: util.NilString(session.AuthorizationID.String),
		StartDatetime:   &session.StartDatetime,
		EndDatetime:     util.NilTime(session.EndDatetime.Time),
		Kwh:             &session.Kwh,
		AuthID:          &session.AuthID,
		AuthMethod:      &session.AuthMethod,
		MeterID:         util.NilString(session.MeterID.String),
		Currency:        &session.Currency,
		TotalCost:       util.NilFloat64(session.TotalCost.Float64),
		Status:          &session.Status,
		LastUpdated:     &session.LastUpdated,
	}
}

func NewCreateSessionParams(dto *SessionDto) db.CreateSessionParams {
	return db.CreateSessionParams{
		Uid:             *dto.ID,
		AuthorizationID: util.SqlNullString(dto.AuthorizationID),
		StartDatetime:   *dto.StartDatetime,
		EndDatetime:     util.SqlNullTime(dto.EndDatetime),
		Kwh:             *dto.Kwh,
		AuthID:          *dto.AuthID,
		AuthMethod:      *dto.AuthMethod,
		MeterID:         util.SqlNullString(dto.MeterID),
		Currency:        *dto.Currency,
		TotalCost:       util.SqlNullFloat64(dto.TotalCost),
		Status:          *dto.Status,
		LastUpdated:     *dto.LastUpdated,
	}
}

func NewUpdateSessionByUidParams(session db.Session) db.UpdateSessionByUidParams {
	return db.UpdateSessionByUidParams{
		Uid:             session.Uid,
		AuthorizationID: session.AuthorizationID,
		StartDatetime:   session.StartDatetime,
		EndDatetime:     session.EndDatetime,
		Kwh:             session.Kwh,
		AuthID:          session.AuthID,
		AuthMethod:      session.AuthMethod,
		LocationID:      session.LocationID,
		MeterID:         session.MeterID,
		Currency:        session.Currency,
		TotalCost:       session.TotalCost,
		Status:          session.Status,
		LastUpdated:     session.LastUpdated,
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
