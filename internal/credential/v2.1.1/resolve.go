package credential

import (
	credentialRepository "github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/businessdetail"
	"github.com/satimoto/go-ocpi/internal/service"
	sync "github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/version"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type CredentialResolver struct {
	Repository             credentialRepository.CredentialRepository
	BusinessDetailResolver *businessdetail.BusinessDetailResolver
	OcpiService            *transportation.OcpiService
	SyncService            *sync.SyncService
	VersionResolver        *version.VersionResolver
	VersionDetailResolver  *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *CredentialResolver {
	return &CredentialResolver{
		Repository:             credentialRepository.NewRepository(repositoryService),
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		OcpiService:            services.OcpiService,
		SyncService:            services.SyncService,
		VersionResolver:        version.NewResolver(repositoryService, services.OcpiService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService, services),
	}
}
