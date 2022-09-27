package configs

//version 0.2

import (
	"crypto/tls"
	_ "esb-worker/pkg/envconfig"
	"fmt"
	"strings"
	"time"
)

var (
	// Start with sensible default values
	defaultCnf = &Config{
		Driver: &DriverConfig{
			SslMode: "disable",
		},
		Broker: &BrokerConfig{
			Host:  "localhost",
			User:  "guest",
			Pass:  "guest",
			Port:  "5672",
			Vhost: "/",
		},
		DefaultQueue: "ami_default_tasks",
		AMQP: &AMQPConfig{
			Exchange:      "ami_default_exchange",
			ExchangeType:  "direct",
			BindingKey:    "",
			PrefetchCount: 3,
		},
	}

	reloadDelay = time.Second * 10
)

// QueueBindingArgs arguments which are used when binding to the exchange
type QueueBindingArgs map[string]interface{}

type Broker struct {
	Broker string `yaml:"broker" envconfig:"BROKER"`
}

// Config holds all configuration for our program
type Config struct {
	Driver          *DriverConfig `yaml:"driver"`
	Broker          *BrokerConfig `yaml:"broker"`
	DefaultQueue    string        `yaml:"default_queue" envconfig:"DEFAULT_QUEUE"`
	ResultsExpireIn int           `yaml:"results_expire_in" envconfig:"RESULTS_EXPIRE_IN"`
	AMQP            *AMQPConfig   `yaml:"amqp"`
	TLSConfig       *tls.Config
}

type BrokerConfig struct {
	Host  string `yaml:"host" envconfig:"RABBIT_CONNECTION"`
	User  string `yaml:"user" envconfig:"RABBIT_CLIENT"`
	Pass  string `yaml:"pass" envconfig:"RABBIT_PASSWORD"`
	Port  string `yaml:"port" envconfig:"RABBIT_POST"`
	Vhost string `yaml:"db" envconfig:"RABBIT_VHOST"`
}

// Driver config
type DriverConfig struct {
	Host    string `yaml:"host" envconfig:"DB_HOST"`
	User    string `yaml:"user" envconfig:"DB_USER"`
	Pass    string `yaml:"pass" envconfig:"DB_PASSWORD"`
	Db      string `yaml:"db" envconfig:"DB_NAME"`
	Port    string `yaml:"port" envconfig:"DB_PORT"`
	Schema  string `yaml:"schema" envconfig:"DB_SCHEMA"`
	MaxConn string `yaml:"max_conn" envconfig:"DB_MAX_CONN"`
	SslMode string `yaml:"ssl_mode" envconfig:"DB_SSL_MODE"`
}

// AMQPConfig wraps RabbitMQ related configuration
type AMQPConfig struct {
	Exchange         string           `yaml:"exchange" envconfig:"AMQP_EXCHANGE"`
	ExchangeType     string           `yaml:"exchange_type" envconfig:"AMQP_EXCHANGE_TYPE"`
	QueueBindingArgs QueueBindingArgs `yaml:"queue_binding_args" envconfig:"AMQP_QUEUE_BINDING_ARGS"`
	BindingKey       string           `yaml:"binding_key" envconfig:"AMQP_BINDING_KEY"`
	PrefetchCount    int              `yaml:"prefetch_count" envconfig:"AMQP_PREFETCH_COUNT"`
}

// Decode from yaml to map (any field whose type or pointer-to-type implements
// envconfig.Decoder can control its own deserialization)
func (args *QueueBindingArgs) Decode(value string) error {
	pairs := strings.Split(value, ",")
	mp := make(map[string]interface{}, len(pairs))
	for _, pair := range pairs {
		kvpair := strings.Split(pair, ":")
		if len(kvpair) != 2 {
			return fmt.Errorf("invalid map item: %q", pair)
		}
		mp[kvpair[0]] = kvpair[1]
	}
	*args = QueueBindingArgs(mp)
	return nil
}
