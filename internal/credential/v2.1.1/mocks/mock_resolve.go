package mocks

import (
	credentialMocks "github.com/satimoto/go-datastore/pkg/credential/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	businessdetail "github.com/satimoto/go-ocpi/internal/businessdetail/mocks"
	credential "github.com/satimoto/go-ocpi/internal/credential/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/service"
	version "github.com/satimoto/go-ocpi/internal/version/mocks"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services *service.ServiceResolver) *credential.CredentialResolver {
	return &credential.CredentialResolver{
		Repository:             credentialMocks.NewRepository(repositoryService),
		BusinessDetailResolver: businessdetail.NewResolver(repositoryService),
		OcpiService:            services.OcpiService,
		SyncService:            services.SyncService,
		VersionResolver:        version.NewResolver(repositoryService, services.OcpiService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService, services),
	}
}
