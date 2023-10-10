package config

import "github.com/sirupsen/logrus"

type Config struct {
	AddressREST string `env:"AddressREST"`
	DSN         string `env:"DATABASE_DSN"`
}

func NewConfig(log *logrus.Logger) *Config {
	configServer := Config{
		AddressREST: "localhost:8088",
		DSN:         "host=localhost port=5432 user=user password=password dbname=gophkeeper sslmode=disable",
	}
	return &configServer
}
