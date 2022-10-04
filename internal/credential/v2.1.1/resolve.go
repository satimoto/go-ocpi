package credential

import (
	credentialRepository "github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/businessdetail"
	sync "github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/version"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type CredentialResolver struct {
	Repository             credentialRepository.CredentialRepository
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	OcpiRequester          *transportation.OcpiRequester
	SyncService            *sync.SyncService
	VersionResolver        *version.VersionResolver
	VersionDetailResolver  *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, syncService *sync.SyncService, ocpiRequester *transportation.OcpiRequester) *CredentialResolver {
	return &CredentialResolver{
		Repository:             credentialRepository.NewRepository(repositoryService),
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		OcpiRequester:          ocpiRequester,
		SyncService:            syncService,
		VersionResolver:        version.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionDetailResolver:  versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
