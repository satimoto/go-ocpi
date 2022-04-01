package calibration

import (
	"context"

	"github.com/satimoto/go-datastore/db"
)

type CalibrationRepository interface {
	CreateCalibration(ctx context.Context, arg db.CreateCalibrationParams) (db.Calibration, error)
	CreateCalibrationValue(ctx context.Context, arg db.CreateCalibrationValueParams) (db.CalibrationValue, error)
	GetCalibration(ctx context.Context, id int64) (db.Calibration, error)
	ListCalibrationValues(ctx context.Context, calibrationID int64) ([]db.CalibrationValue, error)
}

type CalibrationResolver struct {
	Repository CalibrationRepository
}

func NewResolver(repositoryService *db.RepositoryService) *CalibrationResolver {
	repo := CalibrationRepository(repositoryService)
	return &CalibrationResolver{repo}
}

func (r *CalibrationResolver) CreateCalibration(ctx context.Context, payload *CalibrationPayload) *db.Calibration {
	if payload != nil {
		calibrationParams := NewCreateCalibrationParams(payload)

		calibration, err := r.Repository.CreateCalibration(ctx, calibrationParams)

		if err == nil {
			r.createCalibrationValues(ctx, &calibration.ID, *payload)

			return &calibration
		}
	}

	return nil
}

func (r *CalibrationResolver) createCalibrationValues(ctx context.Context, calibrationID *int64, payload CalibrationPayload) {
	if calibrationID != nil {
		for _, calibrationValue := range payload.SignedValues {
			calibrationValueParams := NewCreateCalibrationValueParams(*calibrationID, calibrationValue)
			r.Repository.CreateCalibrationValue(ctx, calibrationValueParams)
		}
	}
}
