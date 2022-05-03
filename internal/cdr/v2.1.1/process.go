package cdr

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
)

func (r *CdrResolver) ReplaceCdr(ctx context.Context, credential db.Credential, dto *CdrDto) *db.Cdr {
	if dto != nil {
		countryCode, partyID := evse.GetEvsesIdentity(dto.Location.Evses)

		return r.ReplaceCdrByIdentifier(ctx, credential, countryCode, partyID, *dto.ID, dto)
	}

	return nil
}

func (r *CdrResolver) ReplaceCdrByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, dto *CdrDto) *db.Cdr {
	if dto != nil {
		cdr, err := r.Repository.GetCdrByUid(ctx, uid)

		if err != nil {
			cdrParams := NewCreateCdrParams(dto)
			cdrParams.CountryCode = util.SqlNullString(countryCode)
			cdrParams.PartyID = util.SqlNullString(partyID)
			cdrParams.CredentialID = credential.ID

			if dto.AuthID != nil {
				if token, err := r.TokenResolver.Repository.GetTokenByAuthID(ctx, *dto.AuthID); err == nil {
					cdrParams.TokenID = token.ID
					cdrParams.UserID = token.UserID
				}
			}
			if dto.Location != nil {
				if location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *dto.Location.ID); err == nil {
					cdrParams.LocationID = location.ID
				}
	
				evseDto := dto.Location.Evses[0]
	
				if evse, err := r.LocationResolver.EvseResolver.Repository.GetEvseByUid(ctx, *evseDto.Uid); err == nil {
					cdrParams.EvseID = evse.ID
				}
	
				connectorDto := evseDto.Connectors[0]
				connectorParams := db.GetConnectorByUidParams{
					EvseID: cdrParams.EvseID,
					Uid: *connectorDto.Id,
				}
	
				if connector, err := r.LocationResolver.ConnectorResolver.Repository.GetConnectorByUid(ctx, connectorParams); err == nil {
					cdrParams.ConnectorID = connector.ID
				}
			}

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
					r.replaceTariffs(ctx, credential, countryCode, partyID, &cdr.ID, dto)
				}
			}
 
			// TODO: Send CdrCreated RPC to LSP node
		}

		return &cdr
	}

	return nil
}

func (r *CdrResolver) ReplaceCdrsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, dto []*CdrDto) {
	for _, cdrDto := range dto {
		r.ReplaceCdrByIdentifier(ctx, credential, countryCode, partyID, *cdrDto.ID, cdrDto)
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

func (r *CdrResolver) replaceTariffs(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, cdrID *int64, dto *CdrDto) {
	for _, tariffDto := range dto.Tariffs {
		r.TariffResolver.ReplaceTariffByIdentifier(ctx, credential, countryCode, partyID, *tariffDto.ID, cdrID, tariffDto)
	}
}
