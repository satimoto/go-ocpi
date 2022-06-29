package cdr

import (
	"github.com/satimoto/go-datastore/pkg/cdr"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-ocpi/internal/calibration"
	"github.com/satimoto/go-ocpi/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1"
	token "github.com/satimoto/go-ocpi/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type CdrResolver struct {
	Repository             cdr.CdrRepository
	OcpiRequester          *transportation.OcpiRequester
	CalibrationResolver    *calibration.CalibrationResolver
	ChargingPeriodResolver *chargingperiod.ChargingPeriodResolver
	LocationResolver       *location.LocationResolver
	NodeRepository         node.NodeRepository
	TariffResolver         *tariff.TariffResolver
	TokenResolver          *token.TokenResolver
	VersionDetailResolver  *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService) *CdrResolver {
	return NewResolverWithServices(repositoryService, transportation.NewOcpiRequester())
}

func NewResolverWithServices(repositoryService *db.RepositoryService, ocpiRequester *transportation.OcpiRequester) *CdrResolver {
	return &CdrResolver{
		Repository:             cdr.NewRepository(repositoryService),
		OcpiRequester:          ocpiRequester,
		CalibrationResolver:    calibration.NewResolver(repositoryService),
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService),
		NodeRepository:         node.NewRepository(repositoryService),
		TariffResolver:         tariff.NewResolver(repositoryService),
		TokenResolver:          token.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService),
	}
}
