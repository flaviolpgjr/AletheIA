package services

import (
	"context"
	"testing"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
)

type fakeLLMClient struct{}

type fakePublicDataProvider struct{}

func (f *fakePublicDataProvider) FindEvidence(
	ctx context.Context,
	text string,
) ([]domain.Evidence, error) {
	return []domain.Evidence{
		{
			Indicator: "hospital_facilities",
			Value:     5115,
			Unit:      "hospitais",
			Reference: "Cadastro Nacional de Estabelecimentos de Saúde (CNES) - Quantidade de hospitais ativos no Brasil",
		},
	}, nil
}

func (f *fakeLLMClient) ExtractPromise(
	ctx context.Context,
	text string,
) (*llm.PromiseExtraction, error) {
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
	service := NewPromiseAnalyzerService(&fakeLLMClient{}, nil, nil)

	analysis, err := service.Analyze(
		context.Background(),
		"reduzir imposto sobre combustível",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if analysis.Score != 33 {
		t.Errorf("expected score 33, got %d", analysis.Score)
	}

	if analysis.Confidence != 35 {
		t.Errorf("expected confidence 35, got %d", analysis.Confidence)
	}

	if len(analysis.Criteria) != 7 {
		t.Errorf("expected 7 criteria, got %d", len(analysis.Criteria))
	}

	if analysis.Summary != "Resumo gerado pela LLM." {
		t.Errorf("expected LLM summary, got %s", analysis.Summary)
	}

	if analysis.Risks[0] != "Risco identificado pela LLM." {
		t.Errorf("expected LLM risk, got %s", analysis.Risks[0])
	}
}

func TestAnalyzeWhenPromiseHasNoKnownKeywords(t *testing.T) {
	service := NewPromiseAnalyzerService(&fakeLLMClient{}, nil, nil)

	analysis, err := service.Analyze(
		context.Background(),
		"melhorar a qualidade dos serviços públicos",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if analysis.Score != 18 {
		t.Errorf("expected score 18, got %d", analysis.Score)
	}

	if analysis.Confidence != 21 {
		t.Errorf("expected confidence 21, got %d", analysis.Confidence)
	}

	if len(analysis.Criteria) != 7 {
		t.Errorf("expected 7 criteria, got %d", len(analysis.Criteria))
	}

	if analysis.Risks[0] != "Risco identificado pela LLM." {
		t.Errorf("expected LLM risk, got %s", analysis.Risks[0])
	}
}

func TestAnalyzeAddsEvidencePlausibilityCriterion(t *testing.T) {
	service := NewPromiseAnalyzerService(
		&fakeLLMClient{},
		nil,
		&fakePublicDataProvider{},
	)

	analysis, err := service.Analyze(
		context.Background(),
		"construir 100 hospitais",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	criterion := findCriterionByKeyForTest(
		analysis.Criteria,
		"evidence_plausibility",
	)

	if criterion == nil {
		t.Fatal("expected evidence_plausibility criterion")
	}

	if criterion.Status != domain.CriterionStatusYes {
		t.Fatalf(
			"expected yes, got %s",
			criterion.Status,
		)
	}
}

func TestAnalyzeIncludesExtractedTarget(t *testing.T) {
	service := NewPromiseAnalyzerService(&fakeLLMClient{}, nil, nil)

	analysis, err := service.Analyze(
		context.Background(),
		"construir 100 hospitais",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if analysis.TargetValue != 100 {
		t.Errorf("expected target value 100, got %f", analysis.TargetValue)
	}

	if analysis.TargetUnit != "hospitais" {
		t.Errorf("expected target unit hospitais, got %s", analysis.TargetUnit)
	}
}

func TestAnalyzeSetsMeasurabilityNoWhenPromiseIsSubjectiveWithoutNumber(t *testing.T) {
	service := NewPromiseAnalyzerService(&fakeLLMClient{}, nil, nil)

	analysis, err := service.Analyze(
		context.Background(),
		"Vou melhorar a felicidade dos brasileiros",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	criterion := findCriterionByKeyForTest(
		analysis.Criteria,
		"measurability",
	)

	if criterion == nil {
		t.Fatal("expected measurability criterion")
	}

	if criterion.Status != domain.CriterionStatusNo {
		t.Fatalf(
			"expected no, got %s",
			criterion.Status,
		)
	}

	expectedExplanation := "A promessa não possui um indicador mensurável claramente definido."
	if criterion.Explanation != expectedExplanation {
		t.Fatalf(
			"expected explanation %q, got %q",
			expectedExplanation,
			criterion.Explanation,
		)
	}
}

func TestAnalyzeSetsMeasurabilityPartialWhenPromiseIsSubjectiveWithNumber(t *testing.T) {
	service := NewPromiseAnalyzerService(&fakeLLMClient{}, nil, nil)

	analysis, err := service.Analyze(
		context.Background(),
		"Vou aumentar a felicidade dos brasileiros em 20%",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	criterion := findCriterionByKeyForTest(
		analysis.Criteria,
		"measurability",
	)

	if criterion == nil {
		t.Fatal("expected measurability criterion")
	}

	if criterion.Status != domain.CriterionStatusPartial {
		t.Fatalf(
			"expected partial, got %s",
			criterion.Status,
		)
	}

	expectedExplanation := "A promessa possui uma meta quantitativa, porém o indicador é subjetivo e de difícil verificação."
	if criterion.Explanation != expectedExplanation {
		t.Fatalf(
			"expected explanation %q, got %q",
			expectedExplanation,
			criterion.Explanation,
		)
	}
}

func TestAnalyzeWithoutEvidenceSetsPublicDataCriteriaToNo(t *testing.T) {
	service := NewPromiseAnalyzerService(&fakeLLMClient{}, nil, nil)

	analysis, err := service.Analyze(
		context.Background(),
		"Vou reduzir impostos federais em 20%",
	)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedStatuses := map[string]domain.CriterionStatus{
		"public_data":            domain.CriterionStatusNo,
		"historical_baseline":    domain.CriterionStatusNo,
		"evidence_plausibility":  domain.CriterionStatusNo,
	}

	for key, expectedStatus := range expectedStatuses {
		criterion := findCriterionByKeyForTest(analysis.Criteria, key)
		if criterion == nil {
			t.Fatalf("expected %s criterion", key)
		}

		if criterion.Status != expectedStatus {
			t.Fatalf(
				"expected %s status %s, got %s",
				key,
				expectedStatus,
				criterion.Status,
			)
		}
	}
}

func findCriterionByKeyForTest(
	criteria []domain.Criterion,
	key string,
) *domain.Criterion {
	for i := range criteria {
		if criteria[i].Key == key {
			return &criteria[i]
		}
	}

	return nil
}