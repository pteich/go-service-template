package router

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"tui.com/baduk/handler"
)

func AddApiRoutes(router *chi.Mux, apiHandler *handler.ApiHandler, logger zerolog.Logger) {

	router.Group(func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(middleware.DefaultCompress)
		r.Use(middleware.Timeout(60 * time.Second))
		r.Use(middleware.StripSlashes)
		r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: &logger}))

		r.Get("/greet/{name}", apiHandler.GetName)
	})

}
