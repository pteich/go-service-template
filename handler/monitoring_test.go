package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMonitoringHandler_GetHealth(t *testing.T) {

	// init handler with version
	handler := NewMonitoringHandler("1.0.0")

	testServer := httptest.NewServer(http.HandlerFunc(handler.GetHealth))

	resp, err := testServer.Client().Get(testServer.URL)
	assert.NoError(t, err)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("content-type"))
	assert.Equal(t, 200, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), "ok")

}
