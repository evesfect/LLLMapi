// internal/config/config.go
package config

import (
	"time"
)

type Config struct {
	API     APIConfig
	Queue   QueueConfig
	Browser BrowserConfig
}

type APIConfig struct {
	Host string
	Port int
}

type QueueConfig struct {
	WorkerCount  int
	MaxQueueSize int
	BatchSize    int
	BatchTimeout time.Duration
}

type BrowserConfig struct {
	MaxInstances int
	Headless     bool
	UserAgent    string
	Timeout      int
}

func Load() (*Config, error) {
	// Return default config for now
	return &Config{
		API: APIConfig{
			Host: "localhost",
			Port: 8080,
		},
		Queue: QueueConfig{
			WorkerCount:  5,
			MaxQueueSize: 1000,
			BatchSize:    10,
			BatchTimeout: 5 * time.Second,
		},
		Browser: BrowserConfig{
			MaxInstances: 3,
			Headless:     true,
			UserAgent:    "Mozilla/5.0",
			Timeout:      30,
		},
	}, nil
}
