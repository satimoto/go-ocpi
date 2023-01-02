package mocks

import (
	credential "github.com/satimoto/go-datastore/pkg/credential/mocks"
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	sync "github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
	version "github.com/satimoto/go-ocpi/internal/version/mocks"
)

func NewService(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiService) *sync.SyncService {
	repo := sync.SyncRepository(repositoryService)

	return &sync.SyncService{
		Repository:           repo,
		CredentialRepository: credential.NewRepository(repositoryService),
		VersionResolver:      version.NewResolver(repositoryService, ocpiRequester),
	}
}
