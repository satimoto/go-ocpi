package itbec

import (
	"context"
	"sync"

	"github.com/satimoto/go-datastore/pkg/connector"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/tariff"
	"github.com/satimoto/go-ocpi/internal/element"
)

type ItBecRepository interface{}

type ItBecService struct {
	Repository           ItBecRepository
	ConnectorRepository  connector.ConnectorRepository
	ElementResolver      *element.ElementResolver
	TariffRepository     tariff.TariffRepository
	shutdownCtx          context.Context
	waitGroup            *sync.WaitGroup
}

func NewService(repositoryService *db.RepositoryService) *ItBecService {
	repo := ItBecRepository(repositoryService)

	return &ItBecService{
		Repository:           repo,
		ConnectorRepository:  connector.NewRepository(repositoryService),
		ElementResolver:      element.NewResolver(repositoryService),
		TariffRepository:     tariff.NewRepository(repositoryService),
	}
}
