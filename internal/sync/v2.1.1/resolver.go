package sync

import (
	"github.com/satimoto/go-datastore/db"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
)

type SyncRepository interface{}

type SyncResolver struct {
	Repository       SyncRepository
	CdrResolver      *cdr.CdrResolver
	LocationResolver *location.LocationResolver
	SessionResolver  *session.SessionResolver
	TariffResolver   *tariff.TariffResolver
}

func NewResolver(repositoryService *db.RepositoryService) *SyncResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *SyncResolver {
	repo := SyncRepository(repositoryService)

	return &SyncResolver{
		Repository:       repo,
		CdrResolver:      cdr.NewResolverWithServices(repositoryService, ocpiRequester),
		LocationResolver: location.NewResolverWithServices(repositoryService, ocpiRequester),
		SessionResolver:  session.NewResolverWithServices(repositoryService, ocpiRequester),
		TariffResolver:   tariff.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
