package cdr

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
)

func (r *CdrResolver) ReplaceCdr(ctx context.Context, dto *CdrDto) *db.Cdr {
	if dto != nil {
		countryCode, partyID := evse.GetEvsesIdentity(dto.Location.Evses)

		return r.ReplaceCdrByIdentifier(ctx, countryCode, partyID, *dto.ID, dto)
	}

	return nil
}

func (r *CdrResolver) ReplaceCdrByIdentifier(ctx context.Context, countryCode *string, partyID *string, uid string, dto *CdrDto) *db.Cdr {
	if dto != nil {
		cdr, err := r.Repository.GetCdrByUid(ctx, uid)

		if err != nil {
			var locationID int64

			cdrParams := NewCreateCdrParams(dto)

			if location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *dto.Location.ID); err == nil {
				countryCode = util.NilString(location.CountryCode)
				partyID = util.NilString(location.PartyID)
				locationID = location.ID
			} else {
				location := r.LocationResolver.ReplaceLocation(ctx, *dto.Location.ID, dto.Location)
				countryCode = util.NilString(location.CountryCode)
				partyID = util.NilString(location.PartyID)
				locationID = location.ID
			}

			cdrParams.CountryCode = util.SqlNullString(countryCode)
			cdrParams.PartyID = util.SqlNullString(partyID)
			cdrParams.LocationID = locationID

			if dto.SignedData != nil {
				if calibration := r.CalibrationResolver.CreateCalibration(ctx, *&dto.SignedData); calibration != nil {
					cdrParams.CalibrationID = util.SqlNullInt64(calibration.ID)
				}
			}

			cdr, err = r.Repository.CreateCdr(ctx, cdrParams)

			if err == nil {
				if dto.ChargingPeriods != nil {
					r.createChargingPeriods(ctx, cdr.ID, dto)
				}

				if dto.Tariffs != nil {
					r.replaceTariffs(ctx, countryCode, partyID, &cdr.ID, dto)
				}
			}
		}

		return &cdr
	}

	return nil
}

func (r *CdrResolver) ReplaceCdrsByIdentifier(ctx context.Context, countryCode *string, partyID *string, dto []*CdrDto) {
	for _, cdrDto := range dto {
		r.ReplaceCdrByIdentifier(ctx, countryCode, partyID, *cdrDto.ID, cdrDto)
	}
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

func (r *CdrResolver) replaceTariffs(ctx context.Context, countryCode *string, partyID *string, cdrID *int64, dto *CdrDto) {
	for _, tariffDto := range dto.Tariffs {
		r.TariffResolver.ReplaceTariffByIdentifier(ctx, countryCode, partyID, *tariffDto.ID, cdrID, tariffDto)
	}
}
