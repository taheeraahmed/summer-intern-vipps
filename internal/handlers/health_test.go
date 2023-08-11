package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	// Create a new HTTP request to the health endpoint.
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to capture the response.
	rr := httptest.NewRecorder()

	// Call the handler function directly.
	HealthHandler(rr, req)

	// Check the response status code.
	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", rr.Code, http.StatusOK)
	}
}
