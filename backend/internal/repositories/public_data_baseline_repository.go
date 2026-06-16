package repositories

import (
	"context"
	"errors"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrPublicDataBaselineNotFound = errors.New(
	"public data baseline not found",
)

type PublicDataBaselineRepository struct {
	db *pgxpool.Pool
}

func NewPublicDataBaselineRepository(
	db *pgxpool.Pool,
) *PublicDataBaselineRepository {
	return &PublicDataBaselineRepository{
		db: db,
	}
}

func (r *PublicDataBaselineRepository) FindByIndicatorAndScope(
	ctx context.Context,
	indicator string,
	scope string,
) (*domain.PublicDataBaseline, error) {
	var baseline domain.PublicDataBaseline

	err := r.db.QueryRow(
		ctx,
		`
		SELECT
			id,
			indicator,
			scope,
			value,
			unit,
			source,
			reference,
			collected_at,
			created_at
		FROM public_data_baselines
		WHERE indicator = $1
		AND scope = $2
		LIMIT 1
		`,
		indicator,
		scope,
	).Scan(
		&baseline.ID,
		&baseline.Indicator,
		&baseline.Scope,
		&baseline.Value,
		&baseline.Unit,
		&baseline.Source,
		&baseline.Reference,
		&baseline.CollectedAt,
		&baseline.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrPublicDataBaselineNotFound
		}

		return nil, err
	}

	return &baseline, nil
}

func (r *PublicDataBaselineRepository) Save(
	ctx context.Context,
	baseline domain.PublicDataBaseline,
) error {
	_, err := r.db.Exec(
		ctx,
		`
		INSERT INTO public_data_baselines (
			indicator,
			scope,
			value,
			unit,
			source,
			reference,
			collected_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (indicator, scope)
		DO UPDATE SET
			value = EXCLUDED.value,
			unit = EXCLUDED.unit,
			source = EXCLUDED.source,
			reference = EXCLUDED.reference,
			collected_at = EXCLUDED.collected_at
		`,
		baseline.Indicator,
		baseline.Scope,
		baseline.Value,
		baseline.Unit,
		baseline.Source,
		baseline.Reference,
		baseline.CollectedAt,
	)

	return err
}