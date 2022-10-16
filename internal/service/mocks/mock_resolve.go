package mocks

import (
	"github.com/satimoto/go-datastore/pkg/db/mocks"
	"github.com/satimoto/go-ocpi/internal/async"
	notification "github.com/satimoto/go-ocpi/internal/notification/mocks"
	"github.com/satimoto/go-ocpi/internal/service"
	sync "github.com/satimoto/go-ocpi/internal/sync/mocks"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

func NewService(repositoryService *mocks.MockRepositoryService, notificationService *notification.MockNotificationService, ocpiService *transportation.OcpiService) *service.ServiceResolver {
	asyncService := async.NewService()
	syncService := sync.NewService(repositoryService, ocpiService)

	return &service.ServiceResolver{
		AsyncService:        asyncService,
		NotificationService: notificationService,
		OcpiService:         ocpiService,
		SyncService:         syncService,
	}
}
