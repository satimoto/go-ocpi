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
)

func (r *TokenAuthorizationResolver) CreateTokenAuthorization(ctx context.Context, token db.Token, locationReferencesDto *dto.LocationReferencesDto) (*db.TokenAuthorization, error) {
	tokenAuthorizationParams := param.NewCreateTokenAuthorizationParams(token.ID)
	tokenAuthorizationParams.Authorized = token.Type == db.TokenTypeOTHER
	tokenAuthorizationParams.SigningKey = r.createTokenAuthorizationSigningKey()

	if locationReferencesDto != nil {
		tokenAuthorizationParams.LocationID = util.SqlNullString(locationReferencesDto.LocationID)
	}

	tokenAuthorization, err := r.Repository.CreateTokenAuthorization(ctx, tokenAuthorizationParams)

	if err != nil {
		util.LogOnError("OCPI206", "Error creating token authorization", err)
		log.Printf("OCPI206: Params=%#v", tokenAuthorizationParams)
		return nil, errors.New("error creating token authorization")
	}

	r.createTokenAuthorizationRelations(ctx, tokenAuthorization.ID, locationReferencesDto)

	if !tokenAuthorizationParams.Authorized {
		// Token authentication is not authorized because its initiated
		// by an RFID card. The request needs to be forwarded to the user's
		// device, which then responds if it is authorized or not.
		// If there is a timeout in waiting for the response, the token
		// authorize request is rejected.
		asyncChan := r.AsyncService.Add(tokenAuthorizationParams.AuthorizationID)

		user, err := r.UserRepository.GetUser(ctx, token.UserID)

		if err != nil {
			util.LogOnError("OCPI285", "Error retrieving node", err)
			log.Printf("OCPI285: UserID=%v", token.UserID)
			return nil, nil
		}

		r.SendNotification(user, tokenAuthorizationParams.AuthorizationID)

		select {
		case asyncResult := <-asyncChan:
			if !asyncResult.Bool {
				return nil, nil
			}
		case <-time.After(55 * time.Second):
			return nil, nil
		}

		updateTokenAuthorizationParams := param.NewUpdateTokenAuthorizationParams(tokenAuthorization)
		updateTokenAuthorizationParams.Authorized = true

		r.Repository.UpdateTokenAuthorizationByAuthorizationID(ctx, updateTokenAuthorizationParams)

		if err != nil {
			util.LogOnError("OCPI287", "Error updating token authorization", err)
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
					util.LogOnError("OCPI207", "Error setting token authorization evse", err)
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
							util.LogOnError("OCPI208", "Error setting token authorization connector", err)
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
