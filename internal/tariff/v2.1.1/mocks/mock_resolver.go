package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	displaytext "github.com/satimoto/go-ocpi-api/internal/displaytext/mocks"
	element "github.com/satimoto/go-ocpi-api/internal/element/mocks"
	energymix "github.com/satimoto/go-ocpi-api/internal/energymix/mocks"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	tariffrestriction "github.com/satimoto/go-ocpi-api/internal/tariffrestriction/mocks"
	"github.com/satimoto/go-ocpi-api/internal/util"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *util.OCPIRequester) *tariff.TariffResolver {
	repo := tariff.TariffRepository(repositoryService)

	return &tariff.TariffResolver{
		Repository:                repo,
		DisplayTextResolver:       displaytext.NewResolver(repositoryService),
		ElementResolver:           element.NewResolver(repositoryService),
		EnergyMixResolver:         energymix.NewResolver(repositoryService),
		TariffRestrictionResolver: tariffrestriction.NewResolver(repositoryService),
	}
}
