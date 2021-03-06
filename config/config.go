package config

import (
	"fmt"
)

// Config holds all config values, env variable names are defined as tags
type Config struct {
	LogLevel         string `env:"LOG_LEVEL"`
	ListenAddr       string `env:"LISTEN_ADDR"`
	LogOutputConsole bool   `env:"LOG_OUTPUT_CONSOLE"`
	Version          string
}

// New returns a new config with defaults and parsed environment variables
func New() (Config, error) {

	config := Config{
		LogLevel:         "debug",
		ListenAddr:       ":8000",
		LogOutputConsole: true,
	}

	err := ParseEnv(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// String returns a string version of the config struct
func (s Config) String() string {
	return fmt.Sprintf("%#v", s)
}
