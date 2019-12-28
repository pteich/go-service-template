package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Greeter interface {
	SayHello(name string) string
}

// ApiHandler is the place for all API routes
type ApiHandler struct {
	greeter Greeter
}

// NewApiHandler returns a new ApiHandler
func NewApiHandler(greeter Greeter) *ApiHandler {
	return &ApiHandler{
		greeter: greeter,
	}
}

func (handler *ApiHandler) GetName(w http.ResponseWriter, r *http.Request) {

	name := chi.URLParam(r, "name")
	greeting := handler.greeter.SayHello(name)

	render.SetContentType(render.ContentTypePlainText)
	render.PlainText(w, r, greeting)
}
