package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	App AppConfig
	Http HttpConfig
	Postgres PostgresConfig
}

type AppConfig struct {
	Debug bool `env:"DEBUG"`
}

type HttpConfig struct {
	Host string `env:"HTTP_HOST"`
	Port uint16 `env:"HTTP_PORT"`
}

type PostgresConfig struct {
	Host string `env:"POSTGRES_HOST"`
	Port uint16 `env:"POSTGRES_PORT"`
	User string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	DbName string `env:POSTGRES_DBNAME"`
	DbSchema string `env:POSTGRES_SCHEMA"`
}

func Load() (*Config, error) {
	return readEnv()
}

func readEnv() (*Config, error) {
	var app AppConfig
	if err := env.Parse(&app); err != nil {
		return nil, err
	}

	var http HttpConfig 
	if err := env.Parse(&http); err != nil {
		return nil, err
	}

	var postgres PostgresConfig
	if err := env.Parse(&postgres); err != nil {
		return nil, err
	}

	return &Config{
		App: app,
		Http: http,
		Postgres: postgres,
	}, nil
}