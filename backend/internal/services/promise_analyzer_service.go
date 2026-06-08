package services

import (
	"context"
	"strings"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
)

type PromiseAnalyzerService struct {
	scoreCalculator *ScoreCalculatorService
	llmClient       llm.Client
}

func NewPromiseAnalyzerService(llmClient llm.Client) *PromiseAnalyzerService {
	return &PromiseAnalyzerService{
		scoreCalculator: NewScoreCalculatorService(),
		llmClient:       llmClient,
	}
}

func (s *PromiseAnalyzerService) Analyze(
	ctx context.Context,
	text string,
) (domain.Analysis, error) {
	text = strings.TrimSpace(text)

	if len(text) < 15 {
		return domain.Analysis{
			Summary:    "Não foi possível identificar uma promessa pública clara no texto informado.",
			Score:      0,
			Confidence: 0,
			Criteria:   []domain.Criterion{},
			Risks: []string{
				"Informe uma promessa ou proposta pública mais específica.",
			},
		}, nil
	}

	summary := "Análise inicial baseada no modelo AletheIA v1."
	risks := []string{
		"Dados públicos adicionais são necessários para aprofundar a avaliação.",
	}

	var extraction *llm.PromiseExtraction

	if s.llmClient != nil {
		llmExtraction, err := s.llmClient.ExtractPromise(ctx, text)
		if err != nil {
			return domain.Analysis{}, err
		}

		extraction = llmExtraction

		if extraction.Summary != "" {
			summary = extraction.Summary
		}

		if len(extraction.Risks) > 0 {
			risks = extraction.Risks
		}
	}

	criteria := buildCriteria(extraction)
	criteria, score := s.scoreCalculator.Calculate(criteria)

	confidence := calculateConfidence(criteria)

	return domain.Analysis{
		Summary:    summary,
		Score:      score,
		Confidence: confidence,
		Criteria:   criteria,
		Risks:      risks,
	}, nil
}

func buildCriteria(
	extraction *llm.PromiseExtraction,
) []domain.Criterion {
	criteria := make([]domain.Criterion, 0, len(domain.ScoringModelV1))

	for _, modelCriterion := range domain.ScoringModelV1 {
		criterion := modelCriterion

		extractedCriterion := findExtractedCriterion(extraction, criterion.Key)

		if extractedCriterion != nil {
			criterion.Status = parseCriterionStatus(extractedCriterion.Status)
			criterion.Explanation = extractedCriterion.Explanation
		} else {
			criterion.Status = domain.CriterionStatusNo
			criterion.Explanation = "Critério não avaliado pela LLM."
		}

		criteria = append(criteria, criterion)
	}

	return criteria
}

func findExtractedCriterion(
	extraction *llm.PromiseExtraction,
	key string,
) *llm.ExtractedCriterion {
	if extraction == nil {
		return nil
	}

	for _, criterion := range extraction.Criteria {
		if criterion.Key == key {
			return &criterion
		}
	}

	return nil
}

func parseCriterionStatus(status string) domain.CriterionStatus {
	switch status {
	case "yes":
		return domain.CriterionStatusYes
	case "partial":
		return domain.CriterionStatusPartial
	case "no":
		return domain.CriterionStatusNo
	default:
		return domain.CriterionStatusNo
	}
}

func calculateConfidence(criteria []domain.Criterion) int {
	if len(criteria) == 0 {
		return 0
	}

	total := 0.0

	for _, criterion := range criteria {
		switch criterion.Status {
		case domain.CriterionStatusYes:
			total += 1
		case domain.CriterionStatusPartial:
			total += 0.5
		}
	}

	return int((total / float64(len(criteria))) * 100)
}