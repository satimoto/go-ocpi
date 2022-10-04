package cdr

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi/pkg/ocpi"
	ocpiCdr "github.com/satimoto/go-ocpi/pkg/ocpi/cdr"
)

func (r *CdrResolver) ReplaceCdr(ctx context.Context, credential db.Credential, uid string, dto *CdrDto) *db.Cdr {
	if dto != nil {
		countryCode, partyID := evse.GetEvsesIdentity(dto.Location.Evses)

		return r.ReplaceCdrByIdentifier(ctx, credential, countryCode, partyID, uid, dto)
	}

	return nil
}

func (r *CdrResolver) ReplaceCdrByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, dto *CdrDto) *db.Cdr {
	if dto != nil {
		cdr, err := r.Repository.GetCdrByUid(ctx, uid)

		if err == nil {
			util.LogOnError("OCPI019", "Error cdr exists", err)
			log.Printf("OCPI019: Uid=%v", uid)
			return nil
		}

		cdrParams := NewCreateCdrParams(dto)
		cdrParams.CountryCode = util.SqlNullString(countryCode)
		cdrParams.PartyID = util.SqlNullString(partyID)
		cdrParams.CredentialID = credential.ID

		if dto.AuthID != nil {
			token, err := r.TokenRepository.GetTokenByAuthID(ctx, *dto.AuthID)

			if err != nil {
				util.LogOnError("OCPI020", "Error retrieving token", err)
				log.Printf("OCPI020: AuthID=%v", *dto.AuthID)
				return nil
			}

			cdrParams.TokenID = token.ID
			cdrParams.UserID = token.UserID
		}

		if dto.Location != nil {
			location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *dto.Location.ID)

			if err != nil {
				util.LogOnError("OCPI021", "Error retrieving location", err)
				log.Printf("OCPI021: Uid=%v", *dto.Location.ID)
			} else {
				cdrParams.LocationID = location.ID
			}

			evseDto := dto.Location.Evses[0]
			evse, err := r.LocationResolver.EvseResolver.Repository.GetEvseByUid(ctx, *evseDto.Uid)

			if err != nil {
				util.LogOnError("OCPI022", "Error retrieving evse", err)
				log.Printf("OCPI022: Uid=%v", *evseDto.Uid)
			} else {
				cdrParams.EvseID = evse.ID
			}

			connectorDto := evseDto.Connectors[0]
			connectorParams := db.GetConnectorByEvseParams{
				EvseID: cdrParams.EvseID,
				Uid:    *connectorDto.Id,
			}
			connector, err := r.LocationResolver.EvseResolver.ConnectorResolver.Repository.GetConnectorByEvse(ctx, connectorParams)

			if err != nil {
				util.LogOnError("OCPI023", "Error retrieving connector", err)
				log.Printf("OCPI023: Params=%#v", connectorParams)
			} else {
				cdrParams.ConnectorID = connector.ID
			}
		}

		if dto.SignedData != nil {
			if calibration := r.CalibrationResolver.CreateCalibration(ctx, dto.SignedData); calibration != nil {
				cdrParams.CalibrationID = util.SqlNullInt64(calibration.ID)
			}
		}

		cdr, err = r.Repository.CreateCdr(ctx, cdrParams)

		if err != nil {
			util.LogOnError("OCPI024", "Error creating cdr", err)
			log.Printf("OCPI024: Params=%#v", cdrParams)
			return nil
		}

		if dto.ChargingPeriods != nil {
			r.createChargingPeriods(ctx, cdr.ID, dto)
		}

		if dto.Tariffs != nil {
			r.replaceTariffs(ctx, credential, countryCode, partyID, &cdr.ID, dto)
		}

		node, err := r.NodeRepository.GetNodeByUserID(ctx, cdr.UserID)

		if err != nil {
			util.LogOnError("OCPI025", "Error retrieving node", err)
			log.Printf("OCPI025: UserID=%v", cdr.UserID)
			return &cdr
		}

		// If cdr received before session is completed, set status to completed
		if cdr.AuthorizationID.Valid {
			if session, err := r.SessionRepository.GetSessionByAuthorizationID(ctx, cdr.AuthorizationID.String); err == nil {
				sessionParams := param.NewUpdateSessionByUidParams(session)
				sessionParams.Status = db.SessionStatusTypeCOMPLETED

				_, err = r.SessionRepository.UpdateSessionByUid(ctx, sessionParams)

				if err != nil {
					util.LogOnError("OCPI283", "Error updating session", err)
					log.Printf("OCPI283: Params=%#v", sessionParams)
				}
			}
		}

		// TODO: Handle failed RPC call more robustly
		ocpiService := ocpi.NewService(node.LspAddr)
		cdrCreatedRequest := ocpiCdr.NewCdrCreatedRequest(cdr)
		cdrCreatedResponse, err := ocpiService.CdrCreated(ctx, cdrCreatedRequest)

		if err != nil {
			util.LogOnError("OCPI026", "Error calling RPC service", err)
			log.Printf("OCPI026: Request=%#v, Response=%#v", cdrCreatedRequest, cdrCreatedResponse)
		}

		return &cdr
	}

	return nil
}

func (r *CdrResolver) ReplaceCdrs(ctx context.Context, credential db.Credential, dto []*CdrDto) {
	for _, cdrDto := range dto {
		r.ReplaceCdr(ctx, credential, *cdrDto.ID, cdrDto)
	}
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
			setCdrChargingPeriodParams := db.SetCdrChargingPeriodParams{
				CdrID:            cdrID,
				ChargingPeriodID: chargingPeriod.ID,
			}
			err := r.Repository.SetCdrChargingPeriod(ctx, setCdrChargingPeriodParams)

			if err != nil {
				util.LogOnError("OCPI027", "Error setting cdr charging period", err)
				log.Printf("OCPI027: Params=%#v", setCdrChargingPeriodParams)
			}
		}
	}
}

func (r *CdrResolver) replaceTariffs(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, cdrID *int64, dto *CdrDto) {
	for _, tariffDto := range dto.Tariffs {
		r.TariffResolver.ReplaceTariffByIdentifier(ctx, credential, countryCode, partyID, *tariffDto.ID, cdrID, tariffDto)
	}
}
