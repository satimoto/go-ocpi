package command

import (
	"github.com/satimoto/go-datastore/pkg/db"
	command "github.com/satimoto/go-ocpi/internal/command/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/credential"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	session "github.com/satimoto/go-ocpi/internal/session/v2.1.1"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type RpcCommandResolver struct {
	CommandResolver    *command.CommandResolver
	CredentialResolver *credential.CredentialResolver
	LocationResolver   *location.LocationResolver
	SessionResolver    *session.SessionResolver
	TokenResolver      *token.TokenResolver
}

func NewResolver(repositoryService *db.RepositoryService) *RpcCommandResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *RpcCommandResolver {
	return &RpcCommandResolver{
		CommandResolver:    command.NewResolverWithServices(repositoryService, ocpiRequester),
		CredentialResolver: credential.NewResolverWithServices(repositoryService, ocpiRequester),
		LocationResolver:   location.NewResolverWithServices(repositoryService, ocpiRequester),
		SessionResolver:    session.NewResolverWithServices(repositoryService, ocpiRequester),
		TokenResolver:      token.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
