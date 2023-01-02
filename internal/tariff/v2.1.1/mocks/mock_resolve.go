package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	tariffMocks "github.com/satimoto/go-datastore/pkg/tariff/mocks"
	displaytext "github.com/satimoto/go-ocpi/internal/displaytext/mocks"
	element "github.com/satimoto/go-ocpi/internal/element/mocks"
	energymix "github.com/satimoto/go-ocpi/internal/energymix/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1"
	tariffrestriction "github.com/satimoto/go-ocpi/internal/tariffrestriction/mocks"
	versiondetail "github.com/satimoto/go-ocpi/internal/versiondetail/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService, services * service.ServiceResolver) *tariff.TariffResolver {
	return &tariff.TariffResolver{
		Repository:                tariffMocks.NewRepository(repositoryService),
		OcpiService:               services.OcpiService,
		DisplayTextResolver:       displaytext.NewResolver(repositoryService),
		ElementResolver:           element.NewResolver(repositoryService),
		EnergyMixResolver:         energymix.NewResolver(repositoryService),
		TariffRestrictionResolver: tariffrestriction.NewResolver(repositoryService),
		VersionDetailResolver:     versiondetail.NewResolver(repositoryService, services),
	}
}
