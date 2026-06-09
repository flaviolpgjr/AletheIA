package llmfactory

import (
	"fmt"
	"log"

	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm/gemini"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm/openrouter"
)

type Config struct {
	Provider string

	GeminiAPIKey string

	OpenRouterAPIKey string
	OpenRouterModel  string
}

func NewClient(config Config) (llm.Client, error) {
	log.Printf("LLM_PROVIDER: %s", config.Provider)

	switch config.Provider {
	case "gemini":
		log.Printf("Using LLM provider: Gemini")
		log.Printf("GEMINI_API_KEY loaded: %v", config.GeminiAPIKey != "")

		return gemini.NewGeminiClient(config.GeminiAPIKey), nil

	case "openrouter":
		log.Printf("Using LLM provider: OpenRouter")
		log.Printf("OPENROUTER_API_KEY loaded: %v", config.OpenRouterAPIKey != "")
		log.Printf("OPENROUTER_MODEL: %s", config.OpenRouterModel)

		return openrouter.NewOpenRouterClient(
			config.OpenRouterAPIKey,
			config.OpenRouterModel,
		), nil

	default:
		return nil, fmt.Errorf("unsupported llm provider: %s", config.Provider)
	}
}