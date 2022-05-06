package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *versiondetail.VersionDetailResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *versiondetail.VersionDetailResolver {
	repo := versiondetail.VersionDetailRepository(repositoryService)

	return &versiondetail.VersionDetailResolver{
		Repository:    repo,
		OcpiRequester: ocpiRequester,
	}
}
