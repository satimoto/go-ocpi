package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	tokenauthorization "github.com/satimoto/go-ocpi-api/internal/tokenauthorization/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OCPIRequester) *token.TokenResolver {
	repo := token.TokenRepository(repositoryService)

	return &token.TokenResolver{
		Repository:                 repo,
		OCPIRequester:              ocpiRequester,
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
	}
}
