package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"tui.com/baduk/config"
	"tui.com/baduk/logger"
)

func main() {

	// initialize config
	config, err := config.NewConfig()
	if err != nil {
		panic("could not read config")
	}

	// init logger with config
	log, err := logger.NewLogger(config)
	if err != nil {
		panic("could not initialize logger")
	}

	log.Info().Msg("service started")

	// create main context
	ctx, done := context.WithCancel(context.Background())

	// handle common signals, could be extended to allow graceful restarts
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		switch <-signalChannel {
		case os.Interrupt:
			done()
		case syscall.SIGTERM:
			done()
		}
	}()


	// wait until done
	<-ctx.Done()
	log.Info().Msg("service stopped")
}
