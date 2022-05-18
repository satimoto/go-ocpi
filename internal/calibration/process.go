package calibration

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
)

func (r *CalibrationResolver) CreateCalibration(ctx context.Context, dto *CalibrationDto) *db.Calibration {
	if dto != nil {
		calibrationParams := NewCreateCalibrationParams(dto)

		calibration, err := r.Repository.CreateCalibration(ctx, calibrationParams)

		if err == nil {
			r.createCalibrationValues(ctx, &calibration.ID, *dto)

			return &calibration
		}
	}

	return nil
}

func (r *CalibrationResolver) createCalibrationValues(ctx context.Context, calibrationID *int64, dto CalibrationDto) {
	if calibrationID != nil {
		for _, calibrationValue := range dto.SignedValues {
			calibrationValueParams := NewCreateCalibrationValueParams(*calibrationID, calibrationValue)
			r.Repository.CreateCalibrationValue(ctx, calibrationValueParams)
		}
	}
}
