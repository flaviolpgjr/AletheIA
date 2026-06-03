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

	if responseBody.Score != 35 {
		t.Errorf("expected score 35, got %d", responseBody.Score)
	}

	if responseBody.Confidence != 30 {
		t.Errorf("expected confidence 30, got %d", responseBody.Confidence)
	}

	if len(responseBody.Criteria) != 6 {
		t.Errorf("expected 6 criteria, got %d", len(responseBody.Criteria))
	}

	if responseBody.Criteria[0].Key == "" {
		t.Errorf("expected criterion key to be present")
	}

	if responseBody.Criteria[0].Name == "" {
		t.Errorf("expected criterion name to be present")
	}

	if responseBody.Criteria[0].Explanation == "" {
		t.Errorf("expected criterion explanation to be present")
	}

	if len(responseBody.Risks) == 0 {
		t.Errorf("expected risks, got none")
	}
}