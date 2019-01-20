package handler

import (
	"github.com/go-chi/render"
	"net/http"
)

// MonitoringHandler is the place for handlers that care about monitoring
type MonitoringHandler struct {
	Version string
}

// NewMonitoringHandler returns a new MonitoringHandler
func NewMonitoringHandler(version string) *MonitoringHandler {
	return &MonitoringHandler{
		Version: version,
	}
}

func (handler *MonitoringHandler) GetHealth(w http.ResponseWriter, r *http.Request) {

	// define ad-hoc struct for health response
	response := struct {
		Status string
		Version string
	}{
		Status: "ok",
		Version: handler.Version,
	}

	render.JSON(w, r, response)
}
