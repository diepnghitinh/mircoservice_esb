package driver

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

var AmqpDefaulChannel = &Amqp{}

type Amqp struct {
	CurrentChannel *amqp.Channel
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (p *Amqp) PushToQueue(pushData interface{}, queue string) error {

	jsonData, err := json.Marshal(pushData)
	failOnError(err, "Parse failed")

	// QueueDeclare Handle
	p.CurrentChannel.QueueDeclare(
		queue,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	return p.CurrentChannel.Publish(
		"",    // exchange
		queue,    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         jsonData,
		})
}
