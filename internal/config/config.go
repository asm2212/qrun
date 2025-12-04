package config

import (
	"log"
	"os"
	"strconv"
)

// config is the main global configuration structure
// every service(api,worker) will load this same config
type Config struct {
	AppName     string
	Environment string
	Version     string

	Server struct {
		Port              string
		ShutdownTimeoutMs int
	}
	Redis struct {
		Host          string
		Port          string
		Password      string
		DB            int
		Stream        string
		Consumer      string
		ConsumerGroup string
	}
	Worker struct {
		Concurrency int
		MaxRetries  int
	}
}

// loadEnv returns an environment variable or a fallback
func loadEnv(key string, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue

}

// loadEnvInt returns an integer environment variable or fallback default
func loadEnvInt(key string, defaultValue int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue
	}
	number, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return number

}

// New loads and returns the full configuration.
func New() *Config {
	cfg := &Config{}

	// Application metadata
	cfg.AppName = loadEnv("APP_NAME", "qrun")
	cfg.Version = loadEnv("APP_VERSION", "1.0.0")
	cfg.Environment = loadEnv("APP_ENV", "development")

	// Server Configuration
	cfg.Server.Port = loadEnv("SERVER_PORT", "8080")
	cfg.Server.ShutdownTimeoutMs = loadEnvInt("SHUTDOWN_TIMEOUT_MS", 5000)

	// Redis / Queue Configuration
	cfg.Redis.Host = loadEnv("REDIS_HOST", "localhost")
	cfg.Redis.Port = loadEnv("REDIS_PORT", "6379")
	cfg.Redis.Password = loadEnv("REDIS_PASSWORD", "")
	cfg.Redis.DB = loadEnvInt("REDIS_DB", 0)
	cfg.Redis.Stream = loadEnv("REDIS_STREAM", "qrun_stream")
	cfg.Redis.Consumer = loadEnv("REDIS_CONSUMER_NAME", "qrun_consumer")
	cfg.Redis.ConsumerGroup = loadEnv("REDIS_CONSUMER_GROUP", "qrun_group")

	// Worker Configuration
	cfg.Worker.Concurrency = loadEnvInt("WORKER_CONCURRENCY", 4)
	cfg.Worker.MaxRetries = loadEnvInt("WORKER_MAX_RETRIES", 3)

	log.Printf("Config loaded — Environment: %s, Port: %s\n", cfg.Environment, cfg.Server.Port)
	return cfg
}
