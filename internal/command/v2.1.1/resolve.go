package command

import (
	"github.com/satimoto/go-datastore/pkg/command"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/async"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/service"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type CommandResolver struct {
	Repository            command.CommandRepository
	OcpiService           *transportation.OcpiService
	AsyncService          *async.AsyncService
	EvseResolver          *evse.EvseResolver
	TokenResolver         *token.TokenResolver
	VersionDetailResolver *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *CommandResolver {
	return &CommandResolver{
		Repository:            command.NewRepository(repositoryService),
		OcpiService:           services.OcpiService,
		AsyncService:          services.AsyncService,
		EvseResolver:          evse.NewResolver(repositoryService, services),
		TokenResolver:         token.NewResolver(repositoryService, services),
		VersionDetailResolver: versiondetail.NewResolver(repositoryService, services),
	}
}
