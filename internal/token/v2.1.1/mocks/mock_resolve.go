package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *ocpi.OCPIRequester) *token.TokenResolver {
	repo := token.TokenRepository(repositoryService)

	return &token.TokenResolver{
		Repository:                 repo,
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
	}
}
