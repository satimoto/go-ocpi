package mocks

import (
	commandMocks "github.com/satimoto/go-datastore/pkg/command/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	command "github.com/satimoto/go-ocpi/internal/command/v2.1.1"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1/mocks"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *command.CommandResolver {
	return &command.CommandResolver{
		Repository:            commandMocks.NewRepository(repositoryService),
		OcpiService:           services.OcpiService,
		AsyncService:          services.AsyncService,
		EvseResolver:          evse.NewResolver(repositoryService, services),
		TokenResolver:         token.NewResolver(repositoryService, services),
		VersionDetailResolver: versiondetail.NewResolver(repositoryService, services),
	}
}
