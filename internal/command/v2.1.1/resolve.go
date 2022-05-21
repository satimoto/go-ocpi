package command

import (
	"github.com/satimoto/go-datastore/pkg/command"
	"github.com/satimoto/go-datastore/pkg/db"
	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

type CommandResolver struct {
	Repository            command.CommandRepository
	OcpiRequester         *transportation.OcpiRequester
	SessionResolver       *session.SessionResolver
	TokenResolver         *token.TokenResolver
	VersionDetailResolver *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CommandResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *CommandResolver {
	return &CommandResolver{
		Repository:            command.NewRepository(repositoryService),
		OcpiRequester:         ocpiRequester,
		SessionResolver:       session.NewResolver(repositoryService),
		TokenResolver:         token.NewResolver(repositoryService),
		VersionDetailResolver: versiondetail.NewResolver(repositoryService),
	}
}
