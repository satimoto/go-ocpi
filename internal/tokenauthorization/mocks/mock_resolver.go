package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/tokenauthorization"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *tokenauthorization.TokenAuthorizationResolver {
	repo := tokenauthorization.TokenAuthorizationRepository(repositoryService)

	return &tokenauthorization.TokenAuthorizationResolver{
		Repository: repo,
	}
}
