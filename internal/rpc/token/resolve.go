package token

import (
	"github.com/satimoto/go-datastore/db"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
)

type RpcCredentialRepository interface{}

type RpcCredentialResolver struct {
	Repository RpcCredentialRepository
	*token.TokenResolver
}

func NewResolver(repositoryService *db.RepositoryService) *RpcCredentialResolver {
	repo := RpcCredentialRepository(repositoryService)
	return &RpcCredentialResolver{
		Repository:    repo,
		TokenResolver: token.NewResolver(repositoryService),
	}
}
