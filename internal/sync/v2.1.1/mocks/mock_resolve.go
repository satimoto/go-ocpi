package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1/mocks"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1/mocks"
	sync "github.com/satimoto/go-ocpi-api/internal/sync/v2.1.1"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *ocpi.OCPIRequester) *sync.SyncResolver {
	repo := sync.SyncRepository(repositoryService)

	return &sync.SyncResolver{
		Repository:       repo,
		CdrResolver:      cdr.NewResolver(repositoryService, requester),
		LocationResolver: location.NewResolver(repositoryService, requester),
		SessionResolver:  session.NewResolver(repositoryService, requester),
		TariffResolver:   tariff.NewResolver(repositoryService, requester),
	}
}
