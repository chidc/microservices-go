package main

import (
	"encoding/json"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/types"
)

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// validation
	if reqBody.UserID == "" {
		http.Error(w, "Missing userID", http.StatusBadRequest)
		return
	}

	response := contracts.APIResponse{Data: "OK"}
	// TODO: call trip service
	writeJSON(w, http.StatusCreated, response)
}
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
