package services

import (
	"math"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

type ScoreCalculatorService struct{}

func NewScoreCalculatorService() *ScoreCalculatorService {
	return &ScoreCalculatorService{}
}

func (s *ScoreCalculatorService) Calculate(criteria []domain.Criterion) ([]domain.Criterion, int) {
	total := 0.0

	for i := range criteria {
		score := float64(criteria[i].Weight) * statusFactor(criteria[i].Status)

		criteria[i].Score = score
		total += score
	}

	return criteria, int(math.Round(total))
}

func statusFactor(status domain.CriterionStatus) float64 {
	switch status {
	case domain.CriterionStatusYes:
		return 1
	case domain.CriterionStatusPartial:
		return 0.5
	default:
		return 0
	}
}

func evidenceScore(evidence []domain.Evidence, promiseText string) float64 {
	for _, e := range evidence {
		if e.Indicator == "health_facilities" && e.Value > 0 {
			// depois vamos extrair a meta da promessa
			return 1
		}
	}

	return 0.5
}