package version

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/version"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type VersionResolver struct {
	Repository  version.VersionRepository
	OcpiService *transportation.OcpiService
}

func NewResolver(repositoryService *db.RepositoryService, ocpiService *transportation.OcpiService) *VersionResolver {
	return &VersionResolver{
		Repository:  version.NewRepository(repositoryService),
		OcpiService: ocpiService,
	}
}
