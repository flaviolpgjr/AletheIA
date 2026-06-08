package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/dto"
	"github.com/flaviolpgjr/aletheia/backend/internal/services"
)

type AnalyzePromiseHandler struct {
	service *services.PromiseAnalyzerService
}

func NewAnalyzePromiseHandler(service *services.PromiseAnalyzerService) *AnalyzePromiseHandler {
	return &AnalyzePromiseHandler{
		service: service,
	}
}

func (h *AnalyzePromiseHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var request dto.AnalyzePromiseRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	analysis := h.service.Analyze(r.Context(), request.Text)

	criteria := make([]dto.CriterionResponse, 0, len(analysis.Criteria))

	for _, criterion := range analysis.Criteria {
		criteria = append(criteria, dto.CriterionResponse{
			Key:         criterion.Key,
			Name:        criterion.Name,
			Weight:      criterion.Weight,
			Status:      string(criterion.Status),
			Score:       criterion.Score,
			Explanation: criterion.Explanation,
		})
	}

	response := dto.AnalyzePromiseResponse{
		Summary:    analysis.Summary,
		Score:      analysis.Score,
		Confidence: analysis.Confidence,
		Criteria:   criteria,
		Risks:      analysis.Risks,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}