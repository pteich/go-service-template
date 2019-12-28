package logger

type Option func(c *Config)

func WithLogLevel(logLevel string) Option {
	return func(c *Config) {
		c.logLevel = logLevel
	}
}

func WithConsoleOutput(logConsole bool) Option {
	return func(c *Config) {
		c.logConsole = logConsole
	}
}

func WithServiceName(serviceName string) Option {
	return func(c *Config) {
		c.serviceName = serviceName
	}
}
