package config

import "github.com/caarlos0/env/v8"

type Config struct {
	App App
	DB  DB
	Log Log
}

type App struct {
	Port             string `env:"APP_PORT" envDefault:"8080"`
	ReportsDirectory string `env:"REPORT_DIRECTORY" envDefault:"./reports"`
}

type DB struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	User     string `env:"DB_USER"`
	Pwd      string `env:"DB_PWD"`
	Name     string `env:"DB_NAME" envDefault:"users-segments"`
	SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
	DNS      string `env:"DB_DNS"`
	TimeZone string `env:"DB_TIMEZONE" envDefault:"Europe/Moscow"`
}

type Log struct {
	Level string `env:"LOG_LEVEL" envDefault:"debug"`
}

func NewConfig() (*Config, error) {
	config := &Config{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}

	return config, nil
}
