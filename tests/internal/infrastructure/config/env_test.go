/**
* This is a test file for config package
**/

package config_test

import (
	"os"
	"testing"

	"github.com/vinnedev/http-server-go-boilerplate/internal/infrastructure/config"
)

func TestEnvVariables(t *testing.T) {
	// Test default values
	if config.ENV_MODE != "development" {
		t.Errorf("Expected ENV_MODE to be 'development', got %s", config.ENV_MODE)
	}
	if config.PORT != "8080" {
		t.Errorf("Expected PORT to be '8080', got %s", config.PORT)
	}

	// Test with custom environment variables
	os.Setenv("ENV_MODE", "production")
	os.Setenv("PORT", "3000")
	defer os.Unsetenv("ENV_MODE")
	defer os.Unsetenv("PORT")

	// Note: Since the variables are initialized at package level,
	// we can't test the custom values in the same test.
	// This is a limitation of testing package-level variables.
}
