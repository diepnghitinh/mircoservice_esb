package libs

import (
	"encoding/json"
	config "esb-worker/configs"
	"esb-worker/internal/core/driver"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"net/http"
	"os"
	"time"
)

var log = logrus.New()

type ESB struct {
	ConfigPath string
	Service    string

	Config *config.Config

	//Broker
	Broker *amqp.Channel

	consumers map[string]interface{}
}

type ErrorMsg struct {
	Error string
	Data  string
}

func (esb *ESB) SetConfig(configPath string, service string) error {
	esb.ConfigPath = configPath
	esb.Service = service
	return nil
}

// Load Config from file or enviroment variable
func (esb *ESB) LoadConfig(refConf interface{}) error {
	var (
		c config.File
		e config.Env

		err error
	)

	if esb.ConfigPath != "" {
		_, err = c.NewFromYaml(refConf, esb.ConfigPath)
	} else {
		_, err = e.NewFromEnvironmentFromConfig(refConf)
	}

	return err
}

func (esb *ESB) Init() {}

func (esb *ESB) RegisterTasks(tasks map[string]interface{}) map[string]interface{} {
	return tasks
}

func (esb *ESB) RegisterConsumers(consumers map[string]interface{}) map[string]interface{} {
	esb.consumers = consumers
	return consumers
}

func (esb *ESB) GetConfig() (*config.Config, error) {

	var _config = &config.Config{}
	err := esb.LoadConfig(_config)
	esb.Config = _config

	return esb.Config, err
}

//Broker setup
var (
	rabbitConn       *amqp.Connection
	rabbitCloseError chan *amqp.Error
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// Server start connect amqp
func (esb *ESB) StartServer() (*amqp.Connection, error) {
	var err error
	rabbitConn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/%s", esb.Config.Broker.User, esb.Config.Broker.Pass,
		esb.Config.Broker.Host, esb.Config.Broker.Port, esb.Config.Broker.Vhost))
	return rabbitConn, err
}

func (esb *ESB) GetBroker() *amqp.Connection {
	return rabbitConn
}

func (esb *ESB) ConsumerInit() error {
	// Create the rabbitmq error channel
	brokerConnect, err := esb.StartServer()
	if err != nil {
		log.Error(err)
		return err
	}

	esb.Broker, err = brokerConnect.Channel()
	failOnError(err, "Failed to open a channel")

	if esb.Config.AMQP.Exchange != "" {
		err = esb.Broker.ExchangeDeclare(
			esb.Config.AMQP.Exchange,     // name
			esb.Config.AMQP.ExchangeType, // type
			true,                         // durable
			false,                        // auto-deleted
			false,                        // internal
			false,                        // no-wait
			nil,                          // arguments
		)
		failOnError(err, "Failed to declare an exchange")
	}

	// QueueDeclare Handle
	q, err := esb.Broker.QueueDeclare(
		esb.Config.DefaultQueue,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	if esb.Config.AMQP.Exchange != "" {
		err = esb.Broker.QueueBind(
			q.Name,
			esb.Config.AMQP.BindingKey,
			esb.Config.AMQP.Exchange,
			false,
			nil,
		)
		failOnError(err, "Failed bind queue")
	}

	Hostname, err := os.Hostname()
	esb.Broker.Qos(esb.Config.AMQP.PrefetchCount, 0, false)
	msgs, err := esb.Broker.Consume(
		q.Name,   // queue
		Hostname, // consumer
		false,    // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)

	go func() {

		serviceInstance := esb.GetService(esb.Service)
		serviceInstance.(IConsumer).Init(esb)

		for d := range msgs {

			go func(d amqp.Delivery) {
				result := serviceInstance.(IConsumer).Handler(d)
				if result == nil {
					d.Ack(false)
				} else {
					// QueueDeclare Handle
					q, err := esb.Broker.QueueDeclare(
						fmt.Sprintf("%s_error", esb.Config.DefaultQueue),
						true,  // durable
						false, // delete when unused
						false, // exclusive
						false, // no-wait
						nil,   // arguments
					)
					failOnError(err, "Failed to declare a queue")

					//Define error msg
					ErrorMsg := &ErrorMsg{}
					ErrorMsg.Error = result.Error()
					ErrorMsg.Data = string(d.Body)

					body, err := json.Marshal(ErrorMsg)

					err = esb.Broker.Publish(
						"",     // exchange
						q.Name, // routing key
						false,  // mandatory
						false,  // immediate
						amqp.Publishing{
							ContentType: "text/plain",
							Body:        []byte(body),
						})
					failOnError(err, "Failed to publish a message")

					d.Ack(false)
				}
			}(d)

		}
	}()

	log.Infof("Launching a worker with the following settings:")
	log.Infof("- Service: %s", esb.Service)
	log.Infof("- Broker: %s", esb.Config.Broker.Host)
	log.Infof("- DefaultQueue: %s", esb.Config.DefaultQueue)

	if esb.Config.AMQP != nil {
		log.Infof("- AMQP: %s", esb.Config.AMQP.Exchange)
		log.Infof("  - Exchange: %s", esb.Config.AMQP.Exchange)
		log.Infof("  - ExchangeType: %s", esb.Config.AMQP.ExchangeType)
		log.Infof("  - BindingKey: %s", esb.Config.AMQP.BindingKey)
		log.Infof("  - PrefetchCount: %d", esb.Config.AMQP.PrefetchCount)
	}

	log.Infof(" [*] Waiting for logs ...")

	return nil
}

func (esb *ESB) HandlerReconnect() {
	rabbitCloseError = rabbitConn.NotifyClose(make(chan *amqp.Error))
	for {
		select {
		case err := <-rabbitCloseError:
			if err == nil {
			}
			log.Warn("Broker is disconnect. Reconnecting ...")

			if esb.ConsumerInit() != nil {
				time.Sleep(15 * time.Second)
			} else {
				rabbitCloseError = rabbitConn.NotifyClose(make(chan *amqp.Error))
			}
		}
	}
}

func (esb *ESB) HandlerReconnectOnlyConnect() {
	rabbitCloseError = rabbitConn.NotifyClose(make(chan *amqp.Error))
	for {
		select {
		case err := <-rabbitCloseError:
			if err == nil {
			}
			log.Warn("Broker is disconnect. Reconnecting ...")

			_, err_start := esb.StartServer()
			if err_start != nil {
				time.Sleep(15 * time.Second)
			} else {
				log.Info("Broker is Connected ...")
				rabbitCloseError = rabbitConn.NotifyClose(make(chan *amqp.Error))
			}
		}
	}
}

// Get Consumer
func (esb *ESB) GetConsumer() error {

	if esb.Service == "" {
		return fmt.Errorf("Service not found")
	}

	forever := make(chan bool)
	err := esb.ConsumerInit()
	if err != nil {
		return err
	}

	//Reconnect handler
	go esb.HandlerReconnect()
	go esb.Healthcheck()
	<-forever

	return nil
}

func (esb *ESB) Healthcheck() {

	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {

		check := true

		// Check database
		if driver.Connection != nil {
			err := driver.Connection.DB().Ping()

			if err != nil {
				check = false
			}
		}

		if check == true {
			w.WriteHeader(200)
			log.Info("healthcheck ok")
			fmt.Fprintf(w, "Ok")
			return
		}

		w.WriteHeader(400)
		log.Info("healthcheck fail")
		fmt.Fprintf(w, "Not healthcheck")
		return

	})

	http.ListenAndServe(":8080", nil)
}

func (esb *ESB) GetService(name string) interface{} {
	serviceInstance := esb.consumers[name]
	return serviceInstance
}
