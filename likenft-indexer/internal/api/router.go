package api

import (
	"encoding/json"
	"log"
	"net/http"

	"likenft-indexer/ent"
	"likenft-indexer/internal/api/openapi"
)

func SetupRoutes(r *http.ServeMux, db *ent.Client) {
	// Initialize handlers
	openapiHandler := openapi.NewOpenAPIHandler(db)

	r.Handle("/api/", http.StripPrefix("/api", openapiHandler))
}

func sendJSON(w http.ResponseWriter, data interface{}) {
	jsonResp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
