package token

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type TokenRepository interface {
	CreateToken(ctx context.Context, arg db.CreateTokenParams) (db.Token, error)
	DeleteTokenByUid(ctx context.Context, uid string) error
	GetTokenByUid(ctx context.Context, uid string) (db.Token, error)
	ListTokens(ctx context.Context, arg db.ListTokensParams) ([]db.Token, error)
	UpdateTokenByUid(ctx context.Context, arg db.UpdateTokenByUidParams) (db.Token, error)
}

type TokenResolver struct {
	Repository TokenRepository
	*credential.CredentialResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TokenResolver {
	repo := TokenRepository(repositoryService)
	return &TokenResolver{
		Repository:          repo,
		CredentialResolver:  credential.NewResolver(repositoryService),
	}
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
