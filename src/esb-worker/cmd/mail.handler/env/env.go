package env

import (
	_ "esb-worker/pkg/envconfig"
)

// Config holds all configuration for our program
type Config struct {
	Services *Services `yaml:"services"`
}

type Services struct {
	Host       string `yaml:"SMTP_HOST" envconfig:"SMTP_HOST"`
	Port       int `yaml:"SMTP_PORT" envconfig:"SMTP_PORT"`
	Username       string `yaml:"SMTP_USERNAME" envconfig:"SMTP_USERNAME"`
	Password   string `yaml:"SMTP_PASSWORD" envconfig:"SMTP_PASSWORD"`
	SenderName string `yaml:"SMTP_SENDER_NAME" envconfig:"SMTP_SENDER_NAME"`
}
