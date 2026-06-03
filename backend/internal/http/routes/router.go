package routes

import (
	"net/http"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/handlers"
	"github.com/flaviolpgjr/aletheia/backend/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(analyzerService *services.PromiseAnalyzerService) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
	}))

	analyzePromiseHandler := handlers.NewAnalyzePromiseHandler(analyzerService)

	r.Get("/health", handlers.HealthHandler)
	r.Post("/promises/analyze", analyzePromiseHandler.Handle)

	return r
}