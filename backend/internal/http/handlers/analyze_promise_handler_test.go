package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/dto"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
	"github.com/flaviolpgjr/aletheia/backend/internal/services"
)

type fakeLLMClient struct{}

func (f *fakeLLMClient) ExtractPromise(
	ctx context.Context,
	text string,
) (*llm.PromiseExtraction, error) {
	return &llm.PromiseExtraction{
		Summary: "Resumo gerado pela LLM.",
		Risks: []string{
			"Risco identificado pela LLM.",
		},
		Criteria: []llm.ExtractedCriterion{
			{
				Key:         "clarity",
				Status:      "yes",
				Explanation: "A promessa é clara.",
			},
			{
				Key:         "measurability",
				Status:      "yes",
				Explanation: "A promessa é mensurável.",
			},
			{
				Key:         "deadline",
				Status:      "no",
				Explanation: "A promessa não possui prazo.",
			},
			{
				Key:         "public_data",
				Status:      "partial",
				Explanation: "Existem dados públicos parciais.",
			},
			{
				Key:         "historical_baseline",
				Status:      "no",
				Explanation: "Não há histórico comparável.",
			},
			{
				Key:         "risks_dependencies",
				Status:      "partial",
				Explanation: "Existem riscos parciais.",
			},
		},
	}, nil
}

func TestAnalyzePromiseHandler(t *testing.T) {
	service := services.NewPromiseAnalyzerService(&fakeLLMClient{})
	handler := NewAnalyzePromiseHandler(service)

	requestBody := dto.AnalyzePromiseRequest{
		Text: "reduzir imposto sobre combustível",
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to marshal request: %v", err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/promises/analyze",
		bytes.NewBuffer(body),
	)

	recorder := httptest.NewRecorder()

	handler.Handle(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			recorder.Code,
		)
	}

	var responseBody dto.AnalyzePromiseResponse

	err = json.NewDecoder(recorder.Body).Decode(&responseBody)
	if err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if responseBody.Score != 55 {
		t.Errorf(
			"expected score 55, got %d",
			responseBody.Score,
		)
	}

	if responseBody.Confidence != 50 {
		t.Errorf(
			"expected confidence 50, got %d",
			responseBody.Confidence,
		)
	}

	if responseBody.Summary != "Resumo gerado pela LLM." {
		t.Errorf(
			"expected summary from llm, got %s",
			responseBody.Summary,
		)
	}

	if len(responseBody.Criteria) != 6 {
		t.Errorf(
			"expected 6 criteria, got %d",
			len(responseBody.Criteria),
		)
	}

	if len(responseBody.Risks) != 1 {
		t.Errorf(
			"expected 1 risk, got %d",
			len(responseBody.Risks),
		)
	}
}