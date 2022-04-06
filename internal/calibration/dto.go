package calibration

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type CalibrationDto struct {
	EncodingMethod        *string                `json:"encoding_method"`
	EncodingMethodVersion *int32                 `json:"encoding_method_version,omitempty"`
	PublicKey             *string                `json:"public_key,omitempty"`
	SignedValues          []*CalibrationValueDto `json:"signed_values"`
	Url                   *string                `json:"url,omitempty"`
}

func (r *CalibrationDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCalibrationDto(calibration db.Calibration) *CalibrationDto {
	return &CalibrationDto{
		EncodingMethod:        &calibration.EncodingMethod,
		EncodingMethodVersion: util.NilInt32(calibration.EncodingMethodVersion.Int32),
		PublicKey:             util.NilString(calibration.PublicKey),
		Url:                   util.NilString(calibration.Url),
	}
}

func (r *CalibrationResolver) CreateCalibrationDto(ctx context.Context, calibration db.Calibration) *CalibrationDto {
	response := NewCalibrationDto(calibration)

	if calibrationValues, err := r.Repository.ListCalibrationValues(ctx, calibration.ID); err == nil {
		response.SignedValues = r.CreateCalibrationValueListDto(ctx, calibrationValues)
	}

	return response
}

type CalibrationValueDto struct {
	Nature     *string `json:"nature"`
	PlainData  *string `json:"plain_data"`
	SignedData *string `json:"signed_data"`
}

func (r *CalibrationValueDto) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCalibrationValueDto(calibrationValue db.CalibrationValue) *CalibrationValueDto {
	return &CalibrationValueDto{
		Nature:     &calibrationValue.Nature,
		PlainData:  &calibrationValue.PlainData,
		SignedData: &calibrationValue.SignedData,
	}
}

func (r *CalibrationResolver) CreateCalibrationValueDto(ctx context.Context, calibrationValue db.CalibrationValue) *CalibrationValueDto {
	return NewCalibrationValueDto(calibrationValue)
}

func (r *CalibrationResolver) CreateCalibrationValueListDto(ctx context.Context, calibrationValues []db.CalibrationValue) []*CalibrationValueDto {
	list := []*CalibrationValueDto{}
	for _, calibrationValue := range calibrationValues {
		list = append(list, r.CreateCalibrationValueDto(ctx, calibrationValue))
	}
	return list
}
