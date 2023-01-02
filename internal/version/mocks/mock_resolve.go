package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	versionMocks "github.com/satimoto/go-datastore/pkg/version/mocks"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/version"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, ocpiService *transportation.OcpiService) *version.VersionResolver {
	return &version.VersionResolver{
		Repository:  versionMocks.NewRepository(repositoryService),
		OcpiService: ocpiService,
	}
}
