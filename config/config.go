package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port       string `envconfig:"PORT" default:"8080"`
	DBDriver   string `envconfig:"DB_DRIVER" default:"postgres"`
	DBSource   string `envconfig:"DB_SOURCE" default:"postgresql://postgres:password@localhost:5432/stuHub?sslmode=disable"`
	DBName     string `envconfig:"DB_NAME" default:"stuHub"`
	DBUser     string `envconfig:"DB_USER" default:"postgres"`
	DBPassword string `envconfig:"DB_PASSWORD" default:"password"`
}

func loadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
