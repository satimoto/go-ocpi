package mocks

import (
	credentialMocks "github.com/satimoto/go-datastore/pkg/credential/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	businessdetail "github.com/satimoto/go-ocpi/internal/businessdetail/mocks"
	credential "github.com/satimoto/go-ocpi/internal/credential"
	credential_2_1_1 "github.com/satimoto/go-ocpi/internal/credential/v2.1.1/mocks"
	sync "github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
	version "github.com/satimoto/go-ocpi/internal/version/mocks"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, syncService *sync.SyncService, ocpiRequester *transportation.OcpiRequester) *credential.CredentialResolver {
	return &credential.CredentialResolver{
		Repository:               credentialMocks.NewRepository(repositoryService),
		OcpiRequester:            ocpiRequester,
		SyncService:              syncService,
		BusinessDetailResolver:   businessdetail.NewResolver(repositoryService),
		CredentialResolver_2_1_1: credential_2_1_1.NewResolver(repositoryService, syncService, ocpiRequester),
		VersionResolver:          version.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionDetailResolver:    versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
