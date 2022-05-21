package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	calibrationMocks "github.com/satimoto/go-datastore/pkg/calibration/mocks"
	"github.com/satimoto/go-ocpi-api/internal/calibration"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *calibration.CalibrationResolver {
	return &calibration.CalibrationResolver{
		Repository: calibrationMocks.NewRepository(repositoryService),
	}
}
