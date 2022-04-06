package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	command "github.com/satimoto/go-ocpi-api/internal/command/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *ocpi.OCPIRequester) *command.CommandResolver {
	repo := command.CommandRepository(repositoryService)

	return &command.CommandResolver{
		Repository:    repo,
		TokenResolver: token.NewResolver(repositoryService, requester),
	}
}
