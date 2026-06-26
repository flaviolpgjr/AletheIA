package captcha

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var ErrInvalidCaptcha = errors.New("invalid captcha")

const defaultTurnstileVerifyURL = "https://challenges.cloudflare.com/turnstile/v0/siteverify"

type TurnstileValidator struct {
	secretKey  string
	verifyURL  string
	httpClient *http.Client
}

type turnstileResponse struct {
	Success bool     `json:"success"`
	Errors  []string `json:"error-codes"`
}

func NewTurnstileValidator(secretKey string) *TurnstileValidator {
	return &TurnstileValidator{
		secretKey: secretKey,
		verifyURL: defaultTurnstileVerifyURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func NewTurnstileValidatorWithURL(
	secretKey string,
	verifyURL string,
) *TurnstileValidator {
	validator := NewTurnstileValidator(secretKey)
	validator.verifyURL = verifyURL

	return validator
}

func (v *TurnstileValidator) Validate(
	ctx context.Context,
	token string,
) error {
	token = strings.TrimSpace(token)

	if token == "" {
		return ErrInvalidCaptcha
	}

	form := url.Values{}
	form.Set("secret", v.secretKey)
	form.Set("response", token)

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		v.verifyURL,
		strings.NewReader(form.Encode()),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := v.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result turnstileResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if !result.Success {
		return ErrInvalidCaptcha
	}

	return nil
}