package version

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

type VersionRepository interface {
	CreateVersion(ctx context.Context, arg db.CreateVersionParams) (db.Version, error)
	DeleteVersions(ctx context.Context, credentialID int64) error
	GetVersion(ctx context.Context, id int64) (db.Version, error)
	ListVersions(ctx context.Context, credentialID int64) ([]db.Version, error)
}

type VersionResolver struct {
	Repository    VersionRepository
	OcpiRequester *transportation.OcpiRequester
}

func NewResolver(repositoryService *db.RepositoryService) *VersionResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *VersionResolver {
	repo := VersionRepository(repositoryService)

	return &VersionResolver{
		Repository:    repo,
		OcpiRequester: ocpiRequester,
	}
}
