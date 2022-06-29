package sync

import (
	"github.com/satimoto/go-datastore/pkg/db"
	cdr "github.com/satimoto/go-ocpi/internal/cdr/v2.1.1"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	session "github.com/satimoto/go-ocpi/internal/session/v2.1.1"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type SyncResolver struct {
	CdrResolver      *cdr.CdrResolver
	LocationResolver *location.LocationResolver
	SessionResolver  *session.SessionResolver
	TariffResolver   *tariff.TariffResolver
}

func NewResolver(repositoryService *db.RepositoryService) *SyncResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *SyncResolver {
	return &SyncResolver{
		CdrResolver:      cdr.NewResolverWithServices(repositoryService, ocpiRequester),
		LocationResolver: location.NewResolverWithServices(repositoryService, ocpiRequester),
		SessionResolver:  session.NewResolverWithServices(repositoryService, ocpiRequester),
		TariffResolver:   tariff.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
