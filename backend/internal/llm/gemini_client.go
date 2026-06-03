package llm

import "context"



type GeminiClient struct{
	apiKey string
}

func NewGeminiClient(apiKey string) *GeminiClient {
	return &GeminiClient{
		apiKey: apiKey,
	}
}

func (g *GeminiClient) ExtractPromise(
	ctx context.Context,
	text string,
) (*PromiseExtraction, error) {
	return nil, nil
}