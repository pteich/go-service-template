package router

import (
	"github.com/sirupsen/logrus"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"tui.com/baduk/handler"
)

func AddApiRoutes(router *chi.Mux, apiHandler *handler.ApiHandler, logger *logrus.Entry) {

	router.Group(func(r chi.Router) {
		r.Use(middleware.Recoverer)
		r.Use(middleware.DefaultCompress)
		r.Use(middleware.Timeout(60 * time.Second))
		r.Use(middleware.StripSlashes)
		r.Use(middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: logger}))

		r.Get("/greet/{name}", apiHandler.GetName)
	})

}
