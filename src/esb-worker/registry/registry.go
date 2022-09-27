package registry

import (
	mail_handler "esb-worker/cmd/mail.handler"
	"esb-worker/internal/core/libs"
)

func Registry() libs.ESB {

	var esb = libs.ESB{}
	esb.Init()
	esb.RegisterConsumers(map[string]interface{}{
		"mail.handler": mail_handler.Init(),
	})
	return esb
}
