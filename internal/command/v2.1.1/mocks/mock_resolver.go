package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	command "github.com/satimoto/go-ocpi-api/internal/command/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *command.CommandResolver {
	repo := command.CommandRepository(repositoryService)

	return &command.CommandResolver{
		Repository:    repo,
		TokenResolver: token.NewResolver(repositoryService, requester),
	}
}
