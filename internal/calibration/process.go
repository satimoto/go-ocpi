package calibration

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *CalibrationResolver) CreateCalibration(ctx context.Context, dto *CalibrationDto) *db.Calibration {
	if dto != nil {
		calibrationParams := NewCreateCalibrationParams(dto)

		calibration, err := r.Repository.CreateCalibration(ctx, calibrationParams)

		if err != nil {
			util.LogOnError("OCPI017", "Error creating calibration", err)
			log.Printf("OCPI017: Params=%#v", calibrationParams)
			return nil
		}

		r.createCalibrationValues(ctx, &calibration.ID, *dto)

		return &calibration
	}

	return nil
}

func (r *CalibrationResolver) createCalibrationValues(ctx context.Context, calibrationID *int64, dto CalibrationDto) {
	if calibrationID != nil {
		for _, calibrationValue := range dto.SignedValues {
			calibrationValueParams := NewCreateCalibrationValueParams(*calibrationID, calibrationValue)

			_, err := r.Repository.CreateCalibrationValue(ctx, calibrationValueParams)

			if err != nil {
				util.LogOnError("OCPI018", "Error creating calibration value", err)
				log.Printf("OCPI018: Params=%#v", calibrationValueParams)
			}
		}
	}
}
