package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/dto"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
	"github.com/flaviolpgjr/aletheia/backend/internal/security/captcha"
	"github.com/flaviolpgjr/aletheia/backend/internal/services"
)

type fakeLLMClient struct {
	err error
}

func (f *fakeLLMClient) ExtractPromise(
	ctx context.Context,
	text string,
) (*llm.PromiseExtraction, error) {
	if f.err != nil {
		return nil, f.err
	}

	return &llm.PromiseExtraction{
		Summary:     "Resumo gerado pela LLM.",
		TargetValue: 100,
		TargetUnit:  "hospitais",
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

type fakeCaptchaValidator struct {
	err error
}

func (f *fakeCaptchaValidator) Validate(
	ctx context.Context,
	token string,
) error {
	return f.err
}

func TestAnalyzePromiseHandler(t *testing.T) {
	service := services.NewPromiseAnalyzerService(
		&fakeLLMClient{},
		nil,
		nil,
	)

	handler := NewAnalyzePromiseHandler(
		service,
		&fakeCaptchaValidator{},
	)

	requestBody := dto.AnalyzePromiseRequest{
		Text:         "reduzir imposto sobre combustível",
		CaptchaToken: "valid-token",
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

	if responseBody.Score != 33 {
		t.Errorf(
			"expected score 33, got %d",
			responseBody.Score,
		)
	}

	if responseBody.Confidence != 35 {
		t.Errorf(
			"expected confidence 35, got %d",
			responseBody.Confidence,
		)
	}

	if responseBody.Summary != "Resumo gerado pela LLM." {
		t.Errorf(
			"expected summary from llm, got %s",
			responseBody.Summary,
		)
	}

	if responseBody.TargetValue != 100 {
		t.Errorf(
			"expected target value 100, got %f",
			responseBody.TargetValue,
		)
	}

	if responseBody.TargetUnit != "hospitais" {
		t.Errorf(
			"expected target unit hospitais, got %s",
			responseBody.TargetUnit,
		)
	}

	if len(responseBody.Criteria) != 7 {
		t.Errorf(
			"expected 7 criteria, got %d",
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

func TestAnalyzePromiseHandlerReturnsForbiddenWhenCaptchaIsMissing(t *testing.T) {
	service := services.NewPromiseAnalyzerService(
		&fakeLLMClient{},
		nil,
		nil,
	)

	handler := NewAnalyzePromiseHandler(
		service,
		&fakeCaptchaValidator{
			err: captcha.ErrInvalidCaptcha,
		},
	)

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

	if recorder.Code != http.StatusForbidden {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusForbidden,
			recorder.Code,
		)
	}
}

func TestAnalyzePromiseHandlerReturnsForbiddenWhenCaptchaIsInvalid(t *testing.T) {
	service := services.NewPromiseAnalyzerService(
		&fakeLLMClient{},
		nil,
		nil,
	)

	handler := NewAnalyzePromiseHandler(
		service,
		&fakeCaptchaValidator{
			err: captcha.ErrInvalidCaptcha,
		},
	)

	requestBody := dto.AnalyzePromiseRequest{
		Text:         "reduzir imposto sobre combustível",
		CaptchaToken: "invalid-token",
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

	if recorder.Code != http.StatusForbidden {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusForbidden,
			recorder.Code,
		)
	}
}

func TestAnalyzePromiseHandlerReturnsBadRequestWhenBodyIsInvalid(t *testing.T) {
	service := services.NewPromiseAnalyzerService(
		&fakeLLMClient{},
		nil,
		nil,
	)

	handler := NewAnalyzePromiseHandler(
		service,
		&fakeCaptchaValidator{},
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/promises/analyze",
		bytes.NewBufferString("{invalid-json"),
	)

	recorder := httptest.NewRecorder()

	handler.Handle(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}
}

func TestAnalyzePromiseHandlerReturnsTooManyRequestsWhenLLMRateLimit(t *testing.T) {
	service := services.NewPromiseAnalyzerService(
		&fakeLLMClient{
			err: llm.ErrRateLimit,
		},
		nil,
		nil,
	)

	handler := NewAnalyzePromiseHandler(
		service,
		&fakeCaptchaValidator{},
	)

	requestBody := dto.AnalyzePromiseRequest{
		Text:         "reduzir imposto sobre combustível",
		CaptchaToken: "valid-token",
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

	if recorder.Code != http.StatusTooManyRequests {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusTooManyRequests,
			recorder.Code,
		)
	}
}

func TestFakeCaptchaValidatorReturnsConfiguredError(t *testing.T) {
	expectedErr := captcha.ErrInvalidCaptcha

	validator := &fakeCaptchaValidator{
		err: expectedErr,
	}

	err := validator.Validate(
		context.Background(),
		"token",
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf("expected %v, got %v", expectedErr, err)
	}
}