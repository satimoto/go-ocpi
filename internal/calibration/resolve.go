package calibration

import (
	"github.com/satimoto/go-datastore/pkg/calibration"
	"github.com/satimoto/go-datastore/pkg/db"
)

type CalibrationResolver struct {
	Repository calibration.CalibrationRepository
}

func NewResolver(repositoryService *db.RepositoryService) *CalibrationResolver {
	return &CalibrationResolver{
		Repository: calibration.NewRepository(repositoryService),
	}
}
