package notificationhubs

import (
	"bytes"
	"context"
	"path"
)

func deleteWithRegistrationId(ctx context.Context, registrationId string) error {
	return DeleteWithRegistrationId(ctx, registrationId);
}

func (h *NotificationHub) DeleteWithRegistrationId(ctx context.Context, registrationId string) (err error) {
	var (
		regURL  = h.generateAPIURL("registrations")
		method  = postMethod
		payload = ""
		headers = map[string]string{
			"Content-Type": "application/atom+xml;type=entry;charset=utf-8",
			"If-Match": "*",
		}
	)

	method = deleteMethod
	regURL.Path = path.Join(regURL.Path, registrationId)

	_, _, err = h.exec(ctx, method, regURL, headers, bytes.NewBufferString(payload))
	return err
}
