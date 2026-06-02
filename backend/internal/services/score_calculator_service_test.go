package services

import (
	"testing"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

func TestScoreCalculatorService_Calculate(t *testing.T) {
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