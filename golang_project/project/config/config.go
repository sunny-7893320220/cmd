package config

import (
	"os"
)

// Config holds the application configuration
type Config struct {
	Port      string
	RedisPort string
}

// LoadConfig creates a new Config instance and populates it from environment variables
func LoadConfig() *Config {
	return &Config{
		Port:      getEnv("PORT", "8080"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
	}
}

// getEnv is a helper function to read an environment variable or return a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
