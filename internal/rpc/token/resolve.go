package token

import (
	"github.com/satimoto/go-datastore/pkg/db"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type RpcTokenRepository interface{}

type RpcTokenResolver struct {
	Repository RpcTokenRepository
	*token.TokenResolver
}

func NewResolver(repositoryService *db.RepositoryService, ocpiService *transportation.OcpiRequester) *RpcTokenResolver {
	repo := RpcTokenRepository(repositoryService)

	return &RpcTokenResolver{
		Repository:    repo,
		TokenResolver: token.NewResolver(repositoryService, ocpiService),
	}
}
