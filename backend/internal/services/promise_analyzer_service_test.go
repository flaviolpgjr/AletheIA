package services

import (
	"context"
	"testing"

	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
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
				Explanation: "Não há histórico comparável informado.",
			},
			{
				Key:         "risks_dependencies",
				Status:      "partial",
				Explanation: "Existem riscos ou dependências parciais.",
			},
		},
	}, nil
}

func TestAnalyzeWhenPromiseMentionsTax(t *testing.T) {
	service := NewPromiseAnalyzerService(&fakeLLMClient{})

	analysis, err := service.Analyze(
		context.Background(),
		"reduzir imposto sobre combustível",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if analysis.Score != 55 {
		t.Errorf("expected score 55, got %d", analysis.Score)
	}

	if analysis.Confidence != 50 {
		t.Errorf("expected confidence 50, got %d", analysis.Confidence)
	}

	if len(analysis.Criteria) != 6 {
		t.Errorf("expected 6 criteria, got %d", len(analysis.Criteria))
	}

	if analysis.Summary != "Resumo gerado pela LLM." {
		t.Errorf("expected LLM summary, got %s", analysis.Summary)
	}

	if analysis.Risks[0] != "Risco identificado pela LLM." {
		t.Errorf("expected LLM risk, got %s", analysis.Risks[0])
	}
}

func TestAnalyzeWhenPromiseHasNoKnownKeywords(t *testing.T) {
	service := NewPromiseAnalyzerService(&fakeLLMClient{})

	analysis, err := service.Analyze(
		context.Background(),
		"melhorar a qualidade dos serviços públicos",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if analysis.Score != 45 {
		t.Errorf("expected score 45, got %d", analysis.Score)
	}

	if analysis.Confidence != 41 {
		t.Errorf("expected confidence 41, got %d", analysis.Confidence)
	}

	if len(analysis.Criteria) != 6 {
		t.Errorf("expected 6 criteria, got %d", len(analysis.Criteria))
	}

	if analysis.Risks[0] != "Risco identificado pela LLM." {
		t.Errorf("expected LLM risk, got %s", analysis.Risks[0])
	}
}