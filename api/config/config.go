package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port    string
	GinMode string
}

// LoadConfig loads configuration from .env file and maps it to a Config struct.
func LoadConfig() Config {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	return Config{
		Port:    getEnv("PORT", "8080"),
		GinMode: getEnv("GIN_MODE", "debug"),
	}
}

// getEnv retrieves the environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
