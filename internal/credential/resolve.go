package credential

import (
	"github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi-api/internal/businessdetail"
	"github.com/satimoto/go-ocpi-api/internal/sync"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/version"
	"github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

type CredentialResolver struct {
	Repository             credential.CredentialRepository
	OcpiRequester          *transportation.OcpiRequester
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	SyncResolver           *sync.SyncResolver
	VersionResolver        *version.VersionResolver
	VersionDetailResolver  *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CredentialResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *CredentialResolver {
	return &CredentialResolver{
		Repository:             credential.NewRepository(repositoryService),
		OcpiRequester:          ocpiRequester,
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		SyncResolver:           sync.NewResolver(repositoryService),
		VersionResolver:        version.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService),
	}
}
