package routes

import (
	"net/http"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/health", handlers.HealthHandler)

	return r
}
