package sync

import (
	"context"
	"sync"

	"github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/sync/htb"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/version"
)

type SyncRepository interface{}

type SyncService struct {
	Repository           SyncRepository
	HtbService           *htb.HtbService
	CredentialRepository credential.CredentialRepository
	VersionResolver      *version.VersionResolver
	syncerHandlers       []*SyncerHandler
	shutdownCtx          context.Context
	waitGroup            *sync.WaitGroup
	activeSyncs          map[string]bool
}

func NewService(repositoryService *db.RepositoryService, ocpiService *transportation.OcpiService) *SyncService {
	repo := SyncRepository(repositoryService)

	return &SyncService{
		Repository:           repo,
		HtbService:           htb.NewService(repositoryService),
		CredentialRepository: credential.NewRepository(repositoryService),
		VersionResolver:      version.NewResolver(repositoryService, ocpiService),
		activeSyncs:          make(map[string]bool),
	}
}
