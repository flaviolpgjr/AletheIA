package openrouter

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
)

const openRouterChatCompletionsURL = "https://openrouter.ai/api/v1/chat/completions"

type OpenRouterClient struct {
	apiKey string
	model  string
}

var _ llm.Client = (*OpenRouterClient)(nil)

func NewOpenRouterClient(apiKey string, model string) *OpenRouterClient {
	return &OpenRouterClient{
		apiKey: apiKey,
		model:  model,
	}
}

func (o *OpenRouterClient) ExtractPromise(
	ctx context.Context,
	text string,
) (*llm.PromiseExtraction, error) {
	requestBody := buildOpenRouterRequest(o.model, text)

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}
	
	req, err := o.buildHTTPRequest(ctx, jsonBody)
	if err != nil {
		return nil, err
	}

	body, err := executeOpenRouterRequest(req)
	if err != nil {
		return nil, err
	}

	extraction, err := parseOpenRouterExtraction(body)
	if err != nil {
		return nil, err
	}

	return extraction, nil
}

func buildOpenRouterRequest(model string, text string) openRouterRequest {
	fullPrompt := llm.PromiseExtractionPrompt + "\n\nPROMESSA:\n" + text

	return openRouterRequest{
		Model: model,
		Messages: []openRouterMessage{
			{
				Role:    "user",
				Content: fullPrompt,
			},
		},
		Temperature: 0,
		ResponseFormat: openRouterResponseFormat{
			Type: "json_object",
		},
	}
}

func (o *OpenRouterClient) buildHTTPRequest(
	ctx context.Context,
	jsonBody []byte,
) (*http.Request, error) {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		openRouterChatCompletionsURL,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.apiKey)

	return req, nil
}

func executeOpenRouterRequest(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusTooManyRequests {
		log.Println("OpenRouter rate limit reached")
		return nil, llm.ErrRateLimit
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("OpenRouter request failed with status %d: %s\n", resp.StatusCode, string(body))
		return nil, fmt.Errorf("openrouter request failed with status %d", resp.StatusCode)
	}

	return body, nil
}

func parseOpenRouterExtraction(body []byte) (*llm.PromiseExtraction, error) {
	var openRouterResp openRouterResponse

	if err := json.Unmarshal(body, &openRouterResp); err != nil {
		return nil, err
	}

	if len(openRouterResp.Choices) == 0 {
		return nil, errors.New("openrouter response has no choices")
	}

	rawText := openRouterResp.Choices[0].Message.Content

	var extraction llm.PromiseExtraction

	if err := json.Unmarshal([]byte(rawText), &extraction); err != nil {
		return nil, err
	}

	return &extraction, nil
}