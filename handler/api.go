package handler

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"tui.com/baduk/greeter"
)

// ApiHandler is the place for all API routes
type ApiHandler struct {
	greeter greeter.Greeter
}

// NewApiHandler returns a new ApiHandler
func NewApiHandler(greeter greeter.Greeter) *ApiHandler {
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
