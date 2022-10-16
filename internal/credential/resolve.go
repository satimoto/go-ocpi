package credential

import (
	"github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/businessdetail"
	credential_2_1_1 "github.com/satimoto/go-ocpi/internal/credential/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/service"
	"github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/version"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type CredentialResolver struct {
	Repository               credential.CredentialRepository
	OcpiService              *transportation.OcpiService
	SyncService              *sync.SyncService
	BusinessDetailResolver   *businessdetail.BusinessDetailResolver
	CredentialResolver_2_1_1 *credential_2_1_1.CredentialResolver
	VersionResolver          *version.VersionResolver
	VersionDetailResolver    *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *CredentialResolver {
	return &CredentialResolver{
		Repository:               credential.NewRepository(repositoryService),
		OcpiService:              services.OcpiService,
		SyncService:              services.SyncService,
		BusinessDetailResolver:   businessdetail.NewResolver(repositoryService),
		CredentialResolver_2_1_1: credential_2_1_1.NewResolver(repositoryService, services),
		VersionResolver:          version.NewResolver(repositoryService, services.OcpiService),
		VersionDetailResolver:    versiondetail.NewResolver(repositoryService, services),
	}
}
