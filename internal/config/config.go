package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPAddr string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	config := Config{
		HTTPAddr: os.Getenv("HTTP_ADDR"),
	}

	if config.HTTPAddr == "" {
		return nil, errors.New("missing required http configuration")
	}

	return &config, nil
}
