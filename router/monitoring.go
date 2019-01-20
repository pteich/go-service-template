package router

import (
	"time"

	"github.com/go-chi/render"
	"github.com/rs/zerolog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"tui.com/baduk/handler"
)

func AddMonitoringRoutes(router *chi.Mux, monitoringHandler *handler.MonitoringHandler, logger zerolog.Logger) {

	router.Group(func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(middleware.DefaultCompress)
		r.Use(middleware.Timeout(60 * time.Second))
		r.Use(render.SetContentType(render.ContentTypeJSON))
		r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: &logger}))

		r.Get("/health", monitoringHandler.GetHealth)
	})

}
