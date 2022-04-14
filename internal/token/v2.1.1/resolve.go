package token

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1"
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
