package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi-api/internal/calibration"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *calibration.CalibrationResolver {
	repo := calibration.CalibrationRepository(repositoryService)

	return &calibration.CalibrationResolver{
		Repository: repo,
	}
}
