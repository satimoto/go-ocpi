package tokenauthorization

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	secp "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	coreLocation "github.com/satimoto/go-ocpi/internal/location"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/pkg/ocpi"
	ocpiSession "github.com/satimoto/go-ocpi/pkg/ocpi/session"
)

func (r *TokenAuthorizationResolver) CreateTokenAuthorization(ctx context.Context, credential db.Credential, token db.Token, locationReferencesDto *dto.LocationReferencesDto) (*db.TokenAuthorization, error) {
	if token.Type == db.TokenTypeRFID {
		// Check if user is restricted, has a node and has been active
		user, err := r.UserRepository.GetUserByTokenID(ctx, token.ID)

		if err != nil {
			metrics.RecordError("OCPI288", "Error retrieving user", err)
			log.Printf("OCPI288: UserID=%v", token.UserID)
			return nil, nil
		}

		if user.IsRestricted || !user.NodeID.Valid || !user.LastActiveDate.Valid {
			return nil, errors.New("Please fund your Satimoto application and try again")
		}

		// Check if the use has been active in the last 5 days
		fiveDaysAgo := time.Now().Add(time.Hour * 24 * -5)

		if fiveDaysAgo.After(user.LastActiveDate.Time) {
			// TODO: Send a notification
			return nil, errors.New("Please open your Satimoto application and try again")
		}
	}

	tokenAuthorizationParams := param.NewCreateTokenAuthorizationParams(token.ID)
	tokenAuthorizationParams.Authorized = token.Type == db.TokenTypeOTHER
	tokenAuthorizationParams.SigningKey = r.createTokenAuthorizationSigningKey()

	if locationReferencesDto != nil {
		tokenAuthorizationParams.LocationID = util.SqlNullString(locationReferencesDto.LocationID)
	}

	tokenAuthorization, err := r.Repository.CreateTokenAuthorization(ctx, tokenAuthorizationParams)

	if err != nil {
		metrics.RecordError("OCPI206", "Error creating token authorization", err)
		log.Printf("OCPI206: Params=%#v", tokenAuthorizationParams)
		return nil, errors.New("Authorization error")
	}

	evseUids, _ := r.createTokenAuthorizationRelations(ctx, tokenAuthorization.ID, locationReferencesDto)

	if !tokenAuthorizationParams.Authorized {
		// Token authentication is not authorized because its initiated
		// by an RFID card. The request needs to be forwarded to the user's
		// device, which then responds if it is authorized or not.
		// If there is a timeout in waiting for the response, the token
		// authorize request is rejected.
		user, err := r.UserRepository.GetUser(ctx, token.UserID)

		if err != nil {
			metrics.RecordError("OCPI285", "Error retrieving user", err)
			log.Printf("OCPI285: UserID=%v", token.UserID)
			return nil, nil
		}

		if !user.DeviceToken.Valid {
			return nil, errors.New("Please enable notifications in your Satimoto application")
		}

		asyncChan := r.AsyncService.Add(tokenAuthorizationParams.AuthorizationID)
		r.SendNotification(user, tokenAuthorizationParams.AuthorizationID)
		timeout := util.GetEnvInt32("TOKEN_AUTHORIZATION_TIMEOUT", 5)

		select {
		case asyncResult := <-asyncChan:
			log.Printf("Token authorization received: %v", tokenAuthorizationParams.AuthorizationID)
			r.AsyncService.Remove(tokenAuthorizationParams.AuthorizationID)

			if !asyncResult.Bool {
				return nil, errors.New("Please fund your Satimoto application and try again")
			}
		case <-time.After(time.Duration(timeout) * time.Second):
			log.Printf("Token authorization timeout: %v", tokenAuthorizationParams.AuthorizationID)
			r.AsyncService.Remove(tokenAuthorizationParams.AuthorizationID)

			return nil, errors.New("Authorization timeout")
		}

		updateTokenAuthorizationParams := param.NewUpdateTokenAuthorizationParams(tokenAuthorization)
		updateTokenAuthorizationParams.Authorized = true

		updatedTokenAuthorization, err := r.Repository.UpdateTokenAuthorizationByAuthorizationID(ctx, updateTokenAuthorizationParams)

		if err != nil {
			metrics.RecordError("OCPI287", "Error updating token authorization", err)
			log.Printf("OCPI287: Params=%#v", updateTokenAuthorizationParams)
		} else {
			tokenAuthorization = updatedTokenAuthorization
		}
	}

	if tokenAuthorization.Authorized && locationReferencesDto.LocationID != nil {
		go r.waitForEvsesStatus(credential, token, tokenAuthorization, *locationReferencesDto.LocationID, evseUids, db.EvseStatusCHARGING, 150)
	}

	return &tokenAuthorization, nil
}

func (r *TokenAuthorizationResolver) createTokenAuthorizationRelations(ctx context.Context, tokenAuthorizationID int64, locationReferencesDto *dto.LocationReferencesDto) (evseUids, connectorUids []string) {
	if locationReferencesDto != nil {
		for _, evseUid := range locationReferencesDto.EvseUids {
			if evse, err := r.EvseResolver.Repository.GetEvseByUid(ctx, *evseUid); err == nil {
				evseUids = append(evseUids, *evseUid)

				setTokenAuthorizationEvseParams := db.SetTokenAuthorizationEvseParams{
					TokenAuthorizationID: tokenAuthorizationID,
					EvseID:               evse.ID,
				}

				err = r.Repository.SetTokenAuthorizationEvse(ctx, setTokenAuthorizationEvseParams)

				if err != nil {
					metrics.RecordError("OCPI207", "Error setting token authorization evse", err)
					log.Printf("OCPI207: Params=%#v", setTokenAuthorizationEvseParams)
				}

				for _, connectorId := range locationReferencesDto.ConnectorIds {
					getConnectorByEvseParams := db.GetConnectorByEvseParams{
						EvseID: evse.ID,
						Uid:    *connectorId,
					}

					if connector, err := r.ConnectorResolver.Repository.GetConnectorByEvse(ctx, getConnectorByEvseParams); err == nil {
						connectorUids = append(connectorUids, *connectorId)

						setTokenAuthorizationConnectorParams := db.SetTokenAuthorizationConnectorParams{
							TokenAuthorizationID: tokenAuthorizationID,
							ConnectorID:          connector.ID,
						}

						err = r.Repository.SetTokenAuthorizationConnector(ctx, setTokenAuthorizationConnectorParams)

						if err != nil {
							metrics.RecordError("OCPI208", "Error setting token authorization connector", err)
							log.Printf("OCPI208: Params=%#v", setTokenAuthorizationConnectorParams)
						}
					}
				}
			}
		}
	}

	return evseUids, connectorUids
}

func (r *TokenAuthorizationResolver) createTokenAuthorizationSigningKey() []byte {
	var privateKey *secp.PrivateKey
	var err error

	for {
		if privateKey, err = secp.GeneratePrivateKey(); err == nil {
			break
		}
	}

	return privateKey.Serialize()
}

func (r *TokenAuthorizationResolver) waitForEvsesStatus(credential db.Credential, token db.Token, tokenAuthorization db.TokenAuthorization, locationUid string, evseUids []string, evseStatus db.EvseStatus, timeoutSeconds int) {
	cancelCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for _, evseUid := range evseUids {
		go r.waitForEvseStatus(credential, token, tokenAuthorization, locationUid, evseUid, evseStatus, cancelCtx, cancel, timeoutSeconds)
	}

	<-cancelCtx.Done()
}

func (r *TokenAuthorizationResolver) waitForEvseStatus(credential db.Credential, token db.Token, tokenAuthorization db.TokenAuthorization, locationUid, evseUid string, evseStatus db.EvseStatus, cancelCtx context.Context, cancel context.CancelFunc, timeoutSeconds int) {
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

		updatedTokenAuthorization, err := r.Repository.GetTokenAuthorizationByAuthorizationID(ctx, tokenAuthorization.AuthorizationID)

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

		evseDto, err := r.EvseResolver.UnmarshalPullDto(response.Body)
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

func (r *TokenAuthorizationResolver) createSession(ctx context.Context, credential db.Credential, token db.Token, tokenAuthorization db.TokenAuthorization, status db.SessionStatusType, locationUid, evseUid string) {
	timeNow := time.Now().UTC()
	location, err := r.LocationRepository.GetLocationByUid(ctx, locationUid)

	if err != nil {
		metrics.RecordError("OCPI317", "Error getting location", err)
		log.Printf("OCPI317: LocationUid=%v", locationUid)
		return
	}

	evse, err := r.EvseResolver.Repository.GetEvseByUid(ctx, evseUid)

	if err != nil {
		metrics.RecordError("OCPI318", "Error getting evse", err)
		log.Printf("OCPI318: EvseUid=%v", evseUid)
		return
	}

	connectors, err := r.EvseResolver.ConnectorResolver.Repository.ListConnectors(ctx, evse.ID)

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

	createSessionParams := db.CreateSessionParams{
		Uid:             tokenAuthorization.AuthorizationID,
		CredentialID:    credential.ID,
		CountryCode:     location.CountryCode,
		PartyID:         location.PartyID,
		AuthorizationID: util.SqlNullString(tokenAuthorization.AuthorizationID),
		StartDatetime:   timeNow,
		Kwh:             0,
		AuthID:          token.AuthID,
		AuthMethod:      db.AuthMethodTypeAUTHREQUEST,
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

func (r *TokenAuthorizationResolver) sendOcpiRequest(session db.Session) {
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
