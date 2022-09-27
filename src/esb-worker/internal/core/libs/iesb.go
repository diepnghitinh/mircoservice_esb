package libs

import (
	"github.com/streadway/amqp"
)

type IConsumer interface {
	Init(esb *ESB)
	Handler(d amqp.Delivery) error
}