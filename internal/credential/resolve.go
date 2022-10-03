package credential

import (
	"github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/businessdetail"
	credential_2_1_1 "github.com/satimoto/go-ocpi/internal/credential/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/version"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type CredentialResolver struct {
	Repository               credential.CredentialRepository
	OcpiRequester            *transportation.OcpiRequester
	BusinessDetailResolver   *businessdetail.BusinessDetailResolver
	CredentialResolver_2_1_1 *credential_2_1_1.CredentialResolver
	SyncService              *sync.SyncService
	VersionResolver          *version.VersionResolver
	VersionDetailResolver    *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, syncService *sync.SyncService, ocpiRequester *transportation.OcpiRequester) *CredentialResolver {
	return &CredentialResolver{
		Repository:               credential.NewRepository(repositoryService),
		OcpiRequester:            ocpiRequester,
		BusinessDetailResolver:   businessdetail.NewResolver(repositoryService),
		CredentialResolver_2_1_1: credential_2_1_1.NewResolver(repositoryService, syncService, ocpiRequester),
		SyncService:              sync.NewService(repositoryService, ocpiRequester),
		VersionResolver:          version.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionDetailResolver:    versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
