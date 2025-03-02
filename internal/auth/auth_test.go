package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Error no header", func(t *testing.T) {
		testHeader := http.Header{}
		testHeader.Add("Authorization", "")

		_, err := GetAPIKey(testHeader)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("Error not api key", func(t *testing.T) {
		testHeader := http.Header{}
		testHeader.Add("Authorization", "Bearer 123")

		_, err := GetAPIKey(testHeader)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("error auth too short", func(t *testing.T) {
		testHeader := http.Header{}
		testHeader.Add("Authorization", "ApiKey")

		_, err := GetAPIKey(testHeader)

		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("happy path, gets error", func(t *testing.T) {
		testHeader := http.Header{}
		testHeader.Add("Authorization", "ApiKey 12345")

		result, err := GetAPIKey(testHeader)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if result != "12345" {
			t.Errorf("unexpected result: expected 12345, got %s", result)
		}
	})
}
