package repositories

import "github.com/hank8451/cockroach_ca/cockroach/entities"

type CockroachMessaging interface {
	PushNotification(m *entities.CockroachPushNotificationDto) error
}