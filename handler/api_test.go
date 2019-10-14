package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/pteich/go-service-template/greeter"
	"github.com/stretchr/testify/assert"
)

func TestApiHandler_GetName(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// set expectations for service calls
	greetingService := greeter.NewMockGreeter(mockCtrl)
	greetingService.EXPECT().SayHello("World").Return("Hello World!").Times(1)

	handler := NewApiHandler(greetingService)

	mockRouter := chi.NewRouter()
	mockRouter.Get("/greet/{name}", handler.GetName)

	req, err := http.NewRequest("GET", "/greet/World", nil)
	assert.NoError(t, err)

	respRecorder := httptest.NewRecorder()
	mockRouter.ServeHTTP(respRecorder, req)

	assert.Equal(t, 200, respRecorder.Code)
	assert.Equal(t, "text/plain; charset=utf-8", respRecorder.Header().Get("content-type"))
	assert.Equal(t, "Hello World!", respRecorder.Body.String())

}
