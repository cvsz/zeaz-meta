package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	ServiceName string
	HTTPAddr    string
	NATSURL     string
	DatabaseURL string
	RedisURL    string
	VaultAddr   string
	TenantMode  string
	ReadTimeout time.Duration
}

func Load(serviceName string) Config {
	return Config{
		ServiceName: value("SERVICE_NAME", serviceName),
		HTTPAddr:    value("HTTP_ADDR", ":8080"),
		NATSURL:     value("NATS_URL", "nats://nats:4222"),
		DatabaseURL: value("DATABASE_URL", "postgres://zeaz:zeaz@postgres:5432/zeaz?sslmode=require"),
		RedisURL:    value("REDIS_URL", "redis://redis:6379/0"),
		VaultAddr:   value("VAULT_ADDR", "https://vault.vault.svc:8200"),
		TenantMode:  value("TENANT_MODE", "strict"),
		ReadTimeout: time.Duration(intValue("READ_TIMEOUT_SECONDS", 5)) * time.Second,
	}
}

func value(key, fallback string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return fallback
}

func intValue(key string, fallback int) int {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		if parsed, err := strconv.Atoi(v); err == nil {
			return parsed
		}
	}
	return fallback
}
