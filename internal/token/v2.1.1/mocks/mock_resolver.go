package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1/mocks"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *token.TokenResolver {
	repo := token.TokenRepository(repositoryService)

	return &token.TokenResolver{
		Repository:         repo,
		CredentialResolver: credential.NewResolver(repositoryService, requester),
	}
}
