package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/dto"
	"github.com/flaviolpgjr/aletheia/backend/internal/services"
)

func AnalyzePromiseHandler(w http.ResponseWriter, r *http.Request) {
	var request dto.AnalyzePromiseRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	service := services.NewPromiseAnalyzerService()

	analysis := service.Analyze(request.Text)

	response := dto.AnalyzePromiseResponse{
		Summary: analysis.Summary,
		Score:   analysis.Score,
		Risks:   analysis.Risks,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
