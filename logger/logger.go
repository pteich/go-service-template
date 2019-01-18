package logger

import (
	"os"

	"github.com/rs/zerolog"

	"tui.com/baduk/config"
)

func NewLogger(appConfig config.AppConfig) zerolog.Logger {

	// parse loglevel from string and set it globally
	zerologlevel, err := zerolog.ParseLevel(appConfig.LogLevel)
	if err != nil {
		zerologlevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(zerologlevel)

	// create new logger with timestamp
	return zerolog.New(os.Stdout).With().Timestamp().Str("service", config.ServiceName).Logger()
}
