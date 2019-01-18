package config

import (
	"github.com/pkg/errors"
	"os"
	"strconv"
)

// Default port for service.
const serviceDefaultPort = "8080"

// default log level.
const serviceDefaultLogLevel = "debug"

// Environment variable name for service port.
const servicePortEnvironmentVariable string = "SERVICE_PORT"

// Environment variable name for log level.
const serviceLogLevelEnvironmentVariable string = "SERVICE_LOGLEVEL"

// Environment variable name for GELF log endpoint.
const serviceLogServerEnvironmentVariable string = "LOG_SERVER"

// AppConfig defines the Config for the App.
type AppConfig struct {
	Port string
	LogLevel string
	GelfLogServer string
}

// NewConfig inits the Config for the App and sets default values if the Envs are not set.
func NewConfig() (AppConfig, error) {
	appConfig := AppConfig{
		Port: serviceDefaultPort,
		LogLevel: serviceDefaultLogLevel,
	}

	servicePort, set := os.LookupEnv(servicePortEnvironmentVariable)
	if set {
		portNumeric, err := strconv.Atoi(servicePort)
		if err != nil {
			return appConfig, errors.Wrap(err, "could not initialize config")
		}

		if portNumeric < 1024 || portNumeric > 65536 {
			return appConfig, errors.New("port value not valid")
		}

		appConfig.Port = servicePort
	}

	logLevel, set := os.LookupEnv(serviceLogLevelEnvironmentVariable)
	if set {
		appConfig.LogLevel = logLevel
	}

	logServer, set := os.LookupEnv(serviceLogServerEnvironmentVariable)
	if set {
		appConfig.GelfLogServer = logServer
	}

	return appConfig, nil
}
