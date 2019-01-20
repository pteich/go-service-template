package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"tui.com/baduk/config"
	"tui.com/baduk/greeter"
	"tui.com/baduk/handler"
	"tui.com/baduk/logger"
	"tui.com/baduk/router"
)

var Version string

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

	// Init services
	greetingService := greeter.NewGreet()

	// Init handlers
	monitoringHandler := handler.NewMonitoringHandler(Version)
	apiHandler := handler.NewApiHandler(greetingService)

	// Init Router
	chiRouter := chi.NewRouter()
	router.AddMonitoringRoutes(chiRouter, monitoringHandler, log)
	router.AddApiRoutes(chiRouter, apiHandler, log)

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

	// Start server
	srv := &http.Server{
		Handler: chiRouter,
		Addr:    fmt.Sprintf(":%s", config.Port),
	}

	go srv.ListenAndServe()

	log.Info().Str("port", config.Port).Msg("service started")

	// wait until done
	<-ctx.Done()
	log.Info().Msg("service stopped")
}
