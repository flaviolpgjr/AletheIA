package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/flaviolpgjr/aletheia/backend/internal/http/dto"
	"github.com/flaviolpgjr/aletheia/backend/internal/llm"
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

	analysis, err := h.service.Analyze(r.Context(), request.Text)
	if err != nil {
		if errors.Is(err, llm.ErrRateLimit) {
			http.Error(
				w,
				"O provedor de IA atingiu o limite temporário. Tente novamente em alguns segundos.",
				http.StatusTooManyRequests,
			)
			return
		}

		http.Error(
			w,
			"failed to analyze promise",
			http.StatusInternalServerError,
		)
		return
	}

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

	sources := make([]dto.PublicSourceResponse, 0, len(analysis.Sources))

	for _, source := range analysis.Sources {
		sources = append(sources, dto.PublicSourceResponse{
			Name:        source.Name,
			Description: source.Description,
		})
	}

	response := dto.AnalyzePromiseResponse{
		Summary:    analysis.Summary,
		Score:      analysis.Score,
		Confidence: analysis.Confidence,
		Criteria:   criteria,
		Risks:      analysis.Risks,
		Sources:    sources,
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}