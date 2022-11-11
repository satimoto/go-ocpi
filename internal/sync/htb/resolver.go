package htb

import (
	"context"
	"net/http"
	"sync"

	"github.com/satimoto/go-datastore/pkg/connector"
	"github.com/satimoto/go-datastore/pkg/dataimport"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/tariff"
	"github.com/satimoto/go-ocpi/internal/element"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type HtbRepository interface{}

type HtbService struct {
	Repository           HtbRepository
	HTTPRequester        transportation.HTTPRequester
	ConnectorRepository  connector.ConnectorRepository
	DataImportRepository dataimport.DataImportRepository
	ElementResolver      *element.ElementResolver
	TariffRepository     tariff.TariffRepository
	shutdownCtx          context.Context
	waitGroup            *sync.WaitGroup
}

func NewService(repositoryService *db.RepositoryService) *HtbService {
	repo := HtbRepository(repositoryService)

	return &HtbService{
		Repository:           repo,
		HTTPRequester:        &http.Client{},
		ConnectorRepository:  connector.NewRepository(repositoryService),
		DataImportRepository: dataimport.NewRepository(repositoryService),
		ElementResolver:      element.NewResolver(repositoryService),
		TariffRepository:     tariff.NewRepository(repositoryService),
	}
}
