package router

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/pteich/go-service-template/handler"
	"github.com/pteich/go-service-template/logger"
)

// AddMonitoringRoutes sets routes for health check and monitoring
func AddMonitoringRoutes(router *chi.Mux, monitoringHandler *handler.MonitoringHandler, logger *logger.Logger) {

	router.Group(func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(middleware.DefaultCompress)
		r.Use(middleware.Timeout(60 * time.Second))
		r.Use(render.SetContentType(render.ContentTypeJSON))
		r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))

		r.Get("/health", monitoringHandler.GetHealth)
	})

}
