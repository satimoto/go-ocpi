package tariff

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/tariff"
	"github.com/satimoto/go-ocpi/internal/displaytext"
	"github.com/satimoto/go-ocpi/internal/element"
	"github.com/satimoto/go-ocpi/internal/energymix"
	"github.com/satimoto/go-ocpi/internal/tariffrestriction"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type TariffResolver struct {
	Repository                tariff.TariffRepository
	OcpiRequester             *transportation.OcpiRequester
	DisplayTextResolver       *displaytext.DisplayTextResolver
	ElementResolver           *element.ElementResolver
	EnergyMixResolver         *energymix.EnergyMixResolver
	TariffRestrictionResolver *tariffrestriction.TariffRestrictionResolver
	VersionDetailResolver     *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *TariffResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *TariffResolver {
	return &TariffResolver{
		Repository:                tariff.NewRepository(repositoryService),
		OcpiRequester:             ocpiRequester,
		DisplayTextResolver:       displaytext.NewResolver(repositoryService),
		ElementResolver:           element.NewResolver(repositoryService),
		EnergyMixResolver:         energymix.NewResolver(repositoryService),
		TariffRestrictionResolver: tariffrestriction.NewResolver(repositoryService),
		VersionDetailResolver:     versiondetail.NewResolverWithServices(repositoryService, ocpiRequester),
	}
}
