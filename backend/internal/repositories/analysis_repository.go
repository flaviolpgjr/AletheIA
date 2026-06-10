package repositories

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/flaviolpgjr/aletheia/backend/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrAnalysisNotFound = errors.New("analysis not found")

type AnalysisRepository struct {
	db *pgxpool.Pool
}

func NewAnalysisRepository(db *pgxpool.Pool) *AnalysisRepository {
	return &AnalysisRepository{
		db: db,
	}
}

func (r *AnalysisRepository) FindByHash(
	ctx context.Context,
	promiseHash string,
) (*domain.Analysis, error) {
	var analysisData []byte

	err := r.db.QueryRow(
		ctx,
		`
		SELECT analysis_data
		FROM analyses
		WHERE promise_hash = $1
		LIMIT 1
		`,
		promiseHash,
	).Scan(&analysisData)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrAnalysisNotFound
		}

		return nil, err
	}

	var analysis domain.Analysis

	if err := json.Unmarshal(analysisData, &analysis); err != nil {
		return nil, err
	}

	return &analysis, nil
}

func (r *AnalysisRepository) Save(
	ctx context.Context,
	promiseText string,
	promiseHash string,
	analysis domain.Analysis,
) error {
	analysisData, err := json.Marshal(analysis)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		ctx,
		`
		INSERT INTO analyses (
			id,
			promise_text,
			promise_hash,
			score,
			confidence,
			analysis_data
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (promise_hash) DO NOTHING
		`,
		uuid.NewString(),
		promiseText,
		promiseHash,
		analysis.Score,
		analysis.Confidence,
		analysisData,
	)

	return err
}