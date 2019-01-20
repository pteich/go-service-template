package handler

import (
	"context"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"tui.com/baduk/greeter"
)

func TestApiHandler_GetName(t *testing.T) {

	// init mock controller
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// init mock and set expectations
	greetingService := greeter.NewMockGreeter(mockCtrl)
	greetingService.EXPECT().SayHello("World").Return("Hello World!").Times(1)

	// init handler with mock
	handler := NewApiHandler(greetingService)

	// create http request with chi context
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/greet/World", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("name", "World")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))

	// execute handler
	handler.GetName(w, r)

	// check results
	resp := w.Result()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("content-type"))
	assert.Equal(t, "Hello World!", string(body))

}
