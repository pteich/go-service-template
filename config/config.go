package config

import (
	"fmt"
)

// Config holds all config values, env variable names are defined as tags
type Config struct {
	LogLevel      string `env:"LOG_LEVEL"`
	ListenAddr    string `env:"LISTEN_ADDR"`
	GraylogServer string `env:"GRAYLOG_SERVER"`
	JwtSignKey    string `env:"JWT_SIGN_KEY"`
	LogJson       bool   `env:"LOG_JSON"`
	Version       string
}

// New returns a new config with defaults and parsed environement variables
func New() Config {

	config := Config{
		LogLevel:   "debug",
		ListenAddr: ":8000",
		JwtSignKey: "abc",
		LogJson:    true,
	}

	ParseEnv(&config)

	return config
}

// String returns a string version of the config struct
func (s *Config) String() string {
	return fmt.Sprintf("%#v", s)
}
