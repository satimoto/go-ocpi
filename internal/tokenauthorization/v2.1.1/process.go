package tokenauthorization

import (
	"context"
	"errors"
	"log"
	"time"

	secp "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
	dto "github.com/satimoto/go-ocpi/internal/dto/v2.1.1"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
)

func (r *TokenAuthorizationResolver) CreateTokenAuthorization(ctx context.Context, token db.Token, locationReferencesDto *dto.LocationReferencesDto) (*db.TokenAuthorization, error) {
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

	r.createTokenAuthorizationRelations(ctx, tokenAuthorization.ID, locationReferencesDto)

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

		r.Repository.UpdateTokenAuthorizationByAuthorizationID(ctx, updateTokenAuthorizationParams)

		if err != nil {
			metrics.RecordError("OCPI287", "Error updating token authorization", err)
			log.Printf("OCPI287: Params=%#v", updateTokenAuthorizationParams)
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
