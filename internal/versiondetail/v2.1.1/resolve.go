package versiondetail

import (
	"context"

	"github.com/satimoto/go-datastore/db"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
)

type VersionDetailRepository interface {
	CreateVersionEndpoint(ctx context.Context, arg db.CreateVersionEndpointParams) (db.VersionEndpoint, error)
	DeleteVersionEndpoints(ctx context.Context, versionID int64) error
	GetVersionEndpoint(ctx context.Context, id int64) (db.VersionEndpoint, error)
	GetVersionEndpointByIdentity(ctx context.Context, arg db.GetVersionEndpointByIdentityParams) (db.VersionEndpoint, error)
	ListVersionEndpoints(ctx context.Context, versionID int64) ([]db.VersionEndpoint, error)
}

type VersionDetailResolver struct {
	Repository VersionDetailRepository
	*ocpi.OCPIRequester
}

func NewResolver(repositoryService *db.RepositoryService) *VersionDetailResolver {
	repo := VersionDetailRepository(repositoryService)

	return &VersionDetailResolver{
		Repository:    repo,
		OCPIRequester: ocpi.NewOCPIRequester(),
	}
}
