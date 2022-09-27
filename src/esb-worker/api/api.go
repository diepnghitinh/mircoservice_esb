package main

import (
	"esb-worker/configs"
	"esb-worker/internal/core/libs"
	"esb-worker/registry"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"gopkg.in/go-playground/validator.v9"
	"net/http"

	"esb-worker/api/schema"
	"github.com/labstack/echo"
)

var (
	log = logrus.New()
	esb libs.ESB
	channel *amqp.Channel
)

func Routing(c echo.Context) (err error) {

	if esb.GetBroker().IsClosed() == true {
		return c.String(http.StatusBadRequest, "failed")
	}

	var exchange = c.Param("exchange")
	var queue = c.Param("queue")

	body := echo.Map{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	if exchange == "*" {
		exchange = ""
	}

	if queue == "*" {
		queue = ""
	}

	channel.Publish(
		exchange, // exchange
		queue,     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType: "text/plain",
			Body:        jsonData,
		})

	log.Println("A msg pushed to queue: ", queue, ", data: ", body)

	return c.JSON(http.StatusOK, &schema.Response{
		StatusCode: 200,
		Message:    "Routing successful",
	})
}

func main() {

	esb = registry.Registry()

	config := &configs.Config{}
	_ = esb.LoadConfig(config)
	esb.Config = config

	amqp, err := esb.StartServer()
	if err != nil {
		fmt.Errorf("AMQP not connect")
	}

	channel, err = amqp.Channel()
	if err != nil {
		fmt.Errorf("AMQP not create channel")
	}

	e := echo.New()
	e.Validator = &schema.CustomValidator{Validator: validator.New()}
	e.POST("/routing/:exchange/:queue", Routing)
	e.GET("/healthcheck",  func(c echo.Context) error {
		if amqp.IsClosed() == true {
			return c.String(http.StatusBadRequest, "failed")
		}
		return c.String(http.StatusOK, "ok")
	})

	forever := make(chan bool)
		//Reconnect handler
		go esb.HandlerReconnectOnlyConnect()
		e.Logger.Fatal(e.Start(":8080"))
	<-forever
}