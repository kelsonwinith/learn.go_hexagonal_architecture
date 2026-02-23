package config

import (
	log "log"

	godotenv "github.com/joho/godotenv"
	envconfig "github.com/kelseyhightower/envconfig"
)

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
