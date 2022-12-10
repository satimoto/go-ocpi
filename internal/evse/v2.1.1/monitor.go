package evse

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/util"
	coreLocation "github.com/satimoto/go-ocpi/internal/location"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/pkg/ocpi"
	ocpiSession "github.com/satimoto/go-ocpi/pkg/ocpi/session"
)

func (r *EvseResolver) WaitForEvseStatus(credential db.Credential, token db.Token, tokenAuthorization db.TokenAuthorization, locationUid, evseUid string, evseStatus db.EvseStatus, cancelCtx context.Context, cancel context.CancelFunc, timeoutSeconds int) {
	defer cancel()

	ctx := context.Background()
	deadline := time.Now().Add(time.Duration(timeoutSeconds) * time.Second)
	log.Printf("Waiting for Evse status change to %v over %v seconds: LocationUid=%v, EvseUid=%v", evseStatus, timeoutSeconds, locationUid, evseUid)

	versionEndpoint, err := r.VersionDetailResolver.GetVersionEndpointByIdentity(ctx, coreLocation.IDENTIFIER, credential.CountryCode, credential.PartyID)

	if err != nil {
		metrics.RecordError("OCPI302", "Error getting version endpoint", err)
		log.Printf("OCPI302: CountryCode=%v, PartyID=%v, Identifier=%v", credential.CountryCode, credential.PartyID, coreLocation.IDENTIFIER)
		return
	}

	evseUrl := fmt.Sprintf("%s/%s/%s", versionEndpoint.Url, locationUid, evseUid)
	requestUrl, err := url.Parse(evseUrl)

	if err != nil {
		metrics.RecordError("OCPI305", "Error parsing url", err)
		log.Printf("OCPI305: Url=%v", evseUrl)
		return
	}

	header := transportation.NewOcpiRequestHeader(&credential.ClientToken.String, nil, nil)

waitLoop:
	for {
		select {
		case <-cancelCtx.Done():
			log.Printf("Cancelled. Stop waiting for Evse status change: LocationUid=%v, EvseUid=%v", locationUid, evseUid)
			break waitLoop
		case <-time.After(10 * time.Second):
		}

		if time.Now().After(deadline) {
			log.Printf("Timeout. Stop waiting for Evse status change: LocationUid=%v, EvseUid=%v", locationUid, evseUid)
			break waitLoop
		}

		_, err := r.SessionRepository.GetSessionByAuthorizationID(ctx, tokenAuthorization.AuthorizationID)

		if err == nil {
			// Session found
			log.Printf("Session. Stop waiting for Evse status change: LocationUid=%v, EvseUid=%v", locationUid, evseUid)
			break waitLoop
		}

		updatedTokenAuthorization, err := r.TokenAuthorizationRepository.GetTokenAuthorizationByAuthorizationID(ctx, tokenAuthorization.AuthorizationID)

		if err != nil && !updatedTokenAuthorization.Authorized {
			// Token authorization has been unauthorized
			log.Printf("Unauthorized. Stop waiting for Evse status change: LocationUid=%v, EvseUid=%v", locationUid, evseUid)
			break waitLoop
		}

		response, err := r.OcpiService.Do(http.MethodGet, requestUrl.String(), header, nil)

		if err != nil {
			metrics.RecordError("OCPI306", "Error making request", err)
			log.Printf("OCPI306: Method=%v, Url=%v, Header=%#v", http.MethodGet, requestUrl.String(), header)
			continue
		}

		evseDto, err := r.UnmarshalPullDto(response.Body)
		defer response.Body.Close()

		if err != nil {
			metrics.RecordError("OCPI307", "Error unmarshaling response", err)
			util.LogHttpResponse("OCPI307", requestUrl.String(), response, true)
			continue
		}

		if evseDto.StatusCode == transportation.STATUS_CODE_OK && evseDto.Data.Status != nil {
			responseEvseStatus := *evseDto.Data.Status

			log.Printf("Evse status is %v: LocationUid=%v, EvseUid=%v", responseEvseStatus, locationUid, evseUid)

			if responseEvseStatus == evseStatus {
				log.Printf("Manually creating session %v", tokenAuthorization.AuthorizationID)
				r.createSession(ctx, credential, token, tokenAuthorization, db.SessionStatusTypeACTIVE, locationUid, evseUid)

				break waitLoop
			}
		}
	}
}

func (r *EvseResolver) createSession(ctx context.Context, credential db.Credential, token db.Token, tokenAuthorization db.TokenAuthorization, status db.SessionStatusType, locationUid, evseUid string) {
	timeNow := time.Now().UTC()
	location, err := r.LocationRepository.GetLocationByUid(ctx, locationUid)

	if err != nil {
		metrics.RecordError("OCPI317", "Error getting location", err)
		log.Printf("OCPI317: LocationUid=%v", locationUid)
		return
	}

	evse, err := r.Repository.GetEvseByUid(ctx, evseUid)

	if err != nil {
		metrics.RecordError("OCPI318", "Error getting evse", err)
		log.Printf("OCPI318: EvseUid=%v", evseUid)
		return
	}

	connectors, err := r.ConnectorResolver.Repository.ListConnectors(ctx, evse.ID)

	if err != nil || len(connectors) == 0 {
		metrics.RecordError("OCPI319", "Error getting connector", err)
		log.Printf("OCPI319: EvseUid=%v", evseUid)
		return
	}

	connector := connectors[0]

	if !connector.TariffID.Valid {
		metrics.RecordError("OCPI320", "Error no valid tariff", err)
		log.Printf("OCPI320: ConnectorUid=%v", connector.Uid)
		return
	}

	tariff, err := r.TariffRespository.GetTariffByUid(ctx, connector.TariffID.String)

	if err != nil {
		metrics.RecordError("OCPI321", "Error getting tariff", err)
		log.Printf("OCPI321: TariffUid=%v", connector.TariffID.String)
		return
	}

	authMethod := db.AuthMethodTypeAUTHREQUEST

	if token.Type == db.TokenTypeRFID {
		authMethod = db.AuthMethodTypeWHITELIST
	}

	createSessionParams := db.CreateSessionParams{
		Uid:             tokenAuthorization.AuthorizationID,
		CredentialID:    credential.ID,
		CountryCode:     location.CountryCode,
		PartyID:         location.PartyID,
		AuthorizationID: util.SqlNullString(tokenAuthorization.AuthorizationID),
		StartDatetime:   timeNow,
		Kwh:             0,
		AuthID:          token.AuthID,
		AuthMethod:      authMethod,
		UserID:          token.UserID,
		TokenID:         token.ID,
		LocationID:      location.ID,
		EvseID:          evse.ID,
		ConnectorID:     connector.ID,
		Currency:        tariff.Currency,
		TotalCost:       util.SqlNullFloat64(0),
		Status:          status,
		LastUpdated:     timeNow,
	}

	session, err := r.SessionRepository.CreateSession(ctx, createSessionParams)

	if err != nil {
		metrics.RecordError("OCPI322", "Error creating session", err)
		log.Printf("OCPI322: Params=%#v", createSessionParams)
		return
	}

	go r.sendOcpiRequest(session)
}

func (r *EvseResolver) sendOcpiRequest(session db.Session) {
	ctx := context.Background()
	node, err := r.NodeRepository.GetNodeByUserID(ctx, session.UserID)

	if err != nil {
		metrics.RecordError("OCPI323", "Error retrieving node", err)
		log.Printf("OCPI323: UserID=%v", session.UserID)
		return
	}

	// TODO: Handle failed RPC call more robustly
	ocpiService := ocpi.NewService(node.LspAddr)

	sessionCreatedRequest := ocpiSession.NewSessionCreatedRequest(session)
	sessionCreatedResponse, err := ocpiService.SessionCreated(ctx, sessionCreatedRequest)

	if err != nil {
		metrics.RecordError("OCPI324", "Error calling RPC service", err)
		log.Printf("OCPI324: Request=%#v, Response=%#v", sessionCreatedRequest, sessionCreatedResponse)
	}
}
