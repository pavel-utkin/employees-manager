package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AddressREST string `env:"AddressREST"`
	DSN         string `env:"DATABASE_DSN"`
}

func NewConfig(log *logrus.Logger) *Config {
	configServer := Config{
		AddressREST: "localhost:8088",
		DSN:         "root:password@tcp(127.0.0.1:3306)/Employees",
	}
	flag.StringVar(&configServer.AddressREST, "a", configServer.AddressREST, "Server address")
	flag.StringVar(&configServer.DSN, "d", configServer.DSN, "DATABASE DSN")
	flag.Parse()
	err := env.Parse(&configServer)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(configServer)
	return &configServer
}
