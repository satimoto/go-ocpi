package tokenauthorization

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/param"
	"github.com/satimoto/go-datastore/pkg/util"
)

func (r *TokenAuthorizationResolver) CreateTokenAuthorization(ctx context.Context, token db.Token, dto *LocationReferencesDto) *db.TokenAuthorization {
	tokenAuthorizationParams := param.NewCreateTokenAuthorizationParams(token.ID)

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
			if connector, err := r.ConnectorResolver.Repository.GetConnectorByUid(ctx, db.GetConnectorByUidParams{Uid: *connectorId}); err == nil {
				r.Repository.SetTokenAuthorizationConnector(ctx, db.SetTokenAuthorizationConnectorParams{
					TokenAuthorizationID: tokenAuthorizationID,
					ConnectorID:          connector.ID,
				})
			}
		}
	}
}

func (r *TokenAuthorizationResolver) createTokenAuthorizationEvses(ctx context.Context, tokenAuthorizationID int64, dto *LocationReferencesDto) {
	if dto != nil {
		for _, evseUid := range dto.EvseUids {
			if evse, err := r.EvseResolver.Repository.GetEvseByUid(ctx, *evseUid); err == nil {
				r.Repository.SetTokenAuthorizationEvse(ctx, db.SetTokenAuthorizationEvseParams{
					TokenAuthorizationID: tokenAuthorizationID,
					EvseID:               evse.ID,
				})
			}
		}
	}
}
