package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	versiondetailMocks "github.com/satimoto/go-datastore/pkg/versiondetail/mocks"
	"github.com/satimoto/go-ocpi/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *versiondetail.VersionDetailResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *versiondetail.VersionDetailResolver {
	return &versiondetail.VersionDetailResolver{
		Repository:    versiondetailMocks.NewRepository(repositoryService),
		OcpiRequester: ocpiRequester,
	}
}
