package command

import (
	"github.com/satimoto/go-datastore/pkg/db"
	command "github.com/satimoto/go-ocpi/internal/command/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/credential"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/service"
	session "github.com/satimoto/go-ocpi/internal/session/v2.1.1"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
)

type RpcCommandResolver struct {
	CommandResolver    *command.CommandResolver
	CredentialResolver *credential.CredentialResolver
	LocationResolver   *location.LocationResolver
	SessionResolver    *session.SessionResolver
	TokenResolver      *token.TokenResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *RpcCommandResolver {
	return &RpcCommandResolver{
		CommandResolver:    command.NewResolver(repositoryService, services),
		CredentialResolver: credential.NewResolver(repositoryService, services),
		LocationResolver:   location.NewResolver(repositoryService, services),
		SessionResolver:    session.NewResolver(repositoryService, services),
		TokenResolver:      token.NewResolver(repositoryService, services),
	}
}
