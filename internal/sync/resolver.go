package sync

import (
	"context"
	"sync"

	"github.com/satimoto/go-datastore/pkg/credential"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/party"
	"github.com/satimoto/go-ocpi/internal/sync/htb"
	"github.com/satimoto/go-ocpi/internal/sync/nlcon"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/version"
)

type SyncRepository interface{}

type SyncService struct {
	Repository           SyncRepository
	HtbService           *htb.HtbService
	NlConService         *nlcon.NlConService
	CredentialRepository credential.CredentialRepository
	PartyRepository      party.PartyRepository
	VersionResolver      *version.VersionResolver
	syncerHandlers       []*SyncerHandler
	shutdownCtx          context.Context
	waitGroup            *sync.WaitGroup
	activeSyncs          map[string]bool
	tariffsSyncing       bool
}

func NewService(repositoryService *db.RepositoryService, ocpiService *transportation.OcpiService) *SyncService {
	repo := SyncRepository(repositoryService)

	return &SyncService{
		Repository:           repo,
		HtbService:           htb.NewService(repositoryService),
		NlConService:         nlcon.NewService(repositoryService),
		CredentialRepository: credential.NewRepository(repositoryService),
		PartyRepository:      party.NewRepository(repositoryService),
		VersionResolver:      version.NewResolver(repositoryService, ocpiService),
		activeSyncs:          make(map[string]bool),
		tariffsSyncing:       false,
	}
}
