package feed

import (
	"esb-worker/internal/services/types"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

func GetDefaultPushManageMsg(namespace string, dataArray interface{}) *PushManageMsg{
	return &PushManageMsg{
		Type: "feed.manage",
		Data: &PushManageMsgData{
			Title:  "feed manage create send segment",
			Data: dataArray,
			NameSpace: namespace,
		},
	}
}

func PushToManager(broker *amqp.Channel, accountType types.AccountType, msg *PushManageMsg) {

	transmitData, _ := json.Marshal(PushManageMsg{
		Type: "feed.manage",
		Data: msg.Data,
	})

	broker.Publish(
		"",                        // exchange
		fmt.Sprintf("feed.%s.push.manage.handler", accountType), // routing key
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         transmitData,
		})
}