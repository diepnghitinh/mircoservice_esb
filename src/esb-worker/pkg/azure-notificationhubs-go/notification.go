package notificationhubs

import (
	"fmt"
	"time"
)

type (
	// Notification is a message that can be sent through the hub
	Notification struct {
		Format  NotificationFormat
		Payload []byte
		expiration time.Time
	}
)

// newNotification initializes and returns a Notification pointer
func newNotification(format NotificationFormat, payload []byte, expiration time.Time) (*Notification, error) {
	if !format.IsValid() {
		return nil, fmt.Errorf("unknown format '%s'", format)
	}

	return &Notification{format, payload, expiration}, nil
}

// String returns Notification string representation
func (n *Notification) String() string {
	return fmt.Sprintf("&{%s %s}", n.Format, string(n.Payload))
}
