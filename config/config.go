package config

import "os"

// Default port for service.
const serviceDefaultPort = "8080"

// Environment variable name for service port.
const servicePortEnvironmentVariable string = "SERVICE_PORT"

// AppConfig defines the Config for the App.
type AppConfig struct {
	Port string
}

// NewConfig inits the Config for the App and sets default values if the Envs are not set.
func NewConfig() AppConfig {
	appConfig := AppConfig{
		Port: serviceDefaultPort,
	}

	servicePort, set := os.LookupEnv(servicePortEnvironmentVariable)
	if set {
		appConfig.Port = servicePort
	}

	return appConfig
}
