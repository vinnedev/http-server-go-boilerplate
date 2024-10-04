/**
* This is a test file for the health handler
**/

package http_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vinnedev/http-server-go-boilerplate/internal/app"
	httpInterface "github.com/vinnedev/http-server-go-boilerplate/internal/interfaces/http"
)

func TestHealthHandler_CheckHealth(t *testing.T) {
	// Initialize the health service and handler
	service := app.NewHealthService()
	handler := httpInterface.NewHealthHandler(service)

	// Create a simulated HTTP request
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Simulate the HTTP response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler.CheckHealth(rr, req)

	// Check the HTTP status
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Parse the response body into a map
	var responseBody map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Could not parse response body: %v", err)
	}

	// Expected JSON object
	expectedBody := map[string]string{"status": "healthy"}

	// Compare the actual response body with the expected body
	if responseBody["status"] != expectedBody["status"] {
		t.Errorf("Expected body %v, but got %v", expectedBody, responseBody)
	}
}
