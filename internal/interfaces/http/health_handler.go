/**
* This is a controller for health check
**/

package http

import (
	"encoding/json"
	"net/http"

	"github.com/vinnedev/http-server-go-boilerplate/internal/app"
)

type HealthHandler struct {
	service *app.HealthService
}

func NewHealthHandler(service *app.HealthService) *HealthHandler {
	return &HealthHandler{service: service}
}

func (h *HealthHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	status := h.service.CheckHealth()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(status)
}
