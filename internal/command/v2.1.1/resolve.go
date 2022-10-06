package command

import (
	"github.com/satimoto/go-datastore/pkg/command"
	"github.com/satimoto/go-datastore/pkg/db"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type CommandResolver struct {
	Repository            command.CommandRepository
	OcpiRequester         *transportation.OcpiRequester
	TokenResolver         *token.TokenResolver
	VersionDetailResolver *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *CommandResolver {
	return &CommandResolver{
		Repository:            command.NewRepository(repositoryService),
		OcpiRequester:         ocpiRequester,
		TokenResolver:         token.NewResolver(repositoryService, ocpiRequester),
		VersionDetailResolver: versiondetail.NewResolver(repositoryService, ocpiRequester),
	}
}
