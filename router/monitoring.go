package router

import (
	"github.com/go-chi/render"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"tui.com/baduk/handler"
)

func AddMonitoringRoutes(router *chi.Mux, monitoringHandler *handler.MonitoringHandler) {

	router.Group(func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(middleware.DefaultCompress)
		r.Use(middleware.Timeout(60 * time.Second))
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Get("/health", monitoringHandler.GetHealth)
	})

}

