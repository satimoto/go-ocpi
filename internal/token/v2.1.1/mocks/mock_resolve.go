package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	tokenMocks "github.com/satimoto/go-datastore/pkg/token/mocks"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	tokenauthorization "github.com/satimoto/go-ocpi/internal/tokenauthorization/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *token.TokenResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *token.TokenResolver {
	return &token.TokenResolver{
		Repository:                 tokenMocks.NewRepository(repositoryService),
		OcpiRequester:              ocpiRequester,
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService),
	}
}
