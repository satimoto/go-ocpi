package session

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *SessionResolver) ReplaceSession(ctx context.Context, credential db.Credential, uid string, sessionDto *dto.SessionDto) *db.Session {
	if sessionDto != nil {
		countryCode, partyID := evse.GetEvsesIdentity(sessionDto.Location, sessionDto.Location.Evses)

		return r.ReplaceSessionByIdentifier(ctx, credential, countryCode, partyID, uid, sessionDto)
	}

	return nil
}

func (r *SessionResolver) ReplaceSessionByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, sessionDto *dto.SessionDto) *db.Session {
	if sessionDto != nil {
		session, err := r.Repository.GetSessionByUid(ctx, uid)
		sessionCreated := false
		statusChanged := false

		if err != nil && sessionDto.AuthorizationID != nil {
			// Check there is no existing session with the same AuthorizationID
			// This would have been created if no session was initially received
			session, err = r.Repository.GetSessionByAuthorizationID(ctx, *sessionDto.AuthorizationID)
			sessionDto.ID = &uid
		}

		if err == nil {
			sessionParams := param.NewUpdateSessionByUidParams(session)

			if sessionDto.ID != nil {
				sessionParams.Uid = *sessionDto.ID
			}

			if sessionDto.AuthMethod != nil {
				sessionParams.AuthMethod = *sessionDto.AuthMethod
			}

			if sessionDto.Currency != nil {
				sessionParams.Currency = *sessionDto.Currency
			}

			if sessionDto.EndDatetime != nil {
				sessionParams.EndDatetime = util.SqlNullTime(sessionDto.EndDatetime)
			}

			if sessionDto.Kwh != nil {
				sessionParams.Kwh = *sessionDto.Kwh
			}

			if sessionDto.LastUpdated != nil {
				sessionParams.LastUpdated = sessionDto.LastUpdated.Time()
			}

			if sessionDto.MeterID != nil {
				sessionParams.MeterID = util.SqlNullString(sessionDto.MeterID)
			}

			if sessionDto.Status != nil && session.Status != db.SessionStatusTypeINVOICED && session.Status != *sessionDto.Status {
				if session.Status == db.SessionStatusTypeINVALID && *sessionDto.Status == db.SessionStatusTypeACTIVE {
					// Session has already been set to INVALID
					log.Printf("Session is already INVALID, not setting it to %v", *sessionDto.Status)

					if token, err := r.TokenRepository.GetToken(ctx, session.TokenID); err == nil && token.Type == db.TokenTypeOTHER {
						_, err := r.CommandResolver.StopSession(ctx, credential, session.Uid)

						if err != nil {
							metrics.RecordError("OCPI310", "Error stopping session", err)
							log.Printf("OCPI310: SessionUid=%#v", session.Uid)
						}
					}
				} else {
					statusChanged = true
					sessionParams.Status = *sessionDto.Status
				}
			}

			if sessionDto.StartDatetime != nil {
				sessionParams.StartDatetime = sessionDto.StartDatetime.Time()
			}

			if sessionDto.TotalCost != nil {
				sessionParams.TotalCost = util.SqlNullFloat64(sessionDto.TotalCost)
			}

			updatedSession, err := r.Repository.UpdateSessionByUid(ctx, sessionParams)

			if err != nil {
				metrics.RecordError("OCPI161", "Error updating session", err)
				log.Printf("OCPI161: Params=%#v", sessionParams)
				return nil
			}

			session = updatedSession
		} else {
			sessionCreated = true
			sessionParams := NewCreateSessionParams(sessionDto)
			sessionParams.CredentialID = credential.ID
			sessionParams.CountryCode = util.SqlNullString(countryCode)
			sessionParams.PartyID = util.SqlNullString(partyID)

			if sessionDto.AuthID != nil {
				token, err := r.TokenRepository.GetTokenByAuthID(ctx, *sessionDto.AuthID)

				if err != nil {
					metrics.RecordError("OCPI162", "Error retrieving token", err)
					log.Printf("OCPI162: AuthID=%v", *sessionDto.AuthID)
					return nil
				}

				sessionParams.TokenID = token.ID
				sessionParams.UserID = token.UserID
			}

			if sessionDto.Location != nil {
				location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *sessionDto.Location.ID)

				if err != nil {
					metrics.RecordError("OCPI163", "Error retrieving location", err)
					log.Printf("OCPI163: Uid=%v", *sessionDto.Location.ID)
				} else {
					sessionParams.LocationID = location.ID
				}

				evseDto := sessionDto.Location.Evses[0]
				evse, err := r.LocationResolver.EvseResolver.Repository.GetEvseByUid(ctx, *evseDto.Uid)

				if err != nil {
					metrics.RecordError("OCPI164", "Error retrieving evse", err)
					log.Printf("OCPI164: Uid=%v", *evseDto.Uid)
				} else {
					sessionParams.EvseID = evse.ID
				}

				connectorDto := evseDto.Connectors[0]
				connectorParams := db.GetConnectorByEvseParams{
					EvseID: sessionParams.EvseID,
					Uid:    *connectorDto.Id,
				}
				connector, err := r.LocationResolver.EvseResolver.ConnectorResolver.Repository.GetConnectorByEvse(ctx, connectorParams)

				if err != nil {
					metrics.RecordError("OCPI165", "Error retrieving connector", err)
					log.Printf("OCPI165: Params=%#v", connectorParams)
				} else {
					sessionParams.ConnectorID = connector.ID
				}
			}

			session, err = r.Repository.CreateSession(ctx, sessionParams)

			if err != nil {
				metrics.RecordError("OCPI166", "Error creating session", err)
				log.Printf("OCPI166: Params=%#v", sessionParams)
				return nil
			}

			if sessionDto.AuthorizationID != nil {
				r.replaceTokenAuthorization(ctx, countryCode, partyID, sessionDto)
			}
		}

		if sessionDto.ChargingPeriods != nil {
			r.replaceChargingPeriods(ctx, session.ID, sessionDto)
		}

		// Metrics
		go r.updateMetrics(session, sessionCreated)

		// Send a session created/update RPC message to LSP
		go r.sendOcpiRequest(session, sessionCreated, statusChanged)

		if sessionCreated || statusChanged {
			go r.updateCommand(session)
		}

		if sessionCreated && session.Status == db.SessionStatusTypePENDING {
			go r.waitForEvseStatus(credential, session.LocationID, session.EvseID, db.EvseStatusCHARGING, session, session.Status, db.SessionStatusTypeACTIVE, 150)
		}

		return &session
	}

	return nil
}

func (r *SessionResolver) ReplaceSessions(ctx context.Context, credential db.Credential, sessionsDto []*dto.SessionDto) {
	for _, sessionDto := range sessionsDto {
		r.ReplaceSession(ctx, credential, *sessionDto.ID, sessionDto)
	}
}

func (r *SessionResolver) ReplaceSessionsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, sessionsDto []*dto.SessionDto) {
	for _, sessionDto := range sessionsDto {
		r.ReplaceSessionByIdentifier(ctx, credential, countryCode, partyID, *sessionDto.ID, sessionDto)
	}
}

func (r *SessionResolver) replaceChargingPeriods(ctx context.Context, sessionID int64, sessionDto *dto.SessionDto) {
	r.Repository.DeleteSessionChargingPeriods(ctx, sessionID)

	for _, chargingPeriodDto := range sessionDto.ChargingPeriods {
		chargingPeriod := r.ChargingPeriodResolver.ReplaceChargingPeriod(ctx, chargingPeriodDto)

		if chargingPeriod != nil {
			setSessionChargingPeriodParams := db.SetSessionChargingPeriodParams{
				SessionID:        sessionID,
				ChargingPeriodID: chargingPeriod.ID,
			}
			err := r.Repository.SetSessionChargingPeriod(ctx, setSessionChargingPeriodParams)

			if err != nil {
				metrics.RecordError("OCPI169", "Error setting session charging period", err)
				log.Printf("OCPI169: Params=%#v", setSessionChargingPeriodParams)
			}
		}
	}
}

func (r *SessionResolver) replaceTokenAuthorization(ctx context.Context, countryCode *string, partyID *string, sessionDto *dto.SessionDto) {
	tokenAuthorization, err := r.TokenAuthorizationRepository.GetTokenAuthorizationByAuthorizationID(ctx, *sessionDto.AuthorizationID)

	if err != nil {
		metrics.RecordError("OCPI209", "Error retrieving token authorization", err)
		log.Printf("OCPI209: AuthorizationID=%v", *sessionDto.AuthorizationID)
		return
	}

	tokenAuthorizationParams := param.NewUpdateTokenAuthorizationParams(tokenAuthorization)
	tokenAuthorizationParams.CountryCode = util.SqlNullString(countryCode)
	tokenAuthorizationParams.PartyID = util.SqlNullString(partyID)

	r.TokenAuthorizationRepository.UpdateTokenAuthorizationByAuthorizationID(ctx, tokenAuthorizationParams)
}
