package mocks

import (
	commandMocks "github.com/satimoto/go-datastore/pkg/command/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	command "github.com/satimoto/go-ocpi/internal/command/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/service"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1/mocks"
)


func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *command.CommandResolver {
	return &command.CommandResolver{
		Repository:    commandMocks.NewRepository(repositoryService),
		OcpiService:   services.OcpiService,
		TokenResolver: token.NewResolver(repositoryService, services),
	}
}
