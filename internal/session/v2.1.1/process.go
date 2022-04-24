package session

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-datastore/util"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1"
)

func (r *SessionResolver) ReplaceSessionByIdentifier(ctx context.Context, credential db.Credential, countryCode *string, partyID *string, uid string, dto *SessionDto) *db.Session {
	if dto != nil {
		session, err := r.Repository.GetSessionByUid(ctx, uid)

		if err == nil {
			sessionParams := NewUpdateSessionByUidParams(session)

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

			session, err = r.Repository.UpdateSessionByUid(ctx, sessionParams)
		} else {
			sessionParams := NewCreateSessionParams(dto)
			sessionParams.CredentialID = credential.ID
			sessionParams.CountryCode = util.SqlNullString(countryCode)
			sessionParams.PartyID = util.SqlNullString(partyID)
	
			if dto.AuthID != nil {
				if token, err := r.TokenResolver.Repository.GetTokenByAuthID(ctx, *dto.AuthID); err == nil {
					sessionParams.TokenID = token.ID
				}
			}

			if dto.Location != nil {
				if location, err := r.LocationResolver.Repository.GetLocationByUid(ctx, *dto.Location.ID); err == nil {
					sessionParams.LocationID = location.ID
				}
	
				evseDto := dto.Location.Evses[0]
	
				if evse, err := r.LocationResolver.EvseResolver.Repository.GetEvseByUid(ctx, *evseDto.Uid); err == nil {
					sessionParams.EvseID = evse.ID
				}
	
				connectorDto := evseDto.Connectors[0]
				connectorParams := db.GetConnectorByUidParams{
					EvseID: sessionParams.EvseID,
					Uid: *connectorDto.Id,
				}
	
				if connector, err := r.LocationResolver.ConnectorResolver.Repository.GetConnectorByUid(ctx, connectorParams); err == nil {
					sessionParams.ConnectorID = connector.ID
				}
			}

			session, err = r.Repository.CreateSession(ctx, sessionParams)
		}

		if dto.AuthorizationID != nil {
			r.replaceTokenAuthorization(ctx, countryCode, partyID, dto)
		}

		if dto.ChargingPeriods != nil {
			r.replaceChargingPeriods(ctx, session.ID, dto)
		}

		return &session
	}

	return nil
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
			r.Repository.SetSessionChargingPeriod(ctx, db.SetSessionChargingPeriodParams{
				SessionID:        sessionID,
				ChargingPeriodID: chargingPeriod.ID,
			})
		}
	}
}

func (r *SessionResolver) replaceTokenAuthorization(ctx context.Context, countryCode *string, partyID *string, dto *SessionDto) {
	tokenAuthorizationParams := tokenauthorization.NewUpdateTokenAuthorizationParams(*dto.AuthorizationID, countryCode, partyID)
	r.TokenAuthorizationResolver.Repository.UpdateTokenAuthorizationByAuthorizationID(ctx, tokenAuthorizationParams)
}
