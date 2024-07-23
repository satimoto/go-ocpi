package deise

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

type DeIseRepository interface{}

type DeIseService struct {
	Repository           DeIseRepository
	HTTPRequester        transportation.HTTPRequester
	ConnectorRepository  connector.ConnectorRepository
	DataImportRepository dataimport.DataImportRepository
	ElementResolver      *element.ElementResolver
	TariffRepository     tariff.TariffRepository
	shutdownCtx          context.Context
	waitGroup            *sync.WaitGroup
}

func NewService(repositoryService *db.RepositoryService) *DeIseService {
	repo := DeIseRepository(repositoryService)

	return &DeIseService{
		Repository:           repo,
		HTTPRequester:        &http.Client{},
		ConnectorRepository:  connector.NewRepository(repositoryService),
		DataImportRepository: dataimport.NewRepository(repositoryService),
		ElementResolver:      element.NewResolver(repositoryService),
		TariffRepository:     tariff.NewRepository(repositoryService),
	}
}
