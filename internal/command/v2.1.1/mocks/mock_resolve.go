package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	commandMocks "github.com/satimoto/go-datastore/pkg/command/mocks"
	command "github.com/satimoto/go-ocpi-api/internal/command/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *command.CommandResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *command.CommandResolver {
	return &command.CommandResolver{
		Repository:    commandMocks.NewRepository(repositoryService),
		OcpiRequester: ocpiRequester,
		TokenResolver: token.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
