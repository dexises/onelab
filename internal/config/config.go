package config

import (
	"os"
)

type Config struct {
	Addr string
}

func Load() *Config {
	addr := getEnv("ADDR", ":8080")
	return &Config{
		Addr: addr,
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
