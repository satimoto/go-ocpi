package session

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	"github.com/satimoto/go-ocpi/internal/async"
	coreCommand "github.com/satimoto/go-ocpi/internal/command"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	coreLocation "github.com/satimoto/go-ocpi/internal/location"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/pkg/ocpi"
	ocpiSession "github.com/satimoto/go-ocpi/pkg/ocpi/session"
)

func (r *SessionResolver) sendOcpiRequest(session db.Session, sessionCreated, statusChanged bool) {
	ctx := context.Background()
	node, err := r.NodeRepository.GetNodeByUserID(ctx, session.UserID)

	if err != nil {
		metrics.RecordError("OCPI167", "Error retrieving node", err)
		log.Printf("OCPI167: UserID=%v", session.UserID)
		return
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

func (r *SessionResolver) updateCommand(session db.Session) {
	ctx := context.Background()

	if session.AuthorizationID.Valid {
		if session.Status == db.SessionStatusTypeACTIVE || session.Status == db.SessionStatusTypePENDING {
			updateCommandStartByAuthorizationIDParams := db.UpdateCommandStartByAuthorizationIDParams{
				AuthorizationID: util.SqlNullString(session.AuthorizationID.String),
				Status:          db.CommandResponseTypeACCEPTED,
				LastUpdated:     time.Now().UTC(),
			}

			command, err := r.CommandResolver.Repository.UpdateCommandStartByAuthorizationID(ctx, updateCommandStartByAuthorizationIDParams)

			if err != nil {
				metrics.RecordError("OCPI327", "Error updating command start", err)
				log.Printf("OCPI327: AuthorizationID=%#v", session.AuthorizationID)
				return
			}

			asyncKey := fmt.Sprintf(coreCommand.START_COMMAND_ASYNC_KEY, command.ID)
			asyncResult := async.AsyncResult{
				String: string(command.Status),
				Bool:   command.Status == db.CommandResponseTypeACCEPTED,
			}

			r.AsyncService.Set(asyncKey, asyncResult)
		} else if session.Status == db.SessionStatusTypeCOMPLETED {
			updateCommandStopBySessionIDParams := db.UpdateCommandStopBySessionIDParams{
				SessionID:   session.Uid,
				Status:      db.CommandResponseTypeACCEPTED,
				LastUpdated: time.Now().UTC(),
			}

			command, err := r.CommandResolver.Repository.UpdateCommandStopBySessionID(ctx, updateCommandStopBySessionIDParams)

			if err != nil {
				metrics.RecordError("OCPI328", "Error updating command stop", err)
				log.Printf("OCPI328: SessionUid=%#v", session.Uid)
				return
			}

			asyncKey := fmt.Sprintf(coreCommand.STOP_COMMAND_ASYNC_KEY, command.ID)
			asyncResult := async.AsyncResult{
				String: string(command.Status),
				Bool:   command.Status == db.CommandResponseTypeACCEPTED,
			}

			r.AsyncService.Set(asyncKey, asyncResult)
		}
	}
}

func (r *SessionResolver) waitForEvseStatus(credential db.Credential, locationID, evseID int64, evseStatus db.EvseStatus, sess db.Session, sessionFromStatus db.SessionStatusType, sessionToStatus db.SessionStatusType, timeoutSeconds int) {
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

waitLoop:
	for {
		time.Sleep(10 * time.Second)

		if time.Now().After(deadline) {
			log.Printf("Timeout. Stopped waiting for Evse status change: LocationID=%v, EvseID=%v", locationID, evseID)
			break waitLoop
		}

		if updatedSession, err := r.Repository.GetSession(ctx, sess.ID); err == nil {
			sess = updatedSession
		}

		if sess.Status != sessionFromStatus {
			log.Printf("Session. Stop waiting for Evse status change: LocationID=%v, EvseID=%v", locationID, evseID)
			break waitLoop
		}

		if sess.AuthorizationID.Valid {
			tokenAuthorization, err := r.TokenAuthorizationRepository.GetTokenAuthorizationByAuthorizationID(ctx, sess.AuthorizationID.String)

			if err != nil && !tokenAuthorization.Authorized {
				// Token authorization has been unauthorized
				log.Printf("Unauthorized. Stop waiting for Evse status change: LocationID=%v, EvseID=%v", locationID, evseID)
				break waitLoop
			}
		}

		// Check the local evse status
		if evse, err = r.LocationResolver.Repository.GetEvse(ctx, evseID); err == nil && evse.Status == evseStatus {
			log.Printf("Local Evse status is %v: LocationID=%v, EvseID=%v", evse.Status, locationID, evseID)
			log.Printf("Manually updating session status to %v: SessionUid=%v", sessionToStatus, sess.Uid)

			sessionDto := dto.SessionDto{
				Status: &sessionToStatus,
			}

			r.ReplaceSessionByIdentifier(ctx, credential, util.NilString(sess.CountryCode), util.NilString(sess.PartyID), sess.Uid, &sessionDto)

			break waitLoop
		}

		// Check the remote evse status
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

			log.Printf("Remote Evse status is %v: LocationID=%v, EvseID=%v", responseEvseStatus, locationID, evseID)

			if responseEvseStatus == evseStatus {
				log.Printf("Manually updating session status to %v: SessionUid=%v", sessionToStatus, sess.Uid)

				sessionDto := dto.SessionDto{
					Status: &sessionToStatus,
				}

				r.ReplaceSessionByIdentifier(ctx, credential, util.NilString(sess.CountryCode), util.NilString(sess.PartyID), sess.Uid, &sessionDto)

				break waitLoop
			}
		}
	}
}
