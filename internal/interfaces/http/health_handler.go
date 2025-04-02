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
	if err := json.NewEncoder(w).Encode(status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
