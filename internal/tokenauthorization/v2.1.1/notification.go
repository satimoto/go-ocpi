package tokenauthorization

import (
	"log"

	"github.com/appleboy/go-fcm"
	"github.com/satimoto/go-datastore/pkg/db"
	metrics "github.com/satimoto/go-ocpi/internal/metric"
	"github.com/satimoto/go-ocpi/internal/notification"
)

func (r *TokenAuthorizationResolver) SendContentNotification(user db.User, title, body string) {
	if user.DeviceToken.Valid {
		message := &fcm.Message{
			To: user.DeviceToken.String,
			Notification: &fcm.Notification{
				Title: title,
				Body:  body,
			},
		}

		_, err := r.NotificationService.SendNotificationWithRetry(message, 10)

		if err != nil {
			metrics.RecordError("OCPI333", "Error sending notification", err)
			log.Printf("OCPI333: Message=%v", message)
		}
	}
}

func (r *TokenAuthorizationResolver) SendDataNotification(user db.User, authorizationID string) {
	dto := notification.CreateTokenAuthorizeNotificationDto(authorizationID)

	r.sendNotification(user, dto)
}

func (r *TokenAuthorizationResolver) sendNotification(user db.User, data notification.NotificationDto) {
	if user.DeviceToken.Valid {
		message := &fcm.Message{
			To:               user.DeviceToken.String,
			ContentAvailable: true,
			Priority:         "high",
			Data:             data,
		}

		_, err := r.NotificationService.SendNotificationWithRetry(message, 10)

		if err != nil {
			// TODO: Cancel session?
			metrics.RecordError("OCPI286", "Error sending notification", err)
			log.Printf("OCPI286: Message=%v", message)
		}
	}
}
