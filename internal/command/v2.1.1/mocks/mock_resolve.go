package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	command "github.com/satimoto/go-ocpi-api/internal/command/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *command.CommandResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *command.CommandResolver {
	repo := command.CommandRepository(repositoryService)

	return &command.CommandResolver{
		Repository:    repo,
		OcpiRequester: ocpiRequester,
		TokenResolver: token.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
