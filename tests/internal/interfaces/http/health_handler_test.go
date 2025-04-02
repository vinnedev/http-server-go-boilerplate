/**
* This is a test file for the health handler
**/

package http_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vinnedev/http-server-go-boilerplate/internal/app"
	httpInterface "github.com/vinnedev/http-server-go-boilerplate/internal/interfaces/http"
)

func TestHealthHandler_CheckHealth(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:           "successful health check with GET",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"status": "healthy"},
		},
		{
			name:           "successful health check with POST",
			method:         "POST",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"status": "healthy"},
		},
		{
			name:           "successful health check with PUT",
			method:         "PUT",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"status": "healthy"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize the health service and handler
			service := app.NewHealthService()
			handler := httpInterface.NewHealthHandler(service)

			// Create a simulated HTTP request
			req, err := http.NewRequest(tt.method, "/health", nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			// Simulate the HTTP response recorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler.CheckHealth(rr, req)

			// Check the HTTP status
			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status code %d, but got %d", tt.expectedStatus, rr.Code)
			}

			// If we expect a body, parse and check it
			if tt.expectedBody != nil {
				var responseBody map[string]string
				err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
				if err != nil {
					t.Fatalf("Could not parse response body: %v", err)
				}

				// Compare the actual response body with the expected body
				if responseBody["status"] != tt.expectedBody["status"] {
					t.Errorf("Expected body %v, but got %v", tt.expectedBody, responseBody)
				}

				// Check Content-Type header
				contentType := rr.Header().Get("Content-Type")
				if contentType != "application/json" {
					t.Errorf("Expected Content-Type to be application/json, got %s", contentType)
				}
			}
		})
	}
}

func TestHealthHandler_CheckHealth_ErrorHandling(t *testing.T) {
	// Initialize the health service and handler
	service := app.NewHealthService()
	handler := httpInterface.NewHealthHandler(service)

	// Create a request with a nil body to simulate an error
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Create a custom response writer that fails on Write
	rr := &errorResponseWriter{
		ResponseRecorder: *httptest.NewRecorder(),
	}

	// Call the handler
	handler.CheckHealth(rr, req)

	// Verify that the error was handled properly
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, rr.Code)
	}
}

// errorResponseWriter is a custom ResponseWriter that fails on Write
type errorResponseWriter struct {
	httptest.ResponseRecorder
}

func (w *errorResponseWriter) Write(b []byte) (int, error) {
	return 0, errors.New("simulated write error")
}

func (w *errorResponseWriter) WriteHeader(statusCode int) {
	w.Code = statusCode
}
