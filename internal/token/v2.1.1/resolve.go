package token

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/token"
	tokenauthorization "github.com/satimoto/go-ocpi/internal/tokenauthorization/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type TokenResolver struct {
	Repository                 token.TokenRepository
	OcpiRequester              *transportation.OcpiRequester
	TokenAuthorizationResolver *tokenauthorization.TokenAuthorizationResolver
	VersionDetailResolver      *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *TokenResolver {
	return &TokenResolver{
		Repository:                 token.NewRepository(repositoryService),
		OcpiRequester:              ocpiRequester,
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
		VersionDetailResolver:      versiondetail.NewResolver(repositoryService, ocpiRequester),
	}
}
