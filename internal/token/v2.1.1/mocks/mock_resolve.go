package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	tokenMocks "github.com/satimoto/go-datastore/pkg/token/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	tokenauthorization "github.com/satimoto/go-ocpi/internal/tokenauthorization/v2.1.1/mocks"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *token.TokenResolver {
	return &token.TokenResolver{
		Repository:                 tokenMocks.NewRepository(repositoryService),
		OcpiService:                services.OcpiService,
		TokenAuthorizationResolver: tokenauthorization.NewResolver(repositoryService, services),
		VersionDetailResolver:      versiondetail.NewResolver(repositoryService, services),
	}
}
