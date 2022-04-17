package command

import (
	"github.com/satimoto/go-datastore/db"
	command "github.com/satimoto/go-ocpi-api/internal/command/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/credential"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
)

type RpcCommandRepository interface{}

type RpcCommandResolver struct {
	Repository RpcCommandRepository
	*command.CommandResolver
	*credential.CredentialResolver
	*location.LocationResolver
	*token.TokenResolver
}

func NewResolver(repositoryService *db.RepositoryService) *RpcCommandResolver {
	repo := RpcCommandRepository(repositoryService)
	return &RpcCommandResolver{
		Repository:         repo,
		CommandResolver:    command.NewResolver(repositoryService),
		CredentialResolver: credential.NewResolver(repositoryService),
		LocationResolver:   location.NewResolver(repositoryService),
		TokenResolver:      token.NewResolver(repositoryService),
	}
}
