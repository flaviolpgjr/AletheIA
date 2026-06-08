package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/routes"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
	"github.com/flaviolpgjr/aletheia/backend/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	geminiAPIKey := os.Getenv("GEMINI_API_KEY")

	log.Printf("GEMINI_API_KEY loaded: %v", geminiAPIKey != "")

	llmClient := llm.NewGeminiClient(geminiAPIKey)
	analyzerService := services.NewPromiseAnalyzerService(llmClient)

	router := routes.NewRouter(analyzerService)

	log.Println("API running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}