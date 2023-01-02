package dto

import (
	"net/http"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/ocpitype"
)

type OcpiCdrsDto struct {
	Data          []*CdrDto     `json:"data,omitempty"`
	StatusCode    int16         `json:"status_code"`
	StatusMessage string        `json:"status_message"`
	Timestamp     ocpitype.Time `json:"timestamp"`
}

type CdrDto struct {
	ID               *string                      `json:"id"`
	AuthorizationID  *string                      `json:"authorization_id,omitempty"`
	StartDateTime    *ocpitype.Time               `json:"start_date_time"`
	StopDateTime     *ocpitype.Time               `json:"stop_date_time,omitempty"`
	AuthID           *string                      `json:"auth_id"`
	AuthMethod       *db.AuthMethodType           `json:"auth_method"`
	Location         *LocationDto                 `json:"location"`
	MeterID          *string                      `json:"meter_id,omitempty"`
	Currency         *string                      `json:"currency"`
	Tariffs          []*TariffDto                 `json:"tariffs"`
	ChargingPeriods  []*coreDto.ChargingPeriodDto `json:"charging_periods"`
	SignedData       *coreDto.CalibrationDto      `json:"signed_data,omitempty"`
	TotalCost        *float64                     `json:"total_cost"`
	TotalEnergy      *float64                     `json:"total_energy"`
	TotalTime        *float64                     `json:"total_time"`
	TotalParkingTime *float64                     `json:"total_parking_time,omitempty"`
	Remark           *string                      `json:"remark,omitempty"`
	LastUpdated      *ocpitype.Time               `json:"last_updated"`
}

func (r *CdrDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCdrDto(cdr db.Cdr) *CdrDto {
	return &CdrDto{
		ID:               &cdr.Uid,
		AuthorizationID:  util.NilString(cdr.AuthorizationID),
		StartDateTime:    ocpitype.NilOcpiTime(&cdr.StartDateTime),
		StopDateTime:     ocpitype.NilOcpiTime(&cdr.StopDateTime.Time),
		AuthID:           &cdr.AuthID,
		AuthMethod:       &cdr.AuthMethod,
		MeterID:          util.NilString(cdr.MeterID),
		Currency:         &cdr.Currency,
		TotalCost:        &cdr.TotalCost,
		TotalEnergy:      &cdr.TotalEnergy,
		TotalTime:        &cdr.TotalTime,
		TotalParkingTime: util.NilFloat64(cdr.TotalParkingTime.Float64),
		Remark:           util.NilString(cdr.Remark),
		LastUpdated:      ocpitype.NilOcpiTime(&cdr.LastUpdated),
	}
}
