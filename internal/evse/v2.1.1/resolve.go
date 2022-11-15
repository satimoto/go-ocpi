package evse

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/evse"
	"github.com/satimoto/go-datastore/pkg/util"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/displaytext"
	"github.com/satimoto/go-ocpi/internal/geolocation"
	"github.com/satimoto/go-ocpi/internal/image"
)

type EvseResolver struct {
	Repository              evse.EvseRepository
	ConnectorResolver       *connector.ConnectorResolver
	DisplayTextResolver     *displaytext.DisplayTextResolver
	GeoLocationResolver     *geolocation.GeoLocationResolver
	ImageResolver           *image.ImageResolver
	RecordEvseStatusPeriods bool
}

func NewResolver(repositoryService *db.RepositoryService) *EvseResolver {
	recordEvseStatusPeriods := util.GetEnvBool("RECORD_EVSE_STATUS_PERIODS", true)

	return &EvseResolver{
		Repository:              evse.NewRepository(repositoryService),
		ConnectorResolver:       connector.NewResolver(repositoryService),
		DisplayTextResolver:     displaytext.NewResolver(repositoryService),
		GeoLocationResolver:     geolocation.NewResolver(repositoryService),
		ImageResolver:           image.NewResolver(repositoryService),
		RecordEvseStatusPeriods: recordEvseStatusPeriods,
	}
}
