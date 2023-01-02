package command

import (
	"github.com/satimoto/go-datastore/pkg/command"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/service"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type CommandResolver struct {
	Repository            command.CommandRepository
	OcpiService           *transportation.OcpiService
	TokenResolver         *token.TokenResolver
	VersionDetailResolver *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *CommandResolver {
	return &CommandResolver{
		Repository:            command.NewRepository(repositoryService),
		OcpiService:           services.OcpiService,
		TokenResolver:         token.NewResolver(repositoryService, services),
		VersionDetailResolver: versiondetail.NewResolver(repositoryService, services),
	}
}
