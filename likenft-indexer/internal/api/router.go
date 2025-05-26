package api

import (
	"encoding/json"
	"log"
	"net/http"

	"likenft-indexer/internal/api/openapi"
	"likenft-indexer/internal/database"

	"github.com/hibiken/asynq"
)

func SetupRoutes(
	r *http.ServeMux,

	indexActionApiKey string,
	likeProtocolAddress string,

	db database.Service,

	asynqClient *asynq.Client,
) {
	// Initialize handlers
	openapiHandler := openapi.NewOpenAPIHandler(
		indexActionApiKey,
		likeProtocolAddress,
		db,
		asynqClient,
	)

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
