package mocks

import (
	calibrationMocks "github.com/satimoto/go-datastore/pkg/calibration/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi/internal/calibration"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *calibration.CalibrationResolver {
	return &calibration.CalibrationResolver{
		Repository: calibrationMocks.NewRepository(repositoryService),
	}
}
