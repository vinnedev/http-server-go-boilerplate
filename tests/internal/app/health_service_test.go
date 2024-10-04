/**
* This is a test file for health service
**/

package app_test

import (
	"testing"

	"github.com/vinnedev/http-server-go-boilerplate/internal/app"
)

func TestHealthService_CheckHealth(t *testing.T) {
	// Initialize the service
	service := app.NewHealthService()

	// Execute the health check function
	status := service.CheckHealth()

	// Verify if the returned status is as expected
	expectedStatus := "healthy"
	if status["status"] != expectedStatus {
		t.Errorf("Expected status %s, but got %s", expectedStatus, status["status"])
	}
}
