package tokenauthorization

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *TokenAuthorizationResolver) CreateTokenAuthorization(ctx context.Context, credential db.Credential, token db.Token, locationReferencesDto *dto.LocationReferencesDto) (*db.TokenAuthorization, error) {
	user, err := r.UserRepository.GetUser(ctx, token.UserID)

	if err != nil {
		metrics.RecordError("OCPI288", "Error retrieving user", err)
		log.Printf("OCPI288: UserID=%v", token.UserID)
		return nil, errors.New("Authorization error")
	}

	if !user.DeviceToken.Valid {
		return nil, errors.New("Please enable notifications in your Satimoto application")
	}

	listSessionInvoicesParams := db.ListSessionInvoicesByUserIDParams{
		ID:        user.ID,
		IsSettled: false,
		IsExpired: false,
	}

	// Check if there are unsettled invoices from a previous session
	if sessionInvoices, err := r.SessionRepository.ListSessionInvoicesByUserID(ctx, listSessionInvoicesParams); err == nil && len(sessionInvoices) > 0 {
		r.SendContentNotification(user, "Authorization Failed", "Please fund your Satimoto application and try again")

		return nil, errors.New("Please fund your Satimoto application and try again")
	}

	if token.Type == db.TokenTypeRFID {
		// Check if user is restricted, has a node and has been active
		if user.IsRestricted || !user.NodeID.Valid {
			r.SendContentNotification(user, "Card Authorization Failed", "Please fund your Satimoto application and try again")

			return nil, errors.New("Please fund your Satimoto application and try again")
		}
	}

	if lastTokenAuthorization, err := r.Repository.GetLastTokenAuthorizationByTokenID(ctx, token.ID); err == nil {
		// Last token authorization for this token has no session, unauthorise the token authorization
		updateTokenAuthorizationParams := param.NewUpdateTokenAuthorizationParams(lastTokenAuthorization)
		updateTokenAuthorizationParams.Authorized = false

		_, err = r.Repository.UpdateTokenAuthorizationByAuthorizationID(ctx, updateTokenAuthorizationParams)

		if err != nil {
			metrics.RecordError("OCPI332", "Error updating token authorization", err)
			log.Printf("OCPI332: Params=%#v", updateTokenAuthorizationParams)
		}
	}

	createTokenAuthorizationParams := param.NewCreateTokenAuthorizationParams(token.ID)
	createTokenAuthorizationParams.Authorized = token.Type == db.TokenTypeOTHER

	if locationReferencesDto != nil {
		createTokenAuthorizationParams.LocationID = util.SqlNullString(locationReferencesDto.LocationID)
	}

	tokenAuthorization, err := r.Repository.CreateTokenAuthorization(ctx, createTokenAuthorizationParams)

	if err != nil {
		metrics.RecordError("OCPI206", "Error creating token authorization", err)
		log.Printf("OCPI206: Params=%#v", createTokenAuthorizationParams)
		return nil, errors.New("Authorization error")
	}

	r.createTokenAuthorizationRelations(ctx, tokenAuthorization.ID, locationReferencesDto)

	if !createTokenAuthorizationParams.Authorized {
		// Token authentication is not authorized because its initiated
		// by an RFID card. The request needs to be forwarded to the user's
		// device, which then responds if it is authorized or not.
		// If there is a timeout in waiting for the response, the token
		// authorize request is rejected.
		asyncChan := r.AsyncService.Add(createTokenAuthorizationParams.AuthorizationID)
		r.SendDataNotification(user, createTokenAuthorizationParams.AuthorizationID)
		timeout := util.GetEnvInt32("TOKEN_AUTHORIZATION_TIMEOUT", 5)

		select {
		case asyncResult := <-asyncChan:
			log.Printf("Token authorization received: %v", createTokenAuthorizationParams.AuthorizationID)
			r.AsyncService.Remove(createTokenAuthorizationParams.AuthorizationID)

			if !asyncResult.Bool {
				return nil, errors.New("Please fund your Satimoto application and try again")
			}
		case <-time.After(time.Duration(timeout) * time.Second):
			log.Printf("Token authorization timeout: %v", createTokenAuthorizationParams.AuthorizationID)
			r.AsyncService.Remove(createTokenAuthorizationParams.AuthorizationID)

			r.SendContentNotification(user, "Card Authorization Failed", "Please open your Satimoto application and try again")

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

	return &tokenAuthorization, nil
}

func (r *TokenAuthorizationResolver) createTokenAuthorizationRelations(ctx context.Context, tokenAuthorizationID int64, locationReferencesDto *dto.LocationReferencesDto) {
	if locationReferencesDto != nil {
		for _, evseUid := range locationReferencesDto.EvseUids {
			if evse, err := r.EvseResolver.Repository.GetEvseByUid(ctx, *evseUid); err == nil {
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
}

func (r *TokenAuthorizationResolver) waitForEvsesStatus(credential db.Credential, token db.Token, tokenAuthorization db.TokenAuthorization, locationReferencesDto *dto.LocationReferencesDto, evseStatus db.EvseStatus, timeoutSeconds int) {
	if locationReferencesDto != nil && locationReferencesDto.LocationID != nil && len(locationReferencesDto.EvseUids) > 0 {
		cancelCtx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for _, evseUid := range locationReferencesDto.EvseUids {
			go r.EvseResolver.WaitForEvseStatus(credential, token, tokenAuthorization, *locationReferencesDto.LocationID, *evseUid, evseStatus, cancelCtx, cancel, timeoutSeconds)
		}

		<-cancelCtx.Done()
	}
}
