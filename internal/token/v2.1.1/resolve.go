package token

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/token"
	"github.com/satimoto/go-ocpi/internal/service"
	tokenauthorization "github.com/satimoto/go-ocpi/internal/tokenauthorization/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type TokenResolver struct {
	Repository                 token.TokenRepository
	OcpiService                *transportation.OcpiService
	TokenAuthorizationResolver *tokenauthorization.TokenAuthorizationResolver
	VersionDetailResolver      *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *TokenResolver {
	return &TokenResolver{
		Repository:                 token.NewRepository(repositoryService),
		OcpiService:                services.OcpiService,
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService, services),
		VersionDetailResolver:      versiondetail.NewResolver(repositoryService, services),
	}
}
