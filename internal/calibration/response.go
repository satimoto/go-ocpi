package calibration

import (
	"context"
	"net/http"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type CalibrationPayload struct {
	EncodingMethod        *string                    `json:"encoding_method"`
	EncodingMethodVersion *int32                     `json:"encoding_method_version,omitempty"`
	PublicKey             *string                    `json:"public_key,omitempty"`
	SignedValues          []*CalibrationValuePayload `json:"signed_values"`
	Url                   *string                    `json:"url,omitempty"`
}

func (r *CalibrationPayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCalibrationPayload(calibration db.Calibration) *CalibrationPayload {
	return &CalibrationPayload{
		EncodingMethod:        &calibration.EncodingMethod,
		EncodingMethodVersion: util.NilInt32(calibration.EncodingMethodVersion.Int32),
		PublicKey:             util.NilString(calibration.PublicKey.String),
		Url:                   util.NilString(calibration.Url.String),
	}
}

func NewCreateCalibrationParams(payload *CalibrationPayload) db.CreateCalibrationParams {
	return db.CreateCalibrationParams{
		EncodingMethod:        *payload.EncodingMethod,
		EncodingMethodVersion: util.SqlNullInt32(payload.EncodingMethodVersion),
		PublicKey:             util.SqlNullString(payload.PublicKey),
		Url:                   util.SqlNullString(payload.Url),
	}
}

func (r *CalibrationResolver) CreateCalibrationPayload(ctx context.Context, calibration db.Calibration) *CalibrationPayload {
	response := NewCalibrationPayload(calibration)

	if calibrationValues, err := r.Repository.ListCalibrationValues(ctx, calibration.ID); err == nil {
		response.SignedValues = r.CreateCalibrationValueListPayload(ctx, calibrationValues)
	}

	return response
}

type CalibrationValuePayload struct {
	Nature     *string `json:"nature"`
	PlainData  *string `json:"plain_data"`
	SignedData *string `json:"signed_data"`
}

func (r *CalibrationValuePayload) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

func NewCalibrationValuePayload(calibrationValue db.CalibrationValue) *CalibrationValuePayload {
	return &CalibrationValuePayload{
		Nature:     &calibrationValue.Nature,
		PlainData:  &calibrationValue.PlainData,
		SignedData: &calibrationValue.SignedData,
	}
}

func NewCreateCalibrationValueParams(id int64, payload *CalibrationValuePayload) db.CreateCalibrationValueParams {
	return db.CreateCalibrationValueParams{
		CalibrationID: id,
		Nature:        *payload.Nature,
		PlainData:     *payload.PlainData,
		SignedData:    *payload.SignedData,
	}
}

func (r *CalibrationResolver) CreateCalibrationValuePayload(ctx context.Context, calibrationValue db.CalibrationValue) *CalibrationValuePayload {
	return NewCalibrationValuePayload(calibrationValue)
}

func (r *CalibrationResolver) CreateCalibrationValueListPayload(ctx context.Context, calibrationValues []db.CalibrationValue) []*CalibrationValuePayload {
	list := []*CalibrationValuePayload{}
	for _, calibrationValue := range calibrationValues {
		list = append(list, r.CreateCalibrationValuePayload(ctx, calibrationValue))
	}
	return list
}
