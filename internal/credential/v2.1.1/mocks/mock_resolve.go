package mocks

import (
	credentialMocks "github.com/satimoto/go-datastore/pkg/credential/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	businessdetail "github.com/satimoto/go-ocpi-api/internal/businessdetail/mocks"
	credential "github.com/satimoto/go-ocpi-api/internal/credential/v2.1.1"
	sync "github.com/satimoto/go-ocpi-api/internal/sync/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	version "github.com/satimoto/go-ocpi-api/internal/version/mocks"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *credential.CredentialResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *credential.CredentialResolver {
	return &credential.CredentialResolver{
		Repository:             credentialMocks.NewRepository(repositoryService),
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		OcpiRequester:          ocpiRequester,
		SyncResolver:           sync.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionResolver:        version.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionDetailResolver:  versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
