package gemini

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

const geminiGenerateContentURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent"

type GeminiClient struct {
	apiKey string
}

var _ llm.Client = (*GeminiClient)(nil)

func NewGeminiClient(apiKey string) *GeminiClient {
	return &GeminiClient{
		apiKey: apiKey,
	}
}

func (g *GeminiClient) ExtractPromise(
	ctx context.Context,
	text string,
) (*llm.PromiseExtraction, error) {
	requestBody := buildGeminiRequest(text)

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := g.buildHTTPRequest(ctx, jsonBody)
	if err != nil {
		return nil, err
	}

	body, err := executeGeminiRequest(req)
	if err != nil {
		return nil, err
	}

	extraction, err := parseGeminiExtraction(body)
	if err != nil {
		return nil, err
	}

	return extraction, nil
}

func buildGeminiRequest(text string) geminiRequest {
	fullPrompt := llm.PromiseExtractionPrompt + "\n\nPROMESSA:\n" + text

	return geminiRequest{
		Contents: []geminiContent{
			{
				Parts: []geminiPart{
					{
						Text: fullPrompt,
					},
				},
			},
		},
		GenerationConfig: geminiGenerationConfig{
			Temperature:      0,
			TopP:             1,
			TopK:             1,
			ResponseMimeType: "application/json",
		},
	}
}

func (g *GeminiClient) buildHTTPRequest(
	ctx context.Context,
	jsonBody []byte,
) (*http.Request, error) {
	url := geminiGenerateContentURL + "?key=" + g.apiKey

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func executeGeminiRequest(req *http.Request) ([]byte, error) {
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
		log.Println("Gemini rate limit reached")
		return nil, llm.ErrRateLimit
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("Gemini request failed with status %d\n", resp.StatusCode)
		return nil, fmt.Errorf("gemini request failed with status %d", resp.StatusCode)
	}

	return body, nil
}

func parseGeminiExtraction(body []byte) (*llm.PromiseExtraction, error) {
	var geminiResp geminiResponse

	if err := json.Unmarshal(body, &geminiResp); err != nil {
		return nil, err
	}

	if len(geminiResp.Candidates) == 0 {
		return nil, errors.New("gemini response has no candidates")
	}

	if len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("gemini response has no parts")
	}

	rawText := geminiResp.Candidates[0].Content.Parts[0].Text

	var extraction llm.PromiseExtraction

	if err := json.Unmarshal([]byte(rawText), &extraction); err != nil {
		return nil, err
	}

	return &extraction, nil
}