/**
* This is a test file for domain package
**/

package domain_test

import (
	"encoding/json"
	"testing"

	"github.com/vinnedev/http-server-go-boilerplate/internal/domain"
)

func TestHealthStatus(t *testing.T) {
	// Test struct creation
	status := domain.HealthStatus{
		Status: "healthy",
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(status)
	if err != nil {
		t.Fatalf("Failed to marshal HealthStatus: %v", err)
	}

	expectedJSON := `{"status":"healthy"}`
	if string(jsonData) != expectedJSON {
		t.Errorf("Expected JSON %s, got %s", expectedJSON, string(jsonData))
	}

	// Test JSON unmarshaling
	var unmarshaledStatus domain.HealthStatus
	err = json.Unmarshal(jsonData, &unmarshaledStatus)
	if err != nil {
		t.Fatalf("Failed to unmarshal HealthStatus: %v", err)
	}

	if unmarshaledStatus.Status != status.Status {
		t.Errorf("Expected status %s, got %s", status.Status, unmarshaledStatus.Status)
	}
}
