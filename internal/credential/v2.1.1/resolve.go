package credential

import (
	credentialRepository "github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	sync "github.com/satimoto/go-ocpi-api/internal/sync/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/version"
	"github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

type CredentialResolver struct {
	Repository             credentialRepository.CredentialRepository
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	OcpiRequester          *transportation.OcpiRequester
	SyncResolver           *sync.SyncResolver
	VersionResolver        *version.VersionResolver
	VersionDetailResolver  *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CredentialResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *CredentialResolver {
	return &CredentialResolver{
		Repository:             credentialRepository.NewRepository(repositoryService),
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		OcpiRequester:          ocpiRequester,
		SyncResolver:           sync.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionResolver:        version.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionDetailResolver:  versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
