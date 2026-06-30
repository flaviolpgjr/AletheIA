package services

import (
	"testing"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

func TestScoreCalculatorServiceCalculate(t *testing.T) {
	calculator := NewScoreCalculatorService()

	criteria := []domain.Criterion{
		{
			Key:    "clarity",
			Weight: 15,
			Status: domain.CriterionStatusYes,
		},
		{
			Key:    "measurability",
			Weight: 20,
			Status: domain.CriterionStatusPartial,
		},
		{
			Key:    "deadline",
			Weight: 10,
			Status: domain.CriterionStatusNo,
		},
	}

	updatedCriteria, score := calculator.Calculate(criteria)

	if score != 25 {
		t.Errorf("expected score 25, got %d", score)
	}

	if updatedCriteria[0].Score != 15 {
		t.Errorf("expected first criterion score 15, got %.2f", updatedCriteria[0].Score)
	}

	if updatedCriteria[1].Score != 10 {
		t.Errorf("expected second criterion score 10, got %.2f", updatedCriteria[1].Score)
	}

	if updatedCriteria[2].Score != 0 {
		t.Errorf("expected third criterion score 0, got %.2f", updatedCriteria[2].Score)
	}
}

func TestScoreCalculatorServiceCalculateWithEmptyCriteria(t *testing.T) {
	calculator := NewScoreCalculatorService()

	updatedCriteria, score := calculator.Calculate([]domain.Criterion{})

	if score != 0 {
		t.Errorf("expected score 0, got %d", score)
	}

	if len(updatedCriteria) != 0 {
		t.Errorf("expected empty criteria, got %d", len(updatedCriteria))
	}
}

func TestScoreCalculatorServiceCalculateRoundsFinalScore(t *testing.T) {
	calculator := NewScoreCalculatorService()

	criteria := []domain.Criterion{
		{
			Key:    "risks_dependencies",
			Weight: 15,
			Status: domain.CriterionStatusPartial,
		},
	}

	updatedCriteria, score := calculator.Calculate(criteria)

	if score != 8 {
		t.Errorf("expected score 8, got %d", score)
	}

	if updatedCriteria[0].Score != 7.5 {
		t.Errorf("expected criterion score 7.5, got %.2f", updatedCriteria[0].Score)
	}
}