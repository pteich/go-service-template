package router

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/pteich/go-service-template/handler"
	"github.com/pteich/go-service-template/logger"
)

// AddApiRoutes sets routes for API access
func AddApiRoutes(router *chi.Mux, apiHandler *handler.ApiHandler, logger *logger.Logger) {

	router.Group(func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(middleware.DefaultCompress)
		r.Use(middleware.Timeout(60 * time.Second))
		r.Use(middleware.StripSlashes)
		r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))

		r.Get("/greet/{name}", apiHandler.GetName)
	})

}
