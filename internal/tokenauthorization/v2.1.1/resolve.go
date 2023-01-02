package tokenauthorization

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/location"
	"github.com/satimoto/go-datastore/pkg/node"
	"github.com/satimoto/go-datastore/pkg/session"
	"github.com/satimoto/go-datastore/pkg/tariff"
	"github.com/satimoto/go-datastore/pkg/tokenauthorization"
	"github.com/satimoto/go-datastore/pkg/user"
	"github.com/satimoto/go-ocpi/internal/async"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/notification"
	"github.com/satimoto/go-ocpi/internal/service"
	"github.com/satimoto/go-ocpi/internal/transportation"
	"github.com/satimoto/go-ocpi/internal/versiondetail"
)

type TokenAuthorizationResolver struct {
	Repository            tokenauthorization.TokenAuthorizationRepository
	OcpiService           *transportation.OcpiService
	AsyncService          *async.AsyncService
	NotificationService   notification.Notification
	ConnectorResolver     *connector.ConnectorResolver
	EvseResolver          *evse.EvseResolver
	LocationRepository    location.LocationRepository
	NodeRepository        node.NodeRepository
	SessionRepository     session.SessionRepository
	TariffRespository     tariff.TariffRepository
	UserRepository        user.UserRepository
	VersionDetailResolver *versiondetail.VersionDetailResolver
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *TokenAuthorizationResolver {
	return &TokenAuthorizationResolver{
		Repository:            tokenauthorization.NewRepository(repositoryService),
		OcpiService:           services.OcpiService,
		AsyncService:          services.AsyncService,
		NotificationService:   services.NotificationService,
		ConnectorResolver:     connector.NewResolver(repositoryService),
		EvseResolver:          evse.NewResolver(repositoryService, services),
		LocationRepository:    location.NewRepository(repositoryService),
		NodeRepository:        node.NewRepository(repositoryService),
		SessionRepository:     session.NewRepository(repositoryService),
		TariffRespository:     tariff.NewRepository(repositoryService),
		UserRepository:        user.NewRepository(repositoryService),
		VersionDetailResolver: versiondetail.NewResolver(repositoryService, services),
	}
}
