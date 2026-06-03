package llm

import "context"

type Client interface {
	ExtractPromise(
		ctx context.Context,
		text string,
	) (*PromiseExtraction, error)
}