package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/flaviolpgjr/aletheia/backend/internal/database"
	"github.com/flaviolpgjr/aletheia/backend/internal/http/routes"
	"github.com/flaviolpgjr/aletheia/backend/internal/llmfactory"
	"github.com/flaviolpgjr/aletheia/backend/internal/publicdata"
	"github.com/flaviolpgjr/aletheia/backend/internal/publicdata/health"
	"github.com/flaviolpgjr/aletheia/backend/internal/repositories"
	"github.com/flaviolpgjr/aletheia/backend/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	ctx := context.Background()

	dbPool, err := database.NewPostgresPool(
		ctx,
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer dbPool.Close()

	analysisRepository := repositories.NewAnalysisRepository(dbPool)
	publicDataBaselineRepository := repositories.NewPublicDataBaselineRepository(dbPool)

	llmClient, err := llmfactory.NewClient(llmfactory.Config{
		Provider: os.Getenv("LLM_PROVIDER"),

		GeminiAPIKey: os.Getenv("GEMINI_API_KEY"),

		OpenRouterAPIKey: os.Getenv("OPENROUTER_API_KEY"),
		OpenRouterModel:  os.Getenv("OPENROUTER_MODEL"),
	})
	if err != nil {
		log.Fatal(err)
	}

	healthClient := health.NewClient(publicDataBaselineRepository)

	publicDataAggregator := publicdata.NewAggregator(
		healthClient,
	)
	analyzerService := services.NewPromiseAnalyzerService(
		llmClient,
		analysisRepository,
		publicDataAggregator,
	)
	router := routes.NewRouter(analyzerService)

	log.Println("API running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}