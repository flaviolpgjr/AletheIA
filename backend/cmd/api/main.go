package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/routes"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
	"github.com/flaviolpgjr/aletheia/backend/internal/services"
)

func main() {
	geminiAPIKey := os.Getenv("GEMINI_API_KEY")

	llmClient := llm.NewGeminiClient(geminiAPIKey)

	analyzerService := services.NewPromiseAnalyzerService(llmClient)

	router := routes.NewRouter(analyzerService)

	log.Println("API running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}