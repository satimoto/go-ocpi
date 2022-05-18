package token

import (
	"context"

	"github.com/satimoto/go-datastore/pkg/db"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

type TokenRepository interface {
	CreateToken(ctx context.Context, arg db.CreateTokenParams) (db.Token, error)
	DeleteTokenByUid(ctx context.Context, uid string) error
	GetToken(ctx context.Context, id int64) (db.Token, error)
	GetTokenByAuthID(ctx context.Context, authID string) (db.Token, error)
	GetTokenByUid(ctx context.Context, uid string) (db.Token, error)
	GetTokenByUserID(ctx context.Context, arg db.GetTokenByUserIDParams) (db.Token, error)
	ListCredentials(ctx context.Context) ([]db.Credential, error)
	ListTokens(ctx context.Context, arg db.ListTokensParams) ([]db.Token, error)
	ListTokensByUserID(ctx context.Context, userID int64) ([]db.Token, error)
	UpdateTokenByUid(ctx context.Context, arg db.UpdateTokenByUidParams) (db.Token, error)
}

type TokenResolver struct {
	Repository                 TokenRepository
	OcpiRequester              *transportation.OcpiRequester
	TokenAuthorizationResolver *tokenauthorization.TokenAuthorizationResolver
	VersionDetailResolver      *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TokenResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *TokenResolver {
	repo := TokenRepository(repositoryService)

	return &TokenResolver{
		Repository:                 repo,
		OcpiRequester:              ocpiRequester,
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
		VersionDetailResolver:      versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
