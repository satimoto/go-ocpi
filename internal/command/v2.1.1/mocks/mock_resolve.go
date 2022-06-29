package mocks

import (
	commandMocks "github.com/satimoto/go-datastore/pkg/command/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	command "github.com/satimoto/go-ocpi/internal/command/v2.1.1"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/transportation"
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
