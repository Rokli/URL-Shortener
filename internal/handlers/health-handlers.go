package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status  string
	Service string
}

func HealthCheckAll(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:  "healthy",
		Service: "server",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JSONResponse{
		Success: true,
		Data:    response,
	})
}
