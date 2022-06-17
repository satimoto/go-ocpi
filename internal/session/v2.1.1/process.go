package session

import (
	"context"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	evse "github.com/satimoto/go-ocpi-api/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi-api/pkg/ocpi"
	ocpiSession "github.com/satimoto/go-ocpi-api/pkg/ocpi/session"
)

func (r *SessionResolver) ReplaceSession(ctx context.Context, credential db.Credential, uid string, dto *SessionDto) *db.Session {
	if dto != nil {
		countryCode, partyID := evse.GetEvsesIdentity(dto.Location.Evses)

		return r.ReplaceSessionByIdentifier(ctx, credential, countryCode, partyID, uid, dto)
	}

	return nil
}

func (r *SessionResolver) ReplaceSessionByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, dto *SessionDto) *db.Session {
	if dto != nil {
		session, err := r.Repository.GetSessionByUid(ctx, uid)

		if err == nil {
			sessionParams := param.NewUpdateSessionByUidParams(session)

			if dto.AuthMethod != nil {
				sessionParams.AuthMethod = *dto.AuthMethod
			}

			if dto.Currency != nil {
				sessionParams.Currency = *dto.Currency
			}

			if dto.EndDatetime != nil {
				sessionParams.EndDatetime = util.SqlNullTime(dto.EndDatetime)
			}

			if dto.Kwh != nil {
				sessionParams.Kwh = *dto.Kwh
			}

			if dto.LastUpdated != nil {
				sessionParams.LastUpdated = *dto.LastUpdated
			}

			if dto.MeterID != nil {
				sessionParams.MeterID = util.SqlNullString(dto.MeterID)
			}

			if dto.Status != nil {
				sessionParams.Status = *dto.Status
			}

			if dto.StartDatetime != nil {
				sessionParams.StartDatetime = *dto.StartDatetime
			}

			if dto.TotalCost != nil {
				sessionParams.TotalCost = util.SqlNullFloat64(dto.TotalCost)
			}

			updatedSession, err := r.Repository.UpdateSessionByUid(ctx, sessionParams)

			if err != nil {
				util.LogOnError("OCPI161", "Error updating session", err)
				log.Printf("OCPI161: Params=%#v", sessionParams)
				return nil
			}

			session = updatedSession
		} else {
			sessionParams := NewCreateSessionParams(dto)
			sessionParams.CredentialID = credential.ID
			sessionParams.CountryCode = util.SqlNullString(countryCode)
			sessionParams.PartyID = util.SqlNullString(partyID)

			if dto.AuthID != nil {
				token, err := r.TokenResolver.Repository.GetTokenByAuthID(ctx, *dto.AuthID)

				if err != nil {
					util.LogOnError("OCPI162", "Error retrieving token", err)
					log.Printf("OCPI162: AuthID=%v", *dto.AuthID)
					return nil
				}

				sessionParams.TokenID = token.ID
				sessionParams.UserID = token.UserID
			}

			if dto.Location != nil {
				location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *dto.Location.ID)

				if err != nil {
					util.LogOnError("OCPI163", "Error retrieving location", err)
					log.Printf("OCPI163: Uid=%v", *dto.Location.ID)
				} else {
					sessionParams.LocationID = location.ID
				}

				evseDto := dto.Location.Evses[0]
				evse, err := r.LocationResolver.EvseResolver.Repository.GetEvseByUid(ctx, *evseDto.Uid)

				if err != nil {
					util.LogOnError("OCPI164", "Error retrieving evse", err)
					log.Printf("OCPI164: Uid=%v", *evseDto.Uid)
				} else {
					sessionParams.EvseID = evse.ID
				}

				connectorDto := evseDto.Connectors[0]
				connectorParams := db.GetConnectorByUidParams{
					EvseID: sessionParams.EvseID,
					Uid:    *connectorDto.Id,
				}
				connector, err := r.LocationResolver.EvseResolver.ConnectorResolver.Repository.GetConnectorByUid(ctx, connectorParams)

				if err != nil {
					util.LogOnError("OCPI165", "Error retrieving connector", err)
					log.Printf("OCPI165: Params=%#v", connectorParams)
				} else {
					sessionParams.ConnectorID = connector.ID
				}
			}

			session, err = r.Repository.CreateSession(ctx, sessionParams)

			if err != nil {
				util.LogOnError("OCPI166", "Error creating session", err)
				log.Printf("OCPI166: Params=%#v", sessionParams)
				return nil
			}
		}

		if dto.AuthorizationID != nil {
			r.replaceTokenAuthorization(ctx, countryCode, partyID, dto)
		}

		if dto.ChargingPeriods != nil {
			r.replaceChargingPeriods(ctx, session.ID, dto)
		}

		node, err := r.NodeRepository.GetNodeByUserID(ctx, session.UserID)

		if err != nil {
			util.LogOnError("OCPI167", "Error retrieving node", err)
			log.Printf("OCPI167: UserID=%v", session.UserID)
			return &session
		}

		// TODO: Handle failed RPC call more robustly
		ocpiService := ocpi.NewService(node.LspAddr)
		sessionCreatedRequest := ocpiSession.NewSessionCreatedRequest(session)
		sessionCreatedResponse, err := ocpiService.SessionCreated(ctx, sessionCreatedRequest)

		if err != nil {
			util.LogOnError("OCPI168", "Error calling RPC service", err)
			log.Printf("OCPI168: Request=%#v, Response=%#v", sessionCreatedRequest, sessionCreatedResponse)
		}

		return &session
	}

	return nil
}

func (r *SessionResolver) ReplaceSessions(ctx context.Context, credential db.Credential, dto []*SessionDto) {
	for _, sessionDto := range dto {
		r.ReplaceSession(ctx, credential, *sessionDto.ID, sessionDto)
	}
}

func (r *SessionResolver) ReplaceSessionsByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, dto []*SessionDto) {
	for _, sessionDto := range dto {
		r.ReplaceSessionByIdentifier(ctx, credential, countryCode, partyID, *sessionDto.ID, sessionDto)
	}
}

func (r *SessionResolver) replaceChargingPeriods(ctx context.Context, sessionID int64, dto *SessionDto) {
	r.Repository.DeleteSessionChargingPeriods(ctx, sessionID)

	for _, chargingPeriodDto := range dto.ChargingPeriods {
		chargingPeriod := r.ChargingPeriodResolver.ReplaceChargingPeriod(ctx, chargingPeriodDto)

		if chargingPeriod != nil {
			setSessionChargingPeriodParams := db.SetSessionChargingPeriodParams{
				SessionID:        sessionID,
				ChargingPeriodID: chargingPeriod.ID,
			}
			err := r.Repository.SetSessionChargingPeriod(ctx, setSessionChargingPeriodParams)

			if err != nil {
				util.LogOnError("OCPI169", "Error setting session charging period", err)
				log.Printf("OCPI169: Params=%#v", setSessionChargingPeriodParams)
			}
		}
	}
}

func (r *SessionResolver) replaceTokenAuthorization(ctx context.Context, countryCode *string, partyID *string, dto *SessionDto) {
	tokenAuthorizationParams := param.NewUpdateTokenAuthorizationParams(*dto.AuthorizationID, countryCode, partyID)
	r.TokenAuthorizationResolver.Repository.UpdateTokenAuthorizationByAuthorizationID(ctx, tokenAuthorizationParams)
}
