package server

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	_defaultPort            = 8090
	_defaultReadTimeout     = 15 * time.Second
	_defaultWriteTimeout    = 15 * time.Second
	_defaultIdleTimeout     = 60 * time.Second
	_defaultShutdownTimeout = 10 * time.Second
)

// Config holds HTTP server configuration. The zero value is usable: any unset
// field falls back to a sensible default.
type Config struct {
	Host            string
	Port            int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

// ConfigFromEnv builds a Config from environment variables. Unset variables are
// left as zero values, which New resolves to defaults. Recognized variables:
//
//	HOST                   listen host (default: all interfaces)
//	PORT                   listen port (default: 8090)
//	HTTP_READ_TIMEOUT      e.g. "15s"
//	HTTP_WRITE_TIMEOUT     e.g. "15s"
//	HTTP_IDLE_TIMEOUT      e.g. "60s"
//	HTTP_SHUTDOWN_TIMEOUT  e.g. "10s"
func ConfigFromEnv() (Config, error) {
	var c Config
	c.Host = os.Getenv("HOST")

	if v, ok := os.LookupEnv("PORT"); ok {
		port, err := strconv.Atoi(v)
		if err != nil {
			return Config{}, fmt.Errorf("parse PORT: %w", err)
		}
		c.Port = port
	}

	durations := []struct {
		key string
		dst *time.Duration
	}{
		{"HTTP_READ_TIMEOUT", &c.ReadTimeout},
		{"HTTP_WRITE_TIMEOUT", &c.WriteTimeout},
		{"HTTP_IDLE_TIMEOUT", &c.IdleTimeout},
		{"HTTP_SHUTDOWN_TIMEOUT", &c.ShutdownTimeout},
	}
	for _, d := range durations {
		v, ok := os.LookupEnv(d.key)
		if !ok {
			continue
		}
		dur, err := time.ParseDuration(v)
		if err != nil {
			return Config{}, fmt.Errorf("parse %s: %w", d.key, err)
		}
		*d.dst = dur
	}

	return c, nil
}

func (c Config) withDefaults() Config {
	if c.Port == 0 {
		c.Port = _defaultPort
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = _defaultReadTimeout
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = _defaultWriteTimeout
	}
	if c.IdleTimeout == 0 {
		c.IdleTimeout = _defaultIdleTimeout
	}
	if c.ShutdownTimeout == 0 {
		c.ShutdownTimeout = _defaultShutdownTimeout
	}
	return c
}
