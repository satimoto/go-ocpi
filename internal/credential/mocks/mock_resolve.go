package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	businessdetail "github.com/satimoto/go-ocpi-api/internal/businessdetail/mocks"
	credential "github.com/satimoto/go-ocpi-api/internal/credential"
	sync "github.com/satimoto/go-ocpi-api/internal/sync/mocks"
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
		OcpiRequester:          ocpiRequester,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		SyncResolver:           sync.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionResolver:        version.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionDetailResolver:  versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
