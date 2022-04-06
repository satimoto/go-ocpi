package version

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

type VersionRepository interface {
	CreateVersion(ctx context.Context, arg db.CreateVersionParams) (db.Version, error)
	DeleteVersions(ctx context.Context, credentialID int64) error
	GetVersion(ctx context.Context, id int64) (db.Version, error)
	ListVersions(ctx context.Context, credentialID int64) ([]db.Version, error)
}

type VersionResolver struct {
	Repository VersionRepository
	*ocpi.OCPIRequester
}

func NewResolver(repositoryService *db.RepositoryService) *VersionResolver {
	repo := VersionRepository(repositoryService)

	return &VersionResolver{
		Repository:    repo,
		OCPIRequester: ocpi.NewOCPIRequester(),
	}
}
