package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
	App    string `json:"app"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "ok",
		App:    "AletheIA",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}
