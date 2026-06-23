package publicdata

import (
	"context"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

type Aggregator struct {
	providers []Provider
}

func NewAggregator(providers ...Provider) *Aggregator {
	return &Aggregator{
		providers: providers,
	}
}

func (a *Aggregator) FindEvidence(
	ctx context.Context,
	text string,
) ([]domain.Evidence, error) {
	evidence := []domain.Evidence{}

	for _, provider := range a.providers {
		foundEvidence, err := provider.FindEvidence(
			ctx,
			text,
		)
		if err != nil {
			return nil, err
		}

		evidence = append(
			evidence,
			foundEvidence...,
		)
	}

	return evidence, nil
}