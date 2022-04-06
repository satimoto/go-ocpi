package sync

import (
	"github.com/satimoto/go-datastore/db"
	cdr "github.com/satimoto/go-ocpi-api/internal/cdr/v2.1.1"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	session "github.com/satimoto/go-ocpi-api/internal/session/v2.1.1"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
)

type SyncRepository interface{}

type SyncResolver struct {
	Repository SyncRepository
	*cdr.CdrResolver
	*location.LocationResolver
	*session.SessionResolver
	*tariff.TariffResolver
}

func NewResolver(repositoryService *db.RepositoryService) *SyncResolver {
	repo := SyncRepository(repositoryService)
	return &SyncResolver{
		Repository:       repo,
		CdrResolver:      cdr.NewResolver(repositoryService),
		LocationResolver: location.NewResolver(repositoryService),
		SessionResolver:  session.NewResolver(repositoryService),
		TariffResolver:   tariff.NewResolver(repositoryService),
	}
}
