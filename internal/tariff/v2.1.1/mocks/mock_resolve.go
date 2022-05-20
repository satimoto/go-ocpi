package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	tariffMocks "github.com/satimoto/go-datastore/pkg/tariff/mocks"
	displaytext "github.com/satimoto/go-ocpi-api/internal/displaytext/mocks"
	element "github.com/satimoto/go-ocpi-api/internal/element/mocks"
	energymix "github.com/satimoto/go-ocpi-api/internal/energymix/mocks"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	tariffrestriction "github.com/satimoto/go-ocpi-api/internal/tariffrestriction/mocks"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	versiondetail "github.com/satimoto/go-ocpi-api/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *tariff.TariffResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *mocks.MockRepositoryService, ocpiRequester *transportation.OcpiRequester) *tariff.TariffResolver {
	return &tariff.TariffResolver{
		Repository:                tariffMocks.NewRepository(repositoryService),
		OcpiRequester:             ocpiRequester,
		DisplayTextResolver:       displaytext.NewResolver(repositoryService),
		ElementResolver:           element.NewResolver(repositoryService),
		EnergyMixResolver:         energymix.NewResolver(repositoryService),
		TariffRestrictionResolver: tariffrestriction.NewResolver(repositoryService),
		VersionDetailResolver:     versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
