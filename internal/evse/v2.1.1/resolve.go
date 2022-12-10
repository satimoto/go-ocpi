package evse

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/evse"
	"github.com/satimoto/go-datastore/pkg/location"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-datastore/pkg/session"
	"github.com/satimoto/go-datastore/pkg/tariff"
	"github.com/satimoto/go-datastore/pkg/tokenauthorization"
	"github.com/satimoto/go-datastore/pkg/util"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/displaytext"
	"github.com/satimoto/go-ocpi/internal/geolocation"
	"github.com/satimoto/go-ocpi/internal/image"
	"github.com/satimoto/go-ocpi/internal/service"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type EvseResolver struct {
	Repository                   evse.EvseRepository
	OcpiService                  *transportation.OcpiService
	ConnectorResolver            *connector.ConnectorResolver
	DisplayTextResolver          *displaytext.DisplayTextResolver
	GeoLocationResolver          *geolocation.GeoLocationResolver
	ImageResolver                *image.ImageResolver
	LocationRepository           location.LocationRepository
	NodeRepository               node.NodeRepository
	SessionRepository            session.SessionRepository
	TariffRespository            tariff.TariffRepository
	TokenAuthorizationRepository tokenauthorization.TokenAuthorizationRepository
	VersionDetailResolver        *versiondetail.VersionDetailResolver
	RecordEvseStatusPeriods      bool
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *EvseResolver {
	recordEvseStatusPeriods := util.GetEnvBool("RECORD_EVSE_STATUS_PERIODS", true)

	return &EvseResolver{
		Repository:                   evse.NewRepository(repositoryService),
		OcpiService:                  services.OcpiService,
		ConnectorResolver:            connector.NewResolver(repositoryService),
		DisplayTextResolver:          displaytext.NewResolver(repositoryService),
		GeoLocationResolver:          geolocation.NewResolver(repositoryService),
		ImageResolver:                image.NewResolver(repositoryService),
		LocationRepository:           location.NewRepository(repositoryService),
		NodeRepository:               node.NewRepository(repositoryService),
		SessionRepository:            session.NewRepository(repositoryService),
		TariffRespository:            tariff.NewRepository(repositoryService),
		TokenAuthorizationRepository: tokenauthorization.NewRepository(repositoryService),
		VersionDetailResolver:        versiondetail.NewResolver(repositoryService, services),
		RecordEvseStatusPeriods:      recordEvseStatusPeriods,
	}
}
