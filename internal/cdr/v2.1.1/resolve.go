package cdr

import (
	"github.com/satimoto/go-datastore/pkg/cdr"
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-datastore/pkg/session"
	"github.com/satimoto/go-datastore/pkg/token"
	"github.com/satimoto/go-ocpi/internal/calibration"
	"github.com/satimoto/go-ocpi/internal/chargingperiod"
	location "github.com/satimoto/go-ocpi/internal/location/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/service"
	tariff "github.com/satimoto/go-ocpi/internal/tariff/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type CdrResolver struct {
	Repository             cdr.CdrRepository
	OcpiService            *transportation.OcpiService
	CalibrationResolver    *calibration.CalibrationResolver
	ChargingPeriodResolver *chargingperiod.ChargingPeriodResolver
	LocationResolver       *location.LocationResolver
	NodeRepository         node.NodeRepository
	SessionRepository      session.SessionRepository
	TariffResolver         *tariff.TariffResolver
	TokenRepository        token.TokenRepository
	VersionDetailResolver  *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *CdrResolver {
	return &CdrResolver{
		Repository:             cdr.NewRepository(repositoryService),
		OcpiService:            services.OcpiService,
		CalibrationResolver:    calibration.NewResolver(repositoryService),
		ChargingPeriodResolver: chargingperiod.NewResolver(repositoryService),
		LocationResolver:       location.NewResolver(repositoryService, services),
		NodeRepository:         node.NewRepository(repositoryService),
		SessionRepository:      session.NewRepository(repositoryService),
		TariffResolver:         tariff.NewResolver(repositoryService, services),
		TokenRepository:        token.NewRepository(repositoryService),
		VersionDetailResolver:  versiondetail.NewResolver(repositoryService, services),
	}
}
