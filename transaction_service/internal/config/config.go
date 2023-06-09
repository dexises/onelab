package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	HTTP     HTTPConf
	DB       DBConf
	JwtToken string
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

	cfg.HTTP.HttpPort = getEnv("HTTP_PORT", ":9090")

	cfg.JwtToken = getEnv("JwtTokenSecret", "abracadabra")

	cfg.DB.Host = getEnv("DB_HOST", "postgre")
	cfg.DB.Port = getEnv("DB_PORT", "5432")
	cfg.DB.User = getEnv("DB_USER", "postgres")
	cfg.DB.Password = getEnv("DB_PASS", "postgres")
	cfg.DB.DBName = getEnv("DB_NAME", "postgre")
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
