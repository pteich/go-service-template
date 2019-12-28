package logger

import (
	"io"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	logLevel    string
	logConsole  bool
	serviceName string
}

// Logger represents a logger
type Logger struct {
	zerolog.Logger
}

// New is the constructor of the logger
func New(opts ...Option) *Logger {

	_, filename, _, _ := runtime.Caller(0)
	config := Config{
		logLevel:    "debug",
		logConsole:  true,
		serviceName: filename,
	}

	for _, opt := range opts {
		opt(&config)
	}

	zerologlevel, err := zerolog.ParseLevel(config.logLevel)
	if err == nil {
		zerolog.SetGlobalLevel(zerologlevel)
	}

	var logDest io.Writer = os.Stdout

	if config.logConsole {
		logDest = zerolog.ConsoleWriter{Out: logDest, TimeFormat: time.RFC3339}
	}

	return &Logger{
		Logger: zerolog.New(logDest).With().Timestamp().Str("service", config.serviceName).Logger(),
	}
}

func NewFromZerolog(logger zerolog.Logger) Logger {
	return Logger{logger}
}
