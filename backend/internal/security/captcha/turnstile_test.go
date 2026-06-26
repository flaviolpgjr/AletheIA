package captcha

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTurnstileValidatorReturnsInvalidCaptchaWhenTokenIsEmpty(t *testing.T) {
	validator := NewTurnstileValidatorWithURL(
		"secret-key",
		"http://example.com",
	)

	err := validator.Validate(
		context.Background(),
		"",
	)

	if !errors.Is(err, ErrInvalidCaptcha) {
		t.Fatalf(
			"expected ErrInvalidCaptcha, got %v",
			err,
		)
	}
}

func TestTurnstileValidatorReturnsNilWhenCaptchaIsValid(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Fatalf(
					"expected method POST, got %s",
					r.Method,
				)
			}

			if err := r.ParseForm(); err != nil {
				t.Fatalf("failed to parse form: %v", err)
			}

			if r.FormValue("secret") != "secret-key" {
				t.Fatalf(
					"expected secret-key, got %s",
					r.FormValue("secret"),
				)
			}

			if r.FormValue("response") != "valid-token" {
				t.Fatalf(
					"expected valid-token, got %s",
					r.FormValue("response"),
				)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"success": true}`))
		}),
	)
	defer server.Close()

	validator := NewTurnstileValidatorWithURL(
		"secret-key",
		server.URL,
	)

	err := validator.Validate(
		context.Background(),
		"valid-token",
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestTurnstileValidatorReturnsInvalidCaptchaWhenCaptchaIsInvalid(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{
				"success": false,
				"error-codes": ["invalid-input-response"]
			}`))
		}),
	)
	defer server.Close()

	validator := NewTurnstileValidatorWithURL(
		"secret-key",
		server.URL,
	)

	err := validator.Validate(
		context.Background(),
		"invalid-token",
	)

	if !errors.Is(err, ErrInvalidCaptcha) {
		t.Fatalf(
			"expected ErrInvalidCaptcha, got %v",
			err,
		)
	}
}