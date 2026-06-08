package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const geminiGenerateContentURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.5-flash:generateContent"

type GeminiClient struct {
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

	fullPrompt := PromiseExtractionPrompt + "\n\nPROMESSA:\n" + text

	requestBody := geminiRequest{
		Contents: []geminiContent{
			{
				Parts: []geminiPart{
					{
						Text: fullPrompt,
					},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

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

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}

	var geminiResp geminiResponse

	err = json.Unmarshal(body, &geminiResp)
	if err != nil {
		return nil, err
	}

	if len(geminiResp.Candidates) == 0 {
		return nil, errors.New("gemini response has no candidates")
	}

	if len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return nil, errors.New("gemini response has no parts")
	}

	rawText := geminiResp.Candidates[0].Content.Parts[0].Text

	var extraction PromiseExtraction

	err = json.Unmarshal([]byte(rawText), &extraction)
	if err != nil {
		return nil, err
	}

	return &extraction, nil
}