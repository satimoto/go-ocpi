package cdr

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
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
	*chargingperiod.ChargingPeriodResolver
	*credential.CredentialResolver
	*location.LocationResolver
	*tariff.TariffResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CdrResolver {
	repo := CdrRepository(repositoryService)
	return &CdrResolver{
		Repository:             repo,
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		CredentialResolver:     credential.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService),
		TariffResolver:         tariff.NewResolver(repositoryService),
	}
}

func (r *CdrResolver) CreateCdr(ctx context.Context, payload *CdrPayload) *db.Cdr {
	if payload != nil {
		cdrParams := NewCreateCdrParams(payload)

		if location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *payload.Location.ID); err == nil {
			cdrParams.LocationID = location.ID
		} else {
			location := r.LocationResolver.ReplaceLocation(ctx, *payload.Location.ID, payload.Location)
			cdrParams.LocationID = location.ID
		}

		if cdr, err := r.Repository.CreateCdr(ctx, cdrParams); err == nil {
			if payload.ChargingPeriods != nil {
				r.createChargingPeriods(ctx, cdr.ID, payload)
			}

			if payload.Tariffs != nil {
				r.createTariffs(ctx, cdr.ID, payload)
			}

			return &cdr
		}
	}

	return nil
}

func (r *CdrResolver) createChargingPeriods(ctx context.Context, cdrID int64, payload *CdrPayload) {
	for _, chargingPeriodPayload := range payload.ChargingPeriods {
		chargingPeriod := r.ChargingPeriodResolver.ReplaceChargingPeriod(ctx, chargingPeriodPayload)

		if chargingPeriod != nil {
			r.Repository.SetCdrChargingPeriod(ctx, db.SetCdrChargingPeriodParams{
				CdrID:            cdrID,
				ChargingPeriodID: chargingPeriod.ID,
			})
		}
	}
}

func (r *CdrResolver) createTariffs(ctx context.Context, cdrID int64, payload *CdrPayload) {
	for _, tariffPayload := range payload.Tariffs {
		r.TariffResolver.ReplaceTariff(ctx, &cdrID, *tariffPayload.ID, tariffPayload)
	}
}
