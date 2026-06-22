package services

import (
	"testing"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
)

func TestEvaluateEvidencePlausibility(t *testing.T) {
	tests := []struct {
		name     string
		target   float64
		baseline float64
		expected domain.CriterionStatus
	}{
		{
			name:     "high plausibility",
			target:   100,
			baseline: 5115,
			expected: domain.CriterionStatusYes,
		},
		{
			name:     "medium plausibility",
			target:   1000,
			baseline: 5115,
			expected: domain.CriterionStatusPartial,
		},
		{
			name:     "low plausibility",
			target:   10000,
			baseline: 5115,
			expected: domain.CriterionStatusNo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := evaluateEvidencePlausibility(
				&llm.PromiseExtraction{
					TargetValue: tt.target,
				},
				[]domain.Evidence{
					{
						Value: tt.baseline,
					},
				},
			)

			if status != tt.expected {
				t.Fatalf(
					"expected %s, got %s",
					tt.expected,
					status,
				)
			}
		})
	}
}