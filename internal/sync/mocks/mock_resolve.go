package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	sync "github.com/satimoto/go-ocpi-api/internal/sync"
	sync2_1_1 "github.com/satimoto/go-ocpi-api/internal/sync/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	version "github.com/satimoto/go-ocpi-api/internal/version/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *sync.SyncResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *sync.SyncResolver {
	repo := sync.SyncRepository(repositoryService)

	return &sync.SyncResolver{
		Repository:        repo,
		Sync2_1_1Resolver: sync2_1_1.NewResolverWithServices(repositoryService, ocpiRequester),
		VersionResolver:   version.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
