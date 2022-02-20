package app

import "time"

var (
	// Conf global app var
	Conf *Config
)

// Config global config
// nolint
type Config struct {
	HTTP ServerConfig
}

// ServerConfig server config.
type ServerConfig struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
