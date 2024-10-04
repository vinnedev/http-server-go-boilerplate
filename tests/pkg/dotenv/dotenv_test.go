/**
* This is a test file for dotenv package
**/

package dotenv_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/vinnedev/http-server-go-boilerplate/pkg/dotenv"
)

func TestReadEnvFromFile(t *testing.T) {
	// Create a temporary .env file for testing
	tempEnvFile := ".env.test"
	envContent := "TEST_KEY=TEST_VALUE"
	err := os.WriteFile(tempEnvFile, []byte(envContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary .env file: %v", err)
	}
	defer os.Remove(tempEnvFile)
	tests := []struct {
		key      string
		fallback string
		file     *string
		expected string
	}{
		{"TEST_KEY", "default", &tempEnvFile, "TEST_VALUE"},
		{"NON_EXISTENT_KEY", "default", &tempEnvFile, "default"},
	}
	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			value := dotenv.ReadEnvFromFile(tt.key, tt.fallback, tt.file)
			fmt.Println(tt.key, tt.fallback, tt.file)
			if value != tt.expected {
				t.Errorf("readEnvFromFile(%s, %s, %v) = %s; want %s", tt.key, tt.fallback, tt.file, value, tt.expected)
			}
		})
	}
}
func TestGetEnv(t *testing.T) {
	// Create a temporary .env file for testing
	_, currentPath, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(currentPath), "../../.env.test")
	envContent := "TEST_KEY=TEST_VALUE"
	err := os.WriteFile(basepath, []byte(envContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary .env file: %v", err)
	}
	defer os.Remove(basepath)

	tests := []struct {
		key      string
		fallback string
		expected string
	}{
		{"TEST_KEY", "default", "TEST_VALUE"},
		{"NON_EXISTENT_KEY", "default", "default"},
	}

	for _, tt := range tests {
		t.Run(tt.key, func(t *testing.T) {
			value := dotenv.GetEnv(tt.key, tt.fallback)
			if tt.key != "" && value != tt.expected {
				t.Errorf("GetEnv(%s, %s) = %s; want %s", tt.key, tt.fallback, value, tt.expected)
			}
		})
	}
}
