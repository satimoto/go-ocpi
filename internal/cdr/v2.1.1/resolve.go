package cdr

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/calibration"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
)

type CdrRepository interface {
	CreateCdr(ctx context.Context, arg db.CreateCdrParams) (db.Cdr, error)
	DeleteCdrChargingPeriods(ctx context.Context, cdrID int64) error
	GetCdrByIdentityOrderByLastUpdated(ctx context.Context, arg db.GetCdrByIdentityOrderByLastUpdatedParams) (db.Cdr, error)
	GetCdrByUid(ctx context.Context, uid string) (db.Cdr, error)
	ListCdrChargingPeriods(ctx context.Context, cdrID int64) ([]db.ChargingPeriod, error)
	SetCdrChargingPeriod(ctx context.Context, arg db.SetCdrChargingPeriodParams) error
}

type CdrResolver struct {
	Repository CdrRepository
	*calibration.CalibrationResolver
	*chargingperiod.ChargingPeriodResolver
	*location.LocationResolver
	*tariff.TariffResolver
	*versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CdrResolver {
	repo := CdrRepository(repositoryService)
	return &CdrResolver{
		Repository:             repo,
		CalibrationResolver:    calibration.NewResolver(repositoryService),
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService),
		TariffResolver:         tariff.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService),
	}
}
