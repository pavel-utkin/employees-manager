package config

import "github.com/sirupsen/logrus"

type Config struct {
	AddressREST string `env:"AddressREST"`
	DSN         string `env:"DATABASE_DSN"`
}

func NewConfig(log *logrus.Logger) *Config {
	configServer := Config{
		AddressREST: "localhost:8088",
		DSN:         "root:password@tcp(127.0.0.1:3306)/Employees",
	}
	return &configServer
}
