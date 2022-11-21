package session

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	coreLocation "github.com/satimoto/go-ocpi/internal/location"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/pkg/ocpi"
	ocpiSession "github.com/satimoto/go-ocpi/pkg/ocpi/session"
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

		if err == nil {
			sessionParams := param.NewUpdateSessionByUidParams(session)

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

			if sessionDto.Status != nil && session.Status != *sessionDto.Status {
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

		if sessionCreated && session.Status == db.SessionStatusTypePENDING {
			go r.waitForEvseStatus(credential, session.LocationID, session.EvseID, db.EvseStatusCHARGING, session.ID, session.Status, db.SessionStatusTypeACTIVE, 150)
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

func (r *SessionResolver) sendOcpiRequest(session db.Session, sessionCreated, statusChanged bool) {
	ctx := context.Background()
	node, err := r.NodeRepository.GetNodeByUserID(ctx, session.UserID)

	if err != nil {
		metrics.RecordError("OCPI167", "Error retrieving node", err)
		log.Printf("OCPI167: UserID=%v", session.UserID)
	}

	// TODO: Handle failed RPC call more robustly
	ocpiService := ocpi.NewService(node.LspAddr)

	if sessionCreated {
		sessionCreatedRequest := ocpiSession.NewSessionCreatedRequest(session)
		sessionCreatedResponse, err := ocpiService.SessionCreated(ctx, sessionCreatedRequest)

		if err != nil {
			metrics.RecordError("OCPI168", "Error calling RPC service", err)
			log.Printf("OCPI168: Request=%#v, Response=%#v", sessionCreatedRequest, sessionCreatedResponse)
		}
	} else if statusChanged {
		sessionUpdatedRequest := ocpiSession.NewSessionUpdatedRequest(session)
		sessionUpdatedResponse, err := ocpiService.SessionUpdated(ctx, sessionUpdatedRequest)

		if err != nil {
			metrics.RecordError("OCPI273", "Error calling RPC service", err)
			log.Printf("OCPI273: Request=%#v, Response=%#v", sessionUpdatedRequest, sessionUpdatedResponse)
		}
	}
}

func (r *SessionResolver) waitForEvseStatus(credential db.Credential, locationID, evseID int64, evseStatus db.EvseStatus, sessionID int64, sessionFromStatus db.SessionStatusType, sessionToStatus db.SessionStatusType, timeoutSeconds int) {
	ctx := context.Background()
	deadline := time.Now().Add(time.Duration(timeoutSeconds) * time.Second)
	log.Printf("Waiting for Evse status change to %v over %v seconds: LocationID=%v, EvseID=%v", evseStatus, timeoutSeconds, locationID, evseID)

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, coreLocation.IDENTIFIER, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI302", "Error getting version endpoint", err)
		log.Printf("OCPI302: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, coreLocation.IDENTIFIER)
		return
	}

	location, err := r.LocationResolver.Repository.GetLocation(ctx, locationID)

	if err != nil {
		metrics.RecordError("OCPI303", "Error getting location", err)
		log.Printf("OCPI303: LocationID=%v", locationID)
		return
	}

	evse, err := r.LocationResolver.Repository.GetEvse(ctx, evseID)

	if err != nil {
		metrics.RecordError("OCPI304", "Error getting evse", err)
		log.Printf("OCPI304: EvseID=%v", evseID)
		return
	}

	evseUrl := fmt.Sprintf("%s/%s/%s", versionEndpoint.Url, location.Uid, evse.Uid)
	requestUrl, err := url.Parse(evseUrl)

	if err != nil {
		metrics.RecordError("OCPI305", "Error parsing url", err)
		log.Printf("OCPI305: Url=%v", evseUrl)
		return
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)

	for {
		time.Sleep(10 * time.Second)

		if time.Now().After(deadline) {
			log.Printf("Stopped waiting for Evse status change: LocationID=%v, EvseID=%v", locationID, evseID)
			break
		}

		response, err := r.OcpiService.Do(http.MethodGet, requestUrl.String(), header, nil)

		if err != nil {
			metrics.RecordError("OCPI306", "Error making request", err)
			log.Printf("OCPI306: Method=%v, Url=%v, Header=%#v", http.MethodGet, requestUrl.String(), header)
			continue
		}

		evseDto, err := r.LocationResolver.EvseResolver.UnmarshalPullDto(response.Body)
		defer response.Body.Close()

		if err != nil {
			metrics.RecordError("OCPI307", "Error unmarshaling response", err)
			util.LogHttpResponse("OCPI307", requestUrl.String(), response, true)
			continue
		}

		if evseDto.StatusCode == transportation.STATUS_CODE_OK && evseDto.Data.Status != nil {
			responseEvseStatus := *evseDto.Data.Status

			log.Printf("Evse status is %v: LocationID=%v, EvseID=%v", responseEvseStatus, locationID, evseID)

			if responseEvseStatus == evseStatus {
				session, err := r.Repository.GetSession(ctx, sessionID)

				if err != nil {
					metrics.RecordError("OCPI308", "Error getting session", err)
					log.Printf("OCPI308: SessionID=%v", sessionID)
					continue
				}

				if session.Status == sessionFromStatus {
					log.Printf("Manually updating session status to %v: SessionUid=%v", sessionToStatus, session.Uid)

					sessionDto := dto.SessionDto{
						Status: &sessionToStatus,
					}

					r.ReplaceSessionByIdentifier(ctx, credential, util.NilString(session.CountryCode), util.NilString(session.PartyID), session.Uid, &sessionDto)
				}

				break
			}
		}
	}
}
