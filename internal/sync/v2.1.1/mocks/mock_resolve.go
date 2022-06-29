package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	cdr "github.com/satimoto/go-ocpi/internal/cdr/v2.1.1/mocks"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1/mocks"
	session "github.com/satimoto/go-ocpi/internal/session/v2.1.1/mocks"
	sync "github.com/satimoto/go-ocpi/internal/sync/v2.1.1"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1/mocks"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *sync.SyncResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *sync.SyncResolver {
	return &sync.SyncResolver{
		CdrResolver:      cdr.NewResolverWithServices(repositoryService, ocpiRequester),
		LocationResolver: location.NewResolverWithServices(repositoryService, ocpiRequester),
		SessionResolver:  session.NewResolverWithServices(repositoryService, ocpiRequester),
		TariffResolver:   tariff.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
