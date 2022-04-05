package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *token.TokenResolver {
	repo := token.TokenRepository(repositoryService)

	return &token.TokenResolver{
		Repository:         repo,
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
	}
}
