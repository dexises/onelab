package config

import (
	"os"
)

type Config struct {
	HTTP                  HTTPConf
	DB                    DBConf
	JwtToken              string
	TransactionServiceURL string
}

type DBConf struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Sslmode  string
}

type HTTPConf struct {
	HttpPort string
}

func New() *Config {
	cfg := &Config{}

	cfg.HTTP.HttpPort = getEnv("HTTP_PORT", ":8080")

	cfg.JwtToken = getEnv("JwtTokenSecret", "abracadabra")

	cfg.TransactionServiceURL = getEnv("TransactionServiceURL", "http://localhost:9090")

	cfg.DB.Host = getEnv("DB_HOST", "postgres")
	cfg.DB.Port = getEnv("DB_PORT", "5432")
	cfg.DB.User = getEnv("DB_USER", "postgres")
	cfg.DB.Password = getEnv("DB_PASS", "postgres")
	cfg.DB.DBName = getEnv("DB_NAME", "postgres")
	cfg.DB.Sslmode = getEnv("DB_SSLMODE", "disable")

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
