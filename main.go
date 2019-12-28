package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/pteich/go-service-template/config"
	"github.com/pteich/go-service-template/greeter"
	"github.com/pteich/go-service-template/handler"
	"github.com/pteich/go-service-template/logger"
	"github.com/pteich/go-service-template/router"
)

var Version string

func main() {

	var err error

	// initialize config
	cfg, err := config.New()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	// init logger with config
	log := logger.New(logger.WithLogLevel(cfg.LogLevel), logger.WithConsoleOutput(cfg.LogOutputConsole))

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

	// handle common signals, could be extended to allow graceful restarts and finish background jobs
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		// in case of any signal just cancel the main context
		// could be a different reaction for some signals e.g. reload on sighup
		switch <-signalChannel {
		case os.Interrupt:
			done()
		case syscall.SIGTERM:
			done()
		case syscall.SIGINT:
			done()
		case syscall.SIGHUP:
			done()
		}
	}()

	// Start server
	srv := &http.Server{
		Handler: chiRouter,
		Addr:    cfg.ListenAddr,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("server not started")
		}
	}()

	log.Info().Str("config", cfg.String()).Msg("service started")

	// wait until done
	<-ctx.Done()

	// shutdown
	err = srv.Close()
	if err != nil {
		log.Error().Err(err).Msg("error closing server")
	}

	log.Info().Msg("service stopped")
}
