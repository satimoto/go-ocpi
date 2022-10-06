package version

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/version"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type VersionResolver struct {
	Repository    version.VersionRepository
	OcpiRequester *transportation.OcpiRequester
}

func NewResolver(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *VersionResolver {
	return &VersionResolver{
		Repository:    version.NewRepository(repositoryService),
		OcpiRequester: ocpiRequester,
	}
}
