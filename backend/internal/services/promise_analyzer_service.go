package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
	"github.com/flaviolpgjr/aletheia/backend/internal/publicdata"
	"github.com/flaviolpgjr/aletheia/backend/internal/repositories"
	"github.com/flaviolpgjr/aletheia/backend/internal/utils"
)
type PromiseAnalyzerService struct {
	scoreCalculator    *ScoreCalculatorService
	llmClient          llm.Client
	analysisRepository AnalysisRepository
	publicDataProvider publicdata.Provider
}

type AnalysisRepository interface {
	FindByHash(ctx context.Context, promiseHash string) (*domain.Analysis, error)
	Save(ctx context.Context, promiseText string, promiseHash string, analysis domain.Analysis) error
}


func NewPromiseAnalyzerService(
	llmClient llm.Client,
	analysisRepository AnalysisRepository,
	publicDataProvider publicdata.Provider,
) *PromiseAnalyzerService {
	return &PromiseAnalyzerService{
		scoreCalculator:     NewScoreCalculatorService(),
		llmClient:           llmClient,
		analysisRepository:  analysisRepository,
		publicDataProvider:  publicDataProvider,
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

	promiseHash := utils.GeneratePromiseHash(text)

	if s.analysisRepository != nil {
		cachedAnalysis, err := s.analysisRepository.FindByHash(ctx, promiseHash)
		if err == nil {
			return *cachedAnalysis, nil
		}

		if !errors.Is(err, repositories.ErrAnalysisNotFound) {
			return domain.Analysis{}, err
		}
	}

	evidence := []domain.Evidence{}

	if s.publicDataProvider != nil {
		foundEvidence, err := s.publicDataProvider.FindEvidence(ctx, text)
		if err != nil {
			return domain.Analysis{}, err
		}

		evidence = foundEvidence
	}

	summary := "Análise inicial baseada no modelo AletheIA v1."
	risks := []string{
		"Dados públicos adicionais são necessários para aprofundar a avaliação.",
	}

	sources := []domain.PublicSource{}

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

		for _, source := range extraction.SuggestedSources {
			sources = append(sources, domain.PublicSource{
				Name:        source.Name,
				Description: source.Description,
			})
		}
	}

	criteria := buildCriteria(text, extraction, evidence)
	criteria, score := s.scoreCalculator.Calculate(criteria)

	confidence := calculateConfidence(criteria)

	targetValue := 0.0
	targetUnit := ""

	if extraction != nil {
		targetValue = extraction.TargetValue
		targetUnit = extraction.TargetUnit
	}

	analysis := domain.Analysis{
		Summary:    summary,
		Score:      score,
		Confidence: confidence,
		TargetValue: targetValue,
		TargetUnit:  targetUnit,
		Criteria:   criteria,
		Risks:      risks,
		Sources:    sources,
		Evidence:   evidence,
	}

	if s.analysisRepository != nil {
		if err := s.analysisRepository.Save(ctx, text, promiseHash, analysis); err != nil {
			return domain.Analysis{}, err
		}
	}

	return analysis, nil
}

func buildCriteria(
	text string,
	extraction *llm.PromiseExtraction,
	evidence []domain.Evidence,
) []domain.Criterion {
	criteria := make([]domain.Criterion, 0, len(domain.ScoringModelV1))
	hasEvidence := len(evidence) > 0

	for _, modelCriterion := range domain.ScoringModelV1 {
		criterion := modelCriterion

		switch criterion.Key {
		case "public_data":
			if !hasEvidence {
				criterion.Status = domain.CriterionStatusNo
				criterion.Explanation = "Não foram encontradas evidências públicas integradas ao AletheIA para avaliar esta promessa."

				criteria = append(criteria, criterion)
				continue
			}

		case "historical_baseline":
			if !hasEvidence {
				criterion.Status = domain.CriterionStatusNo
				criterion.Explanation = "Não há linha de base pública integrada ao AletheIA para comparar esta promessa."

				criteria = append(criteria, criterion)
				continue
			}

		case "evidence_plausibility":
			criterion.Status = evaluateEvidencePlausibility(extraction, evidence)
			criterion.Explanation = explainEvidencePlausibility(
				extraction,
				evidence,
				criterion.Status,
			)

			criteria = append(criteria, criterion)
			continue

		case "deadline":
			if extraction != nil &&
				strings.TrimSpace(extraction.Deadline) != "" {

				criterion.Status = domain.CriterionStatusYes
				criterion.Explanation = "Prazo identificado automaticamente pelo AletheIA."

				criteria = append(criteria, criterion)
				continue
			}

		case "measurability":
			if hasVagueMeasurement(text, extraction) {
				criterion.Status = domain.CriterionStatusPartial
				criterion.Explanation = "A promessa possui número, mas o indicador não está claramente definido ou é subjetivo."

				criteria = append(criteria, criterion)
				continue
			}

			if hasMeasurement(text) && hasIndicators(extraction) {
				criterion.Status = domain.CriterionStatusYes
				criterion.Explanation = "Indicador quantitativo e fonte de medição identificados automaticamente pelo AletheIA."

				criteria = append(criteria, criterion)
				continue
			}

			if hasMeasurement(text) || hasIndicators(extraction) {
				criterion.Status = domain.CriterionStatusPartial
				criterion.Explanation = "A promessa possui algum elemento mensurável, mas ainda depende de definição mais clara do indicador."

				criteria = append(criteria, criterion)
				continue
			}
		}

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

func hasMeasurement(text string) bool {
	for _, r := range text {
		if r >= '0' && r <= '9' {
			return true
		}
	}

	return false
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

func hasIndicators(extraction *llm.PromiseExtraction) bool {
	if extraction == nil {
		return false
	}

	return len(extraction.Indicators) > 0
}

func hasVagueMeasurement(text string, extraction *llm.PromiseExtraction) bool {
	text = strings.ToLower(text)

	vagueTerms := []string{
		"felicidade",
		"confiança",
		"qualidade de vida",
		"eficiência",
		"prosperidade",
		"inovação",
		"melhorar",
		"transformar",
	}

	for _, term := range vagueTerms {
		if strings.Contains(text, term) {
			return true
		}
	}

	if extraction == nil {
		return false
	}

	for _, indicator := range extraction.Indicators {
		indicator = strings.ToLower(indicator)

		for _, term := range vagueTerms {
			if strings.Contains(indicator, term) {
				return true
			}
		}
	}

	return false
}

func evaluateEvidencePlausibility(
	extraction *llm.PromiseExtraction,
	evidence []domain.Evidence,
) domain.CriterionStatus {
	if len(evidence) == 0 {
		return domain.CriterionStatusNo
	}

	if extraction == nil || extraction.TargetValue <= 0 {
		return domain.CriterionStatusPartial
	}

	for _, item := range evidence {
		if item.Value <= 0 {
			continue
		}

		ratio := extraction.TargetValue / item.Value

		switch {
		case ratio <= 0.10:
			return domain.CriterionStatusYes

		case ratio <= 0.50:
			return domain.CriterionStatusPartial

		default:
			return domain.CriterionStatusNo
		}
	}

	return domain.CriterionStatusPartial
}

func explainEvidencePlausibility(
	extraction *llm.PromiseExtraction,
	evidence []domain.Evidence,
	status domain.CriterionStatus,
) string {
	if len(evidence) == 0 {
		return "Não foram encontradas evidências públicas suficientes para avaliar a plausibilidade da promessa."
	}

	if extraction == nil || extraction.TargetValue <= 0 {
		return "Há evidências públicas relacionadas, mas a meta quantitativa não foi identificada com segurança para comparação com a linha de base."
	}

	for _, item := range evidence {
		if item.Value <= 0 {
			continue
		}

		percentage := (extraction.TargetValue / item.Value) * 100

		ratio := extraction.TargetValue / item.Value

		statusDescription := "plausibilidade moderada"

		switch {
		case ratio <= 0.10:
			statusDescription = "plausibilidade alta"

		case ratio <= 0.50:
			statusDescription = "plausibilidade moderada"

		default:
			statusDescription = "plausibilidade baixa"
		}

		return fmt.Sprintf(
			"A meta proposta é de %.0f %s. A linha de base pública identificada é de %.0f %s, segundo %s. A meta representa aproximadamente %.2f%% da linha de base atual, indicando %s. A avaliação continua limitada por fatores como orçamento, execução, prazo e distribuição regional.",
			extraction.TargetValue,
			extraction.TargetUnit,
			item.Value,
			item.Unit,
			item.Reference,
			percentage,
			statusDescription,
		)
	}

	if status == domain.CriterionStatusPartial {
		return "Há evidências públicas relacionadas, mas a linha de base encontrada não possui valor quantitativo suficiente para comparação."
	}

	return "Não foram encontradas evidências públicas suficientes para avaliar a plausibilidade da promessa."
}