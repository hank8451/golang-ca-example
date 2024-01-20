package repositories

import (
	"github.com/hank8451/cockroach_ca/cockroach/entities"
	"github.com/labstack/gommon/log"
)

type cockroachFCMMessaging struct {}

func NewCockroachFCMMessaging() CockroachMessaging {
	return &cockroachFCMMessaging{}
}

func (c *cockroachFCMMessaging) PushNotification(m *entities.CockroachPushNotificationDto) error {
	log.Debugf("Pushed FCM notification with data: %v", m)
	return nil
}