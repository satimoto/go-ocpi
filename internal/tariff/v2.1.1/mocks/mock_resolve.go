package mocks

import (
	mocks "github.com/satimoto/go-datastore-mocks/db"
	displaytext "github.com/satimoto/go-ocpi-api/internal/displaytext/mocks"
	element "github.com/satimoto/go-ocpi-api/internal/element/mocks"
	energymix "github.com/satimoto/go-ocpi-api/internal/energymix/mocks"
	"github.com/satimoto/go-ocpi-api/internal/ocpi"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	tariffrestriction "github.com/satimoto/go-ocpi-api/internal/tariffrestriction/mocks"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/v2.1.1/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, requester *ocpi.OCPIRequester) *tariff.TariffResolver {
	repo := tariff.TariffRepository(repositoryService)

	return &tariff.TariffResolver{
		Repository:                repo,
		DisplayTextResolver:       displaytext.NewResolver(repositoryService),
		ElementResolver:           element.NewResolver(repositoryService),
		EnergyMixResolver:         energymix.NewResolver(repositoryService),
		TariffRestrictionResolver: tariffrestriction.NewResolver(repositoryService),
		VersionDetailResolver:     versiondetail.NewResolver(repositoryService, requester),
	}
}
