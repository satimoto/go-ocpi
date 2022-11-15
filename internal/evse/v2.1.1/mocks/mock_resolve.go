package mocks

import (
	mocks "github.com/satimoto/go-datastore/pkg/db/mocks"
	evseMocks "github.com/satimoto/go-datastore/pkg/evse/mocks"
	"github.com/satimoto/go-datastore/pkg/util"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1/mocks"
	displaytext "github.com/satimoto/go-ocpi/internal/displaytext/mocks"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	geolocation "github.com/satimoto/go-ocpi/internal/geolocation/mocks"
	image "github.com/satimoto/go-ocpi/internal/image/mocks"
)

func NewResolver(repositoryService *mocks.MockRepositoryService) *evse.EvseResolver {
	recordEvseStatusPeriods := util.GetEnvBool("RECORD_EVSE_STATUS_PERIODS", true)

	return &evse.EvseResolver{
		Repository:              evseMocks.NewRepository(repositoryService),
		ConnectorResolver:       connector.NewResolver(repositoryService),
		DisplayTextResolver:     displaytext.NewResolver(repositoryService),
		GeoLocationResolver:     geolocation.NewResolver(repositoryService),
		ImageResolver:           image.NewResolver(repositoryService),
		RecordEvseStatusPeriods: recordEvseStatusPeriods,
	}
}
