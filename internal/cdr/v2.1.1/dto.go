package cdr

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi-api/internal/calibration"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
)

type OcpiCdrsDto struct {
	Data          []*CdrDto `json:"data,omitempty"`
	StatusCode    int16     `json:"status_code"`
	StatusMessage string    `json:"status_message"`
	Timestamp     time.Time `json:"timestamp"`
}

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
	Tariffs          []*tariff.TariffDto                 `json:"tariffs"`
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
		AuthorizationID:  util.NilString(cdr.AuthorizationID),
		StartDateTime:    &cdr.StartDateTime,
		StopDateTime:     util.NilTime(cdr.StopDateTime.Time),
		AuthID:           &cdr.AuthID,
		AuthMethod:       &cdr.AuthMethod,
		MeterID:          util.NilString(cdr.MeterID),
		Currency:         &cdr.Currency,
		TotalCost:        &cdr.TotalCost,
		TotalEnergy:      &cdr.TotalEnergy,
		TotalTime:        &cdr.TotalTime,
		TotalParkingTime: util.NilFloat64(cdr.TotalParkingTime.Float64),
		Remark:           util.NilString(cdr.Remark),
		LastUpdated:      &cdr.LastUpdated,
	}
}

func (r *CdrResolver) CreateCdrDto(ctx context.Context, cdr db.Cdr) *CdrDto {
	response := NewCdrDto(cdr)

	chargingPeriods, err := r.Repository.ListCdrChargingPeriods(ctx, cdr.ID)

	if err != nil {
		util.LogOnError("OCPI223", "Error listing cdr charging periods", err)
		log.Printf("OCPI223: CdrID=%v", cdr.ID)
	} else {
		response.ChargingPeriods = r.ChargingPeriodResolver.CreateChargingPeriodListDto(ctx, chargingPeriods)
	}

	location, err := r.LocationResolver.Repository.GetLocation(ctx, cdr.LocationID)

	if err != nil {
		util.LogOnError("OCPI224", "Error retrieving cdr location", err)
		log.Printf("OCPI224: LocationID=%v", cdr.LocationID)
	} else {
		response.Location = r.LocationResolver.CreateLocationDto(ctx, location)
	}

	if cdr.CalibrationID.Valid {
		calibration, err := r.CalibrationResolver.Repository.GetCalibration(ctx, cdr.CalibrationID.Int64)

		if err != nil {
			util.LogOnError("OCPI225", "Error retrieving cdr calibration", err)
			log.Printf("OCPI225: CalibrationID=%v", cdr.CalibrationID)
		} else {
			response.SignedData = r.CalibrationResolver.CreateCalibrationDto(ctx, calibration)
		}
	}

	tariffs, err := r.TariffResolver.Repository.ListTariffsByCdr(ctx, util.SqlNullInt64(cdr.ID))

	if err != nil {
		util.LogOnError("OCPI226", "Error listing cdr tariffs", err)
		log.Printf("OCPI226: CdrID=%v", cdr.ID)
	} else {
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
