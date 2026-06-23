package publicdata

import (
	"context"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
)

type Provider interface {
	FindEvidence(
		ctx context.Context,
		text string,
	) ([]domain.Evidence, error)
}