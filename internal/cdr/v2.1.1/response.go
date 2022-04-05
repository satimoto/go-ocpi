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

type CdrDto struct {
	ID               *string                             `json:"id"`
	AuthorizationID  *string                             `json:"authorization_id,omitempty"`
	StartDateTime    *time.Time                          `json:"start_date_time"`
	StopDateTime     *time.Time                          `json:"stop_date_time,omitempty"`
	AuthID           *string                             `json:"auth_id"`
	AuthMethod       *db.AuthMethodType                  `json:"auth_method"`
	Location         *location.LocationDto               `json:"location"`
	MeterID          *string                             `json:"meter_id,omitempty"`
	Currency         *string                             `json:"currency"`
	Tariffs          []*tariff.TariffPushDto             `json:"tariffs"`
	ChargingPeriods  []*chargingperiod.ChargingPeriodDto `json:"charging_periods"`
	SignedData       *calibration.CalibrationDto         `json:"signed_data,omitempty"`
	TotalCost        *float64                            `json:"total_cost"`
	TotalEnergy      *float64                            `json:"total_energy"`
	TotalTime        *float64                            `json:"total_time"`
	TotalParkingTime *float64                            `json:"total_parking_time,omitempty"`
	Remark           *string                             `json:"remark,omitempty"`
	LastUpdated      *time.Time                          `json:"last_updated"`
}

func (r *CdrDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCdrDto(cdr db.Cdr) *CdrDto {
	return &CdrDto{
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

func NewCreateCdrParams(dto *CdrDto) db.CreateCdrParams {
	return db.CreateCdrParams{
		Uid:              *dto.ID,
		AuthorizationID:  util.SqlNullString(dto.AuthorizationID),
		StartDateTime:    *dto.StartDateTime,
		StopDateTime:     util.SqlNullTime(dto.StopDateTime),
		AuthID:           *dto.AuthID,
		AuthMethod:       *dto.AuthMethod,
		MeterID:          util.SqlNullString(dto.MeterID),
		Currency:         *dto.Currency,
		TotalCost:        *dto.TotalCost,
		TotalEnergy:      *dto.TotalEnergy,
		TotalTime:        *dto.TotalTime,
		TotalParkingTime: util.SqlNullFloat64(dto.TotalParkingTime),
		Remark:           util.SqlNullString(dto.Remark),
		LastUpdated:      *dto.LastUpdated,
	}
}

func (r *CdrResolver) CreateCdrDto(ctx context.Context, cdr db.Cdr) *CdrDto {
	response := NewCdrDto(cdr)

	if chargingPeriods, err := r.Repository.ListCdrChargingPeriods(ctx, cdr.ID); err == nil {
		response.ChargingPeriods = r.ChargingPeriodResolver.CreateChargingPeriodListDto(ctx, chargingPeriods)
	}

	if location, err := r.LocationResolver.Repository.GetLocation(ctx, cdr.LocationID); err == nil {
		response.Location = r.LocationResolver.CreateLocationDto(ctx, location)
	}

	if cdr.CalibrationID.Valid {
		if calibration, err := r.CalibrationResolver.Repository.GetCalibration(ctx, cdr.CalibrationID.Int64); err == nil {
			response.SignedData = r.CalibrationResolver.CreateCalibrationDto(ctx, calibration)
		}
	}

	if tariffs, err := r.TariffResolver.Repository.ListTariffsByCdr(ctx, util.SqlNullInt64(cdr.ID)); err == nil {
		response.Tariffs = r.TariffResolver.CreateTariffPushListDto(ctx, tariffs)
	}

	return response
}

func (r *CdrResolver) CreateCdrListDto(ctx context.Context, cdrs []db.Cdr) []render.Renderer {
	list := []render.Renderer{}
	for _, cdr := range cdrs {
		list = append(list, r.CreateCdrDto(ctx, cdr))
	}
	return list
}
