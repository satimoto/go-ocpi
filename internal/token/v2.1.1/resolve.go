package token

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1"
)

type TokenRepository interface {
	CreateToken(ctx context.Context, arg db.CreateTokenParams) (db.Token, error)
	DeleteTokenByUid(ctx context.Context, uid string) error
	GetToken(ctx context.Context, id int64) (db.Token, error)
	GetTokenByAuthID(ctx context.Context, authID string) (db.Token, error)
	GetTokenByUid(ctx context.Context, uid string) (db.Token, error)
	GetTokenByUserID(ctx context.Context, arg db.GetTokenByUserIDParams) (db.Token, error)
	ListTokens(ctx context.Context, arg db.ListTokensParams) ([]db.Token, error)
	ListTokensByUserID(ctx context.Context, userID int64) ([]db.Token, error)
	UpdateTokenByUid(ctx context.Context, arg db.UpdateTokenByUidParams) (db.Token, error)
}

type TokenResolver struct {
	Repository TokenRepository
	*transportation.OCPIRequester
	*credential.CredentialResolver
	*tokenauthorization.TokenAuthorizationResolver
	*versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TokenResolver {
	repo := TokenRepository(repositoryService)
	return &TokenResolver{
		Repository:                 repo,
		OCPIRequester:              transportation.NewOCPIRequester(),
		CredentialResolver:         credential.NewResolver(repositoryService),
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
		VersionDetailResolver:      versiondetail.NewResolver(repositoryService),
	}
}
