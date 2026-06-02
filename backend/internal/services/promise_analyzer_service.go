package services

import (
	"strings"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

type PromiseAnalyzerService struct {
	scoreCalculator *ScoreCalculatorService
}

func NewPromiseAnalyzerService() *PromiseAnalyzerService {
	return &PromiseAnalyzerService{
		scoreCalculator: NewScoreCalculatorService(),
	}
}

func (s *PromiseAnalyzerService) Analyze(text string) domain.Analysis {
	normalizedText := strings.ToLower(strings.TrimSpace(text))

	if len(normalizedText) < 15 {
		return domain.Analysis{
			Summary:    "Não foi possível identificar uma promessa pública clara no texto informado.",
			Score:      0,
			Confidence: 0,
			Criteria:   []domain.Criterion{},
			Risks: []string{
				"Informe uma promessa ou proposta pública mais específica.",
			},
		}
	}

	criteria := buildCriteria(normalizedText)
	criteria, score := s.scoreCalculator.Calculate(criteria)

	return domain.Analysis{
		Summary:    "Análise inicial baseada no modelo AletheIA v1.",
		Score:      score,
		Confidence: 30,
		Criteria:   criteria,
		Risks:      buildRisks(normalizedText),
	}
}

func buildCriteria(text string) []domain.Criterion {
	criteria := make([]domain.Criterion, 0, len(domain.ScoringModelV1))

	for _, modelCriterion := range domain.ScoringModelV1 {
		criterion := modelCriterion

		switch criterion.Key {
		case "clarity":
			criterion.Status = domain.CriterionStatusYes
			criterion.Explanation = "A promessa apresenta uma intenção pública identificável."

		case "measurability":
			criterion.Status = detectMeasurability(text)
			criterion.Explanation = "Avalia se a promessa possui meta ou resultado mensurável."

		case "deadline":
			criterion.Status = detectDeadline(text)
			criterion.Explanation = "Avalia se existe prazo explícito ou aproximado para execução."

		case "public_data":
			criterion.Status = domain.CriterionStatusPartial
			criterion.Explanation = "Nesta versão inicial, ainda não há consulta real a bases públicas."

		case "historical_baseline":
			criterion.Status = domain.CriterionStatusNo
			criterion.Explanation = "Ainda não foi consultada uma série histórica comparável."

		case "risks_dependencies":
			criterion.Status = detectRiskStatus(text)
			criterion.Explanation = "Avalia se existem riscos ou dependências aparentes de execução."
		}

		criteria = append(criteria, criterion)
	}

	return criteria
}

func detectMeasurability(text string) domain.CriterionStatus {
	if containsAny(text, []string{"%", "por cento", "mil", "milhão", "milhões", "reduzir", "aumentar", "construir"}) {
		return domain.CriterionStatusYes
	}

	return domain.CriterionStatusPartial
}

func detectDeadline(text string) domain.CriterionStatus {
	if containsAny(text, []string{"ano", "anos", "mês", "meses", "dias", "até", "mandato"}) {
		return domain.CriterionStatusYes
	}

	return domain.CriterionStatusNo
}

func detectRiskStatus(text string) domain.CriterionStatus {
	if containsAny(text, []string{"imposto", "saúde", "educação", "segurança", "obra", "hospital", "escola"}) {
		return domain.CriterionStatusPartial
	}

	return domain.CriterionStatusYes
}

func buildRisks(text string) []string {
	risks := []string{}

	if strings.Contains(text, "imposto") {
		risks = append(risks, "Possível impacto na arrecadação pública.")
	}

	if strings.Contains(text, "saúde") || strings.Contains(text, "hospital") {
		risks = append(risks, "Possível dependência de orçamento, equipe técnica e estrutura pública de saúde.")
	}

	if len(risks) == 0 {
		risks = append(risks, "Dados públicos adicionais são necessários para aprofundar a avaliação.")
	}

	return risks
}

func containsAny(text string, terms []string) bool {
	for _, term := range terms {
		if strings.Contains(text, term) {
			return true
		}
	}

	return false
}