package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	businessdetail "github.com/satimoto/go-ocpi-api/internal/businessdetail/mocks"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	version "github.com/satimoto/go-ocpi-api/internal/version/mocks"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *credential.CredentialResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *credential.CredentialResolver {
	repo := credential.CredentialRepository(repositoryService)

	return &credential.CredentialResolver{
		Repository:             repo,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		VersionResolver:        version.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionDetailResolver:  versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
