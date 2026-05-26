package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	response := httptest.NewRecorder()

	HealthHandler(response, request)

	if response.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	body := response.Body.String()

	if !strings.Contains(body, `"status":"ok"`) {
		t.Errorf("expected body to contain status ok, got %s", body)
	}
}
