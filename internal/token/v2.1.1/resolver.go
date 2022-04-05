package token

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/tokenauthorization"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

type TokenRepository interface {
	CreateToken(ctx context.Context, arg db.CreateTokenParams) (db.Token, error)
	DeleteTokenByUid(ctx context.Context, uid string) error
	GetToken(ctx context.Context, id int64) (db.Token, error)
	GetTokenByUid(ctx context.Context, uid string) (db.Token, error)
	ListTokens(ctx context.Context, arg db.ListTokensParams) ([]db.Token, error)
	UpdateTokenByUid(ctx context.Context, arg db.UpdateTokenByUidParams) (db.Token, error)
}

type TokenResolver struct {
	Repository TokenRepository
	*tokenauthorization.TokenAuthorizationResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TokenResolver {
	repo := TokenRepository(repositoryService)
	return &TokenResolver{
		Repository:                 repo,
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
	}
}

func (r *TokenResolver) ReplaceToken(ctx context.Context, uid string, dto *TokenDto) *db.Token {
	if dto != nil {
		token, err := r.Repository.GetTokenByUid(ctx, uid)

		if err == nil {
			tokenParams := NewUpdateTokenByUidParams(token)

			if dto.AuthID != nil {
				tokenParams.AuthID = *dto.AuthID
			}

			if dto.Issuer != nil {
				tokenParams.Issuer = *dto.Issuer
			}

			if dto.Language != nil {
				tokenParams.Language = util.SqlNullString(dto.Language)
			}

			if dto.LastUpdated != nil {
				tokenParams.LastUpdated = *dto.LastUpdated
			}

			if dto.Type != nil {
				tokenParams.Type = *dto.Type
			}

			if dto.Valid != nil {
				tokenParams.Valid = *dto.Valid
			}

			if dto.VisualNumber != nil {
				tokenParams.VisualNumber = util.SqlNullString(dto.VisualNumber)
			}

			if dto.Whitelist != nil {
				tokenParams.Whitelist = *dto.Whitelist
			}

			token, err = r.Repository.UpdateTokenByUid(ctx, tokenParams)
		} else {
			tokenParams := NewCreateTokenParams(dto)

			token, err = r.Repository.CreateToken(ctx, tokenParams)
		}

		return &token
	}

	return nil
}
