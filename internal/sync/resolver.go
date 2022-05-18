package sync

import (
	"github.com/satimoto/go-datastore/pkg/db"
	sync2_1_1 "github.com/satimoto/go-ocpi-api/internal/sync/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/version"
)

type SyncRepository interface{}

type SyncResolver struct {
	Repository        SyncRepository
	Sync2_1_1Resolver *sync2_1_1.SyncResolver
	VersionResolver   *version.VersionResolver
}

func NewResolver(repositoryService *db.RepositoryService) *SyncResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *SyncResolver {
	repo := SyncRepository(repositoryService)

	return &SyncResolver{
		Repository:        repo,
		Sync2_1_1Resolver: sync2_1_1.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionResolver:   version.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
