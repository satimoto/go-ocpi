package token

import (
	"github.com/satimoto/go-datastore/db"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
)

type RpcTokenRepository interface{}

type RpcTokenResolver struct {
	Repository RpcTokenRepository
	*token.TokenResolver
}

func NewResolver(repositoryService *db.RepositoryService) *RpcTokenResolver {
	repo := RpcTokenRepository(repositoryService)
	return &RpcTokenResolver{
		Repository:    repo,
		TokenResolver: token.NewResolver(repositoryService),
	}
}
