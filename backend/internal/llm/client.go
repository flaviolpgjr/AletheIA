package llm

import (
	"context"
	"errors"
)

var ErrRateLimit = errors.New("llm rate limit exceeded")

type Client interface {
	ExtractPromise(ctx context.Context, text string) (*PromiseExtraction, error)
}