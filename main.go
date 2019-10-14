package main

import (
	"context"
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

	// initialize config
	appConfig := config.New()

	// init logger with config
	log := logger.NewLogger(appConfig)

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
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		switch <-signalChannel {
		case os.Interrupt:
			done()
		case syscall.SIGTERM:
			done()
		case syscall.SIGINT:
			done()
		case syscall.SIGHUP:
			done()
		case syscall.SIGQUIT:
			done()
		}
	}()

	// Start server
	srv := &http.Server{
		Handler: chiRouter,
		Addr:    appConfig.ListenAddr,
	}

	go srv.ListenAndServe()

	log.WithField("config", appConfig.String()).Info("service started")

	// wait until done
	<-ctx.Done()
	log.Info("service stopped")
}
