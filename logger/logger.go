package logger

import (
	"github.com/halink0803/zerolog-graylog-hook/graylog"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"os"
	"tui.com/baduk/config"
)

// NewLogger creates a new logger instance based on a given config
func NewLogger(appConfig config.AppConfig) (zerolog.Logger, error) {

	// parse loglevel from string and set it globally
	zerologlevel, err := zerolog.ParseLevel(appConfig.LogLevel)
	if err != nil {
		zerologlevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(zerologlevel)

	// set up console writer
	output := zerolog.ConsoleWriter{
		Out: os.Stderr,
		NoColor: false,
	}

	// create new logger with timestamp
	logger := zerolog.New(output).With().Timestamp().Str("service", config.ServiceName).Logger()

	// check if GELF server is configured and attach hook
	if appConfig.GelfLogServer != "" {

		hook, err := graylog.NewGraylogHook(appConfig.GelfLogServer)
		if err != nil {
			return logger, errors.Wrap(err, "could not add GELF hook")
		}

		logger.Hook(hook)
	}

	return logger, nil
}
