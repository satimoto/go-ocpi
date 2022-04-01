package token

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type TokenRepository interface {
	CreateToken(ctx context.Context, arg db.CreateTokenParams) (db.Token, error)
	CreateTokenAuthorization(ctx context.Context, arg db.CreateTokenAuthorizationParams) (db.TokenAuthorization, error)
	DeleteTokenByUid(ctx context.Context, uid string) error
	GetToken(ctx context.Context, id int64) (db.Token, error)
	GetTokenByUid(ctx context.Context, uid string) (db.Token, error)
	ListTokens(ctx context.Context, arg db.ListTokensParams) ([]db.Token, error)
	SetTokenAuthorizationConnector(ctx context.Context, arg db.SetTokenAuthorizationConnectorParams) error
	SetTokenAuthorizationEvse(ctx context.Context, arg db.SetTokenAuthorizationEvseParams) error
	UpdateTokenByUid(ctx context.Context, arg db.UpdateTokenByUidParams) (db.Token, error)
}

type TokenResolver struct {
	Repository TokenRepository
	*credential.CredentialResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TokenResolver {
	repo := TokenRepository(repositoryService)
	return &TokenResolver{
		Repository:         repo,
		CredentialResolver: credential.NewResolver(repositoryService),
	}
}

func (r *TokenResolver) CreateTokenAuthorization(ctx context.Context, token db.Token, payload *LocationReferencesPayload) *db.TokenAuthorization {
	tokenAuthorizationParams := NewCreateTokenAuthorizationParams(token.ID)

	if payload != nil {
		tokenAuthorizationParams.LocationID = util.SqlNullString(payload.LocationID)
	}

	if tokenAuthorization, err := r.Repository.CreateTokenAuthorization(ctx, tokenAuthorizationParams); err == nil {
		r.createTokenAuthorizationEvses(ctx, tokenAuthorization.ID, payload)
		r.createTokenAuthorizationConnectors(ctx, tokenAuthorization.ID, payload)

		return &tokenAuthorization
	}

	return nil
}

func (r *TokenResolver) ReplaceToken(ctx context.Context, uid string, payload *TokenPayload) *db.Token {
	if payload != nil {
		token, err := r.Repository.GetTokenByUid(ctx, uid)

		if err == nil {
			tokenParams := NewUpdateTokenByUidParams(token)

			if payload.AuthID != nil {
				tokenParams.AuthID = *payload.AuthID
			}

			if payload.Issuer != nil {
				tokenParams.Issuer = *payload.Issuer
			}

			if payload.Language != nil {
				tokenParams.Language = util.SqlNullString(payload.Language)
			}

			if payload.LastUpdated != nil {
				tokenParams.LastUpdated = *payload.LastUpdated
			}

			if payload.Type != nil {
				tokenParams.Type = *payload.Type
			}

			if payload.Valid != nil {
				tokenParams.Valid = *payload.Valid
			}

			if payload.VisualNumber != nil {
				tokenParams.VisualNumber = util.SqlNullString(payload.VisualNumber)
			}

			if payload.Whitelist != nil {
				tokenParams.Whitelist = *payload.Whitelist
			}

			token, err = r.Repository.UpdateTokenByUid(ctx, tokenParams)
		} else {
			tokenParams := NewCreateTokenParams(payload)

			token, err = r.Repository.CreateToken(ctx, tokenParams)
		}

		return &token
	}

	return nil
}

func (r *TokenResolver) createTokenAuthorizationConnectors(ctx context.Context, tokenAuthorizationID int64, payload *LocationReferencesPayload) {
	if payload != nil {
		for _, connectorId := range payload.ConnectorIds {
			r.Repository.SetTokenAuthorizationConnector(ctx, db.SetTokenAuthorizationConnectorParams{
				TokenAuthorizationID: tokenAuthorizationID,
				ConnectorUid:         *connectorId,
			})
		}
	}
}

func (r *TokenResolver) createTokenAuthorizationEvses(ctx context.Context, tokenAuthorizationID int64, payload *LocationReferencesPayload) {
	if payload != nil {
		for _, evseUid := range payload.EvseUids {
			r.Repository.SetTokenAuthorizationEvse(ctx, db.SetTokenAuthorizationEvseParams{
				TokenAuthorizationID: tokenAuthorizationID,
				EvseUid:              *evseUid,
			})
		}
	}
}
