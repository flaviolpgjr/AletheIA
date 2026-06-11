package services

import (
	"context"
	"errors"
	"strings"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
	"github.com/flaviolpgjr/aletheia/backend/internal/repositories"
	"github.com/flaviolpgjr/aletheia/backend/internal/utils"
)
type PromiseAnalyzerService struct {
	scoreCalculator    *ScoreCalculatorService
	llmClient          llm.Client
	analysisRepository AnalysisRepository
	publicDataProvider PublicDataProvider
}

type AnalysisRepository interface {
	FindByHash(ctx context.Context, promiseHash string) (*domain.Analysis, error)
	Save(ctx context.Context, promiseText string, promiseHash string, analysis domain.Analysis) error
}

type PublicDataProvider interface {
	FindEvidence(ctx context.Context, text string) ([]domain.Evidence, error)
}

func NewPromiseAnalyzerService(
	llmClient llm.Client,
	analysisRepository AnalysisRepository,
	publicDataProvider PublicDataProvider,
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

		analysis := domain.Analysis{
		Summary:    summary,
		Score:      score,
		Confidence: confidence,
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

	for _, modelCriterion := range domain.ScoringModelV1 {
		criterion := modelCriterion

		switch criterion.Key {

			case "evidence_plausibility":
				criterion.Status = evaluateEvidencePlausibility(evidence)
				criterion.Explanation = explainEvidencePlausibility(criterion.Status)

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

func evaluateEvidencePlausibility(evidence []domain.Evidence) domain.CriterionStatus {
	if len(evidence) == 0 {
		return domain.CriterionStatusNo
	}

	for _, item := range evidence {
		if item.Value > 0 {
			return domain.CriterionStatusPartial
		}
	}

	return domain.CriterionStatusPartial
}

func explainEvidencePlausibility(status domain.CriterionStatus) string {
	switch status {
	case domain.CriterionStatusYes:
		return "A promessa apresenta plausibilidade compatível com evidências públicas disponíveis."
	case domain.CriterionStatusPartial:
		return "Há evidências públicas relacionadas, mas ainda não há comparação quantitativa suficiente entre meta e linha de base."
	default:
		return "Não foram encontradas evidências públicas suficientes para avaliar a plausibilidade da promessa."
	}
}