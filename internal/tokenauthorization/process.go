package tokenauthorization

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func (r *TokenAuthorizationResolver) CreateTokenAuthorization(ctx context.Context, token db.Token, dto *LocationReferencesDto) *db.TokenAuthorization {
	tokenAuthorizationParams := NewCreateTokenAuthorizationParams(token.ID)

	if dto != nil {
		tokenAuthorizationParams.LocationID = util.SqlNullString(dto.LocationID)
	}

	if tokenAuthorization, err := r.Repository.CreateTokenAuthorization(ctx, tokenAuthorizationParams); err == nil {
		r.createTokenAuthorizationEvses(ctx, tokenAuthorization.ID, dto)
		r.createTokenAuthorizationConnectors(ctx, tokenAuthorization.ID, dto)

		return &tokenAuthorization
	}

	return nil
}

func (r *TokenAuthorizationResolver) createTokenAuthorizationConnectors(ctx context.Context, tokenAuthorizationID int64, dto *LocationReferencesDto) {
	if dto != nil {
		for _, connectorId := range dto.ConnectorIds {
			r.Repository.SetTokenAuthorizationConnector(ctx, db.SetTokenAuthorizationConnectorParams{
				TokenAuthorizationID: tokenAuthorizationID,
				ConnectorUid:         *connectorId,
			})
		}
	}
}

func (r *TokenAuthorizationResolver) createTokenAuthorizationEvses(ctx context.Context, tokenAuthorizationID int64, dto *LocationReferencesDto) {
	if dto != nil {
		for _, evseUid := range dto.EvseUids {
			r.Repository.SetTokenAuthorizationEvse(ctx, db.SetTokenAuthorizationEvseParams{
				TokenAuthorizationID: tokenAuthorizationID,
				EvseUid:              *evseUid,
			})
		}
	}
}
