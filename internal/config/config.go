package config

import (
	"os"
)

type Config struct {
	Port string
}

func New() *Config {
	port := getEnv("PORT", ":8080")
	cfg := &Config{
		Port: port,
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
