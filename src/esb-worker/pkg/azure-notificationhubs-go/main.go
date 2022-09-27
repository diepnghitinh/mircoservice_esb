package notificationhubs

import (
	"context"
	"time"
)

// NewNotificationHub initializes and returns NotificationHub pointer
func NewNotificationHub(connectionString, hubPath string) *NotificationHub {
	return newNotificationHub(connectionString, hubPath)
}

// NewNotification initalizes and returns Notification pointer
func NewNotification(format NotificationFormat, payload []byte, expiration time.Time) (*Notification, error) {
	return newNotification(format, payload, expiration)
}

// NewRegistration initializes and returns a Notification pointer
func NewRegistration(deviceID string, expirationTime *time.Time, notificationFormat NotificationFormat,
	registrationID string, tags string) *Registration {
	return newRegistration(deviceID, expirationTime, notificationFormat, registrationID, tags)
}

// NewTemplateRegistration initializes and returns a TemplateNotification pointer
func NewTemplateRegistration(deviceID string, expirationTime *time.Time, registrationID string, tags string,
	platform TargetPlatform, template string) *TemplateRegistration {
	return newTemplateRegistration(deviceID, expirationTime, registrationID, tags, platform, template)
}

func DeleteWithRegistrationId(context context.Context, registrationID string) error {
	return deleteWithRegistrationId(context, registrationID)
}
