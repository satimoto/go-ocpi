package tokenauthorization

import (
	"context"
	"errors"
	"log"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *TokenAuthorizationResolver) CreateTokenAuthorization(ctx context.Context, token db.Token, dto *LocationReferencesDto) (*db.TokenAuthorization, error) {
	tokenAuthorizationParams := param.NewCreateTokenAuthorizationParams(token.ID)

	if dto != nil {
		tokenAuthorizationParams.LocationID = util.SqlNullString(dto.LocationID)
	}

	tokenAuthorization, err := r.Repository.CreateTokenAuthorization(ctx, tokenAuthorizationParams)

	if err != nil {
		util.LogOnError("OCPI206", "Error creating token authorization", err)
		log.Printf("OCPI206: Params=%#v", tokenAuthorizationParams)
		return nil, errors.New("error creating token authorization")
	}

	r.createTokenAuthorizationEvses(ctx, tokenAuthorization.ID, dto)
	r.createTokenAuthorizationConnectors(ctx, tokenAuthorization.ID, dto)

	return &tokenAuthorization, nil
}

func (r *TokenAuthorizationResolver) createTokenAuthorizationConnectors(ctx context.Context, tokenAuthorizationID int64, dto *LocationReferencesDto) {
	if dto != nil {
		for _, connectorId := range dto.ConnectorIds {
			getConnectorByUidParams := db.GetConnectorByUidParams{
				Uid: *connectorId,
			}
			connector, err := r.ConnectorResolver.Repository.GetConnectorByUid(ctx, getConnectorByUidParams)

			if err != nil {
				util.LogOnError("OCPI207", "Error creating token authorization", err)
				log.Printf("OCPI207: Params=%#v", getConnectorByUidParams)
				continue
			}

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

func (r *TokenAuthorizationResolver) createTokenAuthorizationEvses(ctx context.Context, tokenAuthorizationID int64, dto *LocationReferencesDto) {
	if dto != nil {
		for _, evseUid := range dto.EvseUids {
			evse, err := r.EvseResolver.Repository.GetEvseByUid(ctx, *evseUid)

			if err != nil {
				util.LogOnError("OCPI209", "Error retrieving evse", err)
				log.Printf("OCPI209: Uid=%v", *evseUid)
				continue
			}

			setTokenAuthorizationEvseParams := db.SetTokenAuthorizationEvseParams{
				TokenAuthorizationID: tokenAuthorizationID,
				EvseID:               evse.ID,
			}
			err = r.Repository.SetTokenAuthorizationEvse(ctx, setTokenAuthorizationEvseParams)

			if err != nil {
				util.LogOnError("OCPI210", "Error setting token authorization evse", err)
				log.Printf("OCPI210: Params=%#v", setTokenAuthorizationEvseParams)
			}
		}
	}
}
