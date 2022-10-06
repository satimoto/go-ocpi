package command

import (
	"github.com/satimoto/go-datastore/pkg/db"
	command "github.com/satimoto/go-ocpi/internal/command/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/credential"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	session "github.com/satimoto/go-ocpi/internal/session/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/sync"
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

func NewResolver(repositoryService *db.RepositoryService, syncService *sync.SyncService, ocpiRequester *transportation.OcpiRequester) *RpcCommandResolver {
	return &RpcCommandResolver{
		CommandResolver:    command.NewResolver(repositoryService, ocpiRequester),
		CredentialResolver: credential.NewResolver(repositoryService, syncService, ocpiRequester),
		LocationResolver:   location.NewResolver(repositoryService, ocpiRequester),
		SessionResolver:    session.NewResolver(repositoryService, ocpiRequester),
		TokenResolver:      token.NewResolver(repositoryService, ocpiRequester),
	}
}
