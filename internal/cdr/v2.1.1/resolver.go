package cdr

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/calibration"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type CdrRepository interface {
	CreateCdr(ctx context.Context, arg db.CreateCdrParams) (db.Cdr, error)
	DeleteCdrChargingPeriods(ctx context.Context, cdrID int64) error
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
}

func NewResolver(repositoryService *db.RepositoryService) *CdrResolver {
	repo := CdrRepository(repositoryService)
	return &CdrResolver{
		Repository:             repo,
		CalibrationResolver:    calibration.NewResolver(repositoryService),
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService),
		TariffResolver:         tariff.NewResolver(repositoryService),
	}
}

func (r *CdrResolver) CreateCdr(ctx context.Context, dto *CdrDto) *db.Cdr {
	if dto != nil {
		cdrParams := NewCreateCdrParams(dto)

		if location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *dto.Location.ID); err == nil {
			cdrParams.LocationID = location.ID
		} else {
			location := r.LocationResolver.ReplaceLocation(ctx, *dto.Location.ID, dto.Location)
			cdrParams.LocationID = location.ID
		}

		if dto.SignedData != nil {
			if calibration := r.CalibrationResolver.CreateCalibration(ctx, *&dto.SignedData); calibration != nil {
				cdrParams.CalibrationID = util.SqlNullInt64(calibration.ID)
			}
		}

		if cdr, err := r.Repository.CreateCdr(ctx, cdrParams); err == nil {
			if dto.ChargingPeriods != nil {
				r.createChargingPeriods(ctx, cdr.ID, dto)
			}

			if dto.Tariffs != nil {
				r.createTariffs(ctx, cdr.ID, dto)
			}

			return &cdr
		}
	}

	return nil
}

func (r *CdrResolver) createChargingPeriods(ctx context.Context, cdrID int64, dto *CdrDto) {
	for _, chargingPeriodDto := range dto.ChargingPeriods {
		chargingPeriod := r.ChargingPeriodResolver.ReplaceChargingPeriod(ctx, chargingPeriodDto)

		if chargingPeriod != nil {
			r.Repository.SetCdrChargingPeriod(ctx, db.SetCdrChargingPeriodParams{
				CdrID:            cdrID,
				ChargingPeriodID: chargingPeriod.ID,
			})
		}
	}
}

func (r *CdrResolver) createTariffs(ctx context.Context, cdrID int64, dto *CdrDto) {
	for _, tariffDto := range dto.Tariffs {
		r.TariffResolver.ReplaceTariff(ctx, &cdrID, *tariffDto.ID, tariffDto)
	}
}
