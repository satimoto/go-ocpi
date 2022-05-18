package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/version"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *version.VersionResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *version.VersionResolver {
	repo := version.VersionRepository(repositoryService)

	return &version.VersionResolver{
		Repository:    repo,
		OcpiRequester: ocpiRequester,
	}
}
