package cdr

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/pkg/ocpi"
	ocpiCdr "github.com/satimoto/go-ocpi/pkg/ocpi/cdr"
)

func (r *CdrResolver) ReplaceCdr(ctx context.Context, credential db.Credential, uid string, cdrDto *dto.CdrDto) *db.Cdr {
	if cdrDto != nil {
		countryCode, partyID := evse.GetEvsesIdentity(cdrDto.Location, cdrDto.Location.Evses)

		return r.ReplaceCdrByIdentifier(ctx, credential, countryCode, partyID, uid, cdrDto)
	}

	return nil
}

func (r *CdrResolver) ReplaceCdrByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, cdrDto *dto.CdrDto) *db.Cdr {
	if cdrDto != nil {
		cdr, err := r.Repository.GetCdrByUid(ctx, uid)

		if err == nil {
			log.Printf("Ignoring existing cdr %v", uid)
			return &cdr
		}

		cdrParams := NewCreateCdrParams(cdrDto)
		cdrParams.CountryCode = util.SqlNullString(countryCode)
		cdrParams.PartyID = util.SqlNullString(partyID)
		cdrParams.CredentialID = credential.ID

		if cdrDto.AuthID != nil {
			token, err := r.TokenRepository.GetTokenByAuthID(ctx, *cdrDto.AuthID)

			if err != nil {
				metrics.RecordError("OCPI020", "Error retrieving token", err)
				log.Printf("OCPI020: AuthID=%v", *cdrDto.AuthID)
				return nil
			}

			cdrParams.TokenID = token.ID
			cdrParams.UserID = token.UserID
		}

		if cdrDto.Location != nil {
			location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *cdrDto.Location.ID)

			if err != nil {
				metrics.RecordError("OCPI021", "Error retrieving location", err)
				log.Printf("OCPI021: Uid=%v", *cdrDto.Location.ID)
			} else {
				cdrParams.LocationID = location.ID
			}

			evseDto := cdrDto.Location.Evses[0]
			evse, err := r.LocationResolver.EvseResolver.Repository.GetEvseByUid(ctx, *evseDto.Uid)

			if err != nil {
				metrics.RecordError("OCPI022", "Error retrieving evse", err)
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
				metrics.RecordError("OCPI023", "Error retrieving connector", err)
				log.Printf("OCPI023: Params=%#v", connectorParams)
			} else {
				cdrParams.ConnectorID = connector.ID

				if (cdrDto.Tariffs == nil || len(cdrDto.Tariffs) == 0) && connector.TariffID.Valid {
					connectorTariff, err := r.TariffResolver.Repository.GetTariffByUid(ctx, connector.TariffID.String)

					if err != nil {
						metrics.RecordError("OCPI301", "Error getting connector tariff", err)
						log.Printf("OCPI301: ConnectorID=%v, Params=%#v", connector.ID, connector.TariffID)
					} else {
						tariffDo := r.TariffResolver.CreateTariffDto(ctx, connectorTariff)
						cdrDto.Tariffs = []*dto.TariffDto{tariffDo}
					}
				}
			}
		}

		if cdrDto.SignedData != nil {
			if calibration := r.CalibrationResolver.CreateCalibration(ctx, cdrDto.SignedData); calibration != nil {
				cdrParams.CalibrationID = util.SqlNullInt64(calibration.ID)
			}
		}

		cdr, err = r.Repository.CreateCdr(ctx, cdrParams)

		if err != nil {
			metrics.RecordError("OCPI024", "Error creating cdr", err)
			log.Printf("OCPI024: Params=%#v", cdrParams)
			return nil
		}

		if cdrDto.ChargingPeriods != nil {
			r.createChargingPeriods(ctx, cdr.ID, cdrDto)
		}

		if cdrDto.Tariffs != nil {
			r.replaceTariffs(ctx, credential, countryCode, partyID, &cdr.ID, cdrDto)
		}

		// If cdr received before session is completed, set status to completed
		if cdr.AuthorizationID.Valid {
			if session, err := r.SessionRepository.GetSessionByAuthorizationID(ctx, cdr.AuthorizationID.String); err == nil {
				// The session exists, set it to completed
				sessionParams := param.NewUpdateSessionByUidParams(session)
				sessionParams.Status = db.SessionStatusTypeCOMPLETED

				_, err = r.SessionRepository.UpdateSessionByUid(ctx, sessionParams)

				if err != nil {
					metrics.RecordError("OCPI283", "Error updating session", err)
					log.Printf("OCPI283: Params=%#v", sessionParams)
				}

				r.updateCommand(cdr, session)
			} else {
				// A session was never received for this cdr, create it
				createSessionParams := db.CreateSessionParams{
					Uid:             cdr.Uid,
					CredentialID:    cdr.CredentialID,
					CountryCode:     cdr.CountryCode,
					PartyID:         cdr.PartyID,
					AuthorizationID: cdr.AuthorizationID,
					StartDatetime:   cdr.StartDateTime,
					EndDatetime:     cdr.StopDateTime,
					Kwh:             cdr.TotalEnergy,
					AuthID:          cdr.AuthID,
					AuthMethod:      cdr.AuthMethod,
					UserID:          cdr.UserID,
					TokenID:         cdr.TokenID,
					LocationID:      cdr.LocationID,
					EvseID:          cdr.EvseID,
					ConnectorID:     cdr.ConnectorID,
					Currency:        cdr.Currency,
					TotalCost:       util.SqlNullFloat64(cdr.TotalCost),
					Status:          db.SessionStatusTypeCOMPLETED,
					LastUpdated:     cdr.LastUpdated,
				}

				_, err = r.SessionRepository.CreateSession(ctx, createSessionParams)

				if err != nil {
					metrics.RecordError("OCPI316", "Error creating session", err)
					log.Printf("OCPI316: Params=%#v", createSessionParams)
				}
			}
		}

		// Metrics
		metricCdrsTotal.Inc()

		go r.sendOcpiRequest(cdr)

		return &cdr
	}

	return nil
}

func (r *CdrResolver) ReplaceCdrs(ctx context.Context, credential db.Credential, cdrsDto []*dto.CdrDto) {
	for _, cdrDto := range cdrsDto {
		r.ReplaceCdr(ctx, credential, *cdrDto.ID, cdrDto)
	}
}

func (r *CdrResolver) ReplaceCdrsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, cdrsDto []*dto.CdrDto) {
	for _, cdrDto := range cdrsDto {
		r.ReplaceCdrByIdentifier(ctx, credential, countryCode, partyID, *cdrDto.ID, cdrDto)
	}
}

func (r *CdrResolver) createChargingPeriods(ctx context.Context, cdrID int64, cdrDto *dto.CdrDto) {
	for _, chargingPeriodDto := range cdrDto.ChargingPeriods {
		chargingPeriod := r.ChargingPeriodResolver.ReplaceChargingPeriod(ctx, chargingPeriodDto)

		if chargingPeriod != nil {
			setCdrChargingPeriodParams := db.SetCdrChargingPeriodParams{
				CdrID:            cdrID,
				ChargingPeriodID: chargingPeriod.ID,
			}
			err := r.Repository.SetCdrChargingPeriod(ctx, setCdrChargingPeriodParams)

			if err != nil {
				metrics.RecordError("OCPI027", "Error setting cdr charging period", err)
				log.Printf("OCPI027: Params=%#v", setCdrChargingPeriodParams)
			}
		}
	}
}

func (r *CdrResolver) replaceTariffs(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, cdrID *int64, cdrDto *dto.CdrDto) {
	for _, tariffDto := range cdrDto.Tariffs {
		r.TariffResolver.ReplaceTariffByIdentifier(ctx, credential, countryCode, partyID, *tariffDto.ID, cdrID, tariffDto)
	}
}

func (r *CdrResolver) sendOcpiRequest(cdr db.Cdr) {
	ctx := context.Background()

	// TODO: Handle failed RPC call more robustly
	node, err := r.NodeRepository.GetNodeByUserID(ctx, cdr.UserID)

	if err != nil {
		metrics.RecordError("OCPI025", "Error retrieving node", err)
		log.Printf("OCPI025: UserID=%v", cdr.UserID)
	}

	ocpiService := ocpi.NewService(node.RpcAddr)
	cdrCreatedRequest := ocpiCdr.NewCdrCreatedRequest(cdr)
	cdrCreatedResponse, err := ocpiService.CdrCreated(ctx, cdrCreatedRequest)

	if err != nil {
		metrics.RecordError("OCPI026", "Error calling RPC service", err)
		log.Printf("OCPI026: Request=%#v, Response=%#v", cdrCreatedRequest, cdrCreatedResponse)
	}
}
