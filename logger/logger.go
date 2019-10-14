package logger

import (
	"github.com/gemnasium/logrus-graylog-hook/v3"
	"github.com/pteich/go-service-template/config"
	"github.com/sirupsen/logrus"
)

// NewLogger creates a new logger instance based on a given config
func NewLogger(appConfig config.AppConfig) *logrus.Entry {

	// parse loglevel from config string and set it globally
	loglevel, err := logrus.ParseLevel(appConfig.LogLevel)
	if err != nil {
		loglevel = logrus.DebugLevel
	}
	logrus.SetLevel(loglevel)

	// enable timestamp in console output
	formatter := &logrus.TextFormatter{
		FullTimestamp: true,
	}
	logrus.SetFormatter(formatter)

	logger := logrus.WithField("service", config.ServiceName)

	// check if GELF server is configured and attach hook
	if appConfig.GelfLogServer != "" {

		hook := graylog.NewGraylogHook(appConfig.GelfLogServer, map[string]interface{}{})
		logrus.AddHook(hook)
	}

	return logger
}
