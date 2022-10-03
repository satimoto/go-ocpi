package sync

import (
	"context"
	"sync"

	"github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/version"
)

type SyncRepository interface{}

type SyncService struct {
	Repository           SyncRepository
	CredentialRepository credential.CredentialRepository
	VersionResolver      *version.VersionResolver
	syncerHandlers       []*SyncerHandler
	shutdownCtx          context.Context
	waitGroup            *sync.WaitGroup
}

func NewService(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *SyncService {
	repo := SyncRepository(repositoryService)

	return &SyncService{
		Repository:           repo,
		CredentialRepository: credential.NewRepository(repositoryService),
		VersionResolver:      version.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
