package service

import (
	"os"

	"github.com/satimoto/go-datastore/pkg/db"
	"github.com/satimoto/go-ocpi/internal/async"
	"github.com/satimoto/go-ocpi/internal/notification"
	"github.com/satimoto/go-ocpi/internal/sync"
	"github.com/satimoto/go-ocpi/internal/transportation"
)

type ServiceResolver struct {
	AsyncService        *async.AsyncService
	NotificationService notification.Notification
	OcpiService         *transportation.OcpiService
	SyncService         *sync.SyncService
}

func NewService(repositoryService *db.RepositoryService) *ServiceResolver {
	asyncService := async.NewService()
	notificationService := notification.NewService(os.Getenv("FCM_API_KEY"))
	ocpiService := transportation.NewOcpiService()
	syncService := sync.NewService(repositoryService, ocpiService)

	return &ServiceResolver{
		AsyncService:        asyncService,
		OcpiService:         ocpiService,
		SyncService:         syncService,
		NotificationService: notificationService,
	}
}
