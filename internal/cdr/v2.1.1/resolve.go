package cdr

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/cdr"
	"github.com/satimoto/go-ocpi-api/internal/calibration"
	"github.com/satimoto/go-ocpi-api/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi-api/internal/location/v2.1.1"
	tariff "github.com/satimoto/go-ocpi-api/internal/tariff/v2.1.1"
	token "github.com/satimoto/go-ocpi-api/internal/token/v2.1.1"
	"github.com/satimoto/go-ocpi-api/internal/transportation"
	"github.com/satimoto/go-ocpi-api/internal/versiondetail"
)

type CdrResolver struct {
	Repository             cdr.CdrRepository
	OcpiRequester          *transportation.OcpiRequester
	CalibrationResolver    *calibration.CalibrationResolver
	ChargingPeriodResolver *chargingperiod.ChargingPeriodResolver
	LocationResolver       *location.LocationResolver
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
		TariffResolver:         tariff.NewResolver(repositoryService),
		TokenResolver:          token.NewResolver(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService),
	}
}
