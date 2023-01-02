package calibration

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	coreDto "github.com/satimoto/go-ocpi/internal/dto"
	"github.com/satimoto/go-ocpi/internal/metric"
)

func (r *CalibrationResolver) CreateCalibrationDto(ctx context.Context, calibration db.Calibration) *coreDto.CalibrationDto {
	response := coreDto.NewCalibrationDto(calibration)

	calibrationValues, err := r.Repository.ListCalibrationValues(ctx, calibration.ID)

	if err != nil {
		metrics.RecordError("OCPI222", "Error retrieving image", err)
		log.Printf("OCPI222: CalibrationID=%v", calibration.ID)
		return response
	}

	response.SignedValues = r.CreateCalibrationValueListDto(ctx, calibrationValues)

	return response
}

func (r *CalibrationResolver) CreateCalibrationValueDto(ctx context.Context, calibrationValue db.CalibrationValue) *coreDto.CalibrationValueDto {
	return coreDto.NewCalibrationValueDto(calibrationValue)
}

func (r *CalibrationResolver) CreateCalibrationValueListDto(ctx context.Context, calibrationValues []db.CalibrationValue) []*coreDto.CalibrationValueDto {
	list := []*coreDto.CalibrationValueDto{}

	for _, calibrationValue := range calibrationValues {
		list = append(list, r.CreateCalibrationValueDto(ctx, calibrationValue))
	}

	return list
}
