package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	versionMocks "github.com/satimoto/go-datastore/pkg/version/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/version"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *version.VersionResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *version.VersionResolver {
	return &version.VersionResolver{
		Repository:    versionMocks.NewRepository(repositoryService),
		OcpiRequester: ocpiRequester,
	}
}
