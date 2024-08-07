package calibration

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *CalibrationResolver) CreateCalibration(ctx context.Context, calibrationDto *coreDto.CalibrationDto) *db.Calibration {
	if calibrationDto != nil {
		calibrationParams := NewCreateCalibrationParams(calibrationDto)

		calibration, err := r.Repository.CreateCalibration(ctx, calibrationParams)

		if err != nil {
			metrics.RecordError("OCPI017", "Error creating calibration", err)
			log.Printf("OCPI017: Params=%#v", calibrationParams)
			return nil
		}

		r.createCalibrationValues(ctx, &calibration.ID, *calibrationDto)

		return &calibration
	}

	return nil
}

func (r *CalibrationResolver) createCalibrationValues(ctx context.Context, calibrationID *int64, calibrationDto coreDto.CalibrationDto) {
	if calibrationID != nil {
		for _, calibrationValue := range calibrationDto.SignedValues {
			calibrationValueParams := NewCreateCalibrationValueParams(*calibrationID, calibrationValue)

			_, err := r.Repository.CreateCalibrationValue(ctx, calibrationValueParams)

			if err != nil {
				metrics.RecordError("OCPI018", "Error creating calibration value", err)
				log.Printf("OCPI018: Params=%#v", calibrationValueParams)
			}
		}
	}
}
