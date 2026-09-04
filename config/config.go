package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	DBDSN    string
	MongoURI string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:     getEnv("PORT", "8080"),
		DBDSN:    getEnv("DB_DSN", ""),
		MongoURI: getEnv("MONGO_URI", ""),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
