package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test default values
	os.Unsetenv("PORT")
	os.Unsetenv("REDIS_PORT")
	cfg := LoadConfig()
	if cfg.Port != "8080" {
		t.Errorf("Expected default PORT 8080, got %s", cfg.Port)
	}
	if cfg.RedisPort != "6379" {
		t.Errorf("Expected default REDIS_PORT 6379, got %s", cfg.RedisPort)
	}

	// Test environment variable override
	os.Setenv("PORT", "9090")
	os.Setenv("REDIS_PORT", "6380")
	cfg = LoadConfig()
	if cfg.Port != "9090" {
		t.Errorf("Expected PORT 9090, got %s", cfg.Port)
	}
	if cfg.RedisPort != "6380" {
		t.Errorf("Expected REDIS_PORT 6380, got %s", cfg.RedisPort)
	}
}
