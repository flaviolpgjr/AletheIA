package services

import (
	"strings"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

type PromiseAnalyzerService struct{}

func NewPromiseAnalyzerService() *PromiseAnalyzerService {
	return &PromiseAnalyzerService{}
}

func (s *PromiseAnalyzerService) Analyze(text string) domain.Analysis {
	if len(strings.TrimSpace(text)) < 15 {
		return domain.Analysis{
			Summary: "Não foi possível identificar uma promessa pública clara no texto informado.",
			Score:   -1,
			Risks: []string{
				"Informe uma promessa ou proposta pública mais específica.",
			},
		}
	}
	
	risks := []string{}
	score := 75
	summary := "Promessa com viabilidade inicial moderada."

	normalizedText := strings.ToLower(text)

	if strings.Contains(normalizedText, "imposto") {
		risks = append(risks, "Possível redução de arrecadação")
		score -= 15
	}

	if strings.Contains(normalizedText, "saúde") {
		risks = append(risks, "Possível aumento de gasto público")
		score -= 10
	}

	if len(risks) == 0 {
		risks = append(risks, "Dados adicionais necessários para avaliação")
	}


	return domain.Analysis{
		Summary: summary,
		Score:   score,
		Risks:   risks,
	}
}
