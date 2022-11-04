package tokenauthorization

import (
	"log"

	"github.com/appleboy/go-fcm"
	"github.com/satimoto/go-datastore/pkg/db"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/notification"
)

func (r *TokenAuthorizationResolver) SendNotification(user db.User, authorizationID string) {
	dto := notification.CreateTokenAuthorizeNotificationDto(authorizationID)

	r.sendNotification(user, dto)
}

func (r *TokenAuthorizationResolver) sendNotification(user db.User, data notification.NotificationDto) {
	message := &fcm.Message{
		To:               user.DeviceToken,
		ContentAvailable: true,
		Data:             data,
	}

	_, err := r.NotificationService.SendNotificationWithRetry(message, 10)

	if err != nil {
		// TODO: Cancel session?
		metrics.RecordError("OCPI286", "Error sending notification", err)
		log.Printf("OCPI286: Message=%v", message)
	}
}
