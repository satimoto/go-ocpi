package cdr

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/calibration"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type CdrPayload struct {
	ID               *string                                 `json:"id"`
	AuthorizationID  *string                                 `json:"authorization_id,omitempty"`
	StartDateTime    *time.Time                              `json:"start_date_time"`
	StopDateTime     *time.Time                              `json:"stop_date_time,omitempty"`
	AuthID           *string                                 `json:"auth_id"`
	AuthMethod       *db.AuthMethodType                      `json:"auth_method"`
	Location         *location.LocationPayload               `json:"location"`
	MeterID          *string                                 `json:"meter_id,omitempty"`
	Currency         *string                                 `json:"currency"`
	Tariffs          []*tariff.TariffPushPayload             `json:"tariffs"`
	ChargingPeriods  []*chargingperiod.ChargingPeriodPayload `json:"charging_periods"`
	SignedData       *calibration.CalibrationPayload         `json:"signed_data,omitempty"`
	TotalCost        *float64                                `json:"total_cost"`
	TotalEnergy      *float64                                `json:"total_energy"`
	TotalTime        *float64                                `json:"total_time"`
	TotalParkingTime *float64                                `json:"total_parking_time,omitempty"`
	Remark           *string                                 `json:"remark,omitempty"`
	LastUpdated      *time.Time                              `json:"last_updated"`
}

func (r *CdrPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCdrPayload(cdr db.Cdr) *CdrPayload {
	return &CdrPayload{
		ID:               &cdr.Uid,
		AuthorizationID:  util.NilString(cdr.AuthorizationID.String),
		StartDateTime:    &cdr.StartDateTime,
		StopDateTime:     util.NilTime(cdr.StopDateTime.Time),
		AuthID:           &cdr.AuthID,
		AuthMethod:       &cdr.AuthMethod,
		MeterID:          util.NilString(cdr.MeterID.String),
		Currency:         &cdr.Currency,
		TotalCost:        &cdr.TotalCost,
		TotalEnergy:      &cdr.TotalEnergy,
		TotalTime:        &cdr.TotalTime,
		TotalParkingTime: util.NilFloat64(cdr.TotalParkingTime.Float64),
		Remark:           util.NilString(cdr.Remark.String),
		LastUpdated:      &cdr.LastUpdated,
	}
}

func NewCreateCdrParams(payload *CdrPayload) db.CreateCdrParams {
	return db.CreateCdrParams{
		Uid:              *payload.ID,
		AuthorizationID:  util.SqlNullString(payload.AuthorizationID),
		StartDateTime:    *payload.StartDateTime,
		StopDateTime:     util.SqlNullTime(payload.StopDateTime),
		AuthID:           *payload.AuthID,
		AuthMethod:       *payload.AuthMethod,
		MeterID:          util.SqlNullString(payload.MeterID),
		Currency:         *payload.Currency,
		TotalCost:        *payload.TotalCost,
		TotalEnergy:      *payload.TotalEnergy,
		TotalTime:        *payload.TotalTime,
		TotalParkingTime: util.SqlNullFloat64(payload.TotalParkingTime),
		Remark:           util.SqlNullString(payload.Remark),
		LastUpdated:      *payload.LastUpdated,
	}
}

func (r *CdrResolver) CreateCdrPayload(ctx context.Context, cdr db.Cdr) *CdrPayload {
	response := NewCdrPayload(cdr)

	if chargingPeriods, err := r.Repository.ListCdrChargingPeriods(ctx, cdr.ID); err == nil {
		response.ChargingPeriods = r.ChargingPeriodResolver.CreateChargingPeriodListPayload(ctx, chargingPeriods)
	}

	if location, err := r.LocationResolver.Repository.GetLocation(ctx, cdr.LocationID); err == nil {
		response.Location = r.LocationResolver.CreateLocationPayload(ctx, location)
	}

	if cdr.CalibrationID.Valid {
		if calibration, err := r.CalibrationResolver.Repository.GetCalibration(ctx, cdr.CalibrationID.Int64); err == nil {
			response.SignedData = r.CalibrationResolver.CreateCalibrationPayload(ctx, calibration)
		}
	}

	if tariffs, err := r.TariffResolver.Repository.ListTariffsByCdr(ctx, util.SqlNullInt64(cdr.ID)); err == nil {
		response.Tariffs = r.TariffResolver.CreateTariffListPayload(ctx, tariffs)
	}

	return response
}

func (r *CdrResolver) CreateCdrListPayload(ctx context.Context, cdrs []db.Cdr) []render.Renderer {
	list := []render.Renderer{}
	for _, cdr := range cdrs {
		list = append(list, r.CreateCdrPayload(ctx, cdr))
	}
	return list
}
