package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/dto"
)

func TestAnalyzePromiseHandler(t *testing.T) {
	requestBody := map[string]string{
		"text": "reduzir imposto sobre combustível",
	}

	jsonBody, _ := json.Marshal(requestBody)

	request := httptest.NewRequest(
		http.MethodPost,
		"/promises/analyze",
		bytes.NewBuffer(jsonBody),
	)

	request.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	AnalyzePromiseHandler(recorder, request)

	response := recorder.Result()

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", response.StatusCode)
	}

	var responseBody dto.AnalyzePromiseResponse

	err := json.NewDecoder(response.Body).Decode(&responseBody)

	if err != nil {
		t.Errorf("failed to decode response")
	}

	if responseBody.Score != 60 {
		t.Errorf("expected score 60, got %d", responseBody.Score)
	}
}
