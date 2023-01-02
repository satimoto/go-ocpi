package tokenauthorization

import (
	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-datastore/pkg/tokenauthorization"
	"github.com/satimoto/go-datastore/pkg/user"
	"github.com/satimoto/go-ocpi/internal/async"
	connector "github.com/satimoto/go-ocpi/internal/connector/v2.1.1"
	evse "github.com/satimoto/go-ocpi/internal/evse/v2.1.1"
	"github.com/satimoto/go-ocpi/internal/notification"
	"github.com/satimoto/go-ocpi/internal/service"
)

type TokenAuthorizationResolver struct {
	Repository          tokenauthorization.TokenAuthorizationRepository
	AsyncService        *async.AsyncService
	NotificationService notification.Notification
	ConnectorResolver   *connector.ConnectorResolver
	EvseResolver        *evse.EvseResolver
	UserRepository      user.UserRepository
}

func NewResolver(repositoryService *db.RepositoryService, services *service.ServiceResolver) *TokenAuthorizationResolver {
	return &TokenAuthorizationResolver{
		Repository:          tokenauthorization.NewRepository(repositoryService),
		AsyncService:        services.AsyncService,
		NotificationService: services.NotificationService,
		ConnectorResolver:   connector.NewResolver(repositoryService),
		EvseResolver:        evse.NewResolver(repositoryService),
		UserRepository:      user.NewRepository(repositoryService),
	}
}
