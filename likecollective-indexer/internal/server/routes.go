package server

import (
	"encoding/json"
	"log"
	"net/http"

	"likecollective-indexer/internal/api/openapi"
	"likecollective-indexer/internal/server/routes/alchemy"
)

func (s *Server) RegisterRoutes() http.Handler {
	openapiHandler := openapi.NewOpenAPIHandler(
		s.db,
		s.timescaleDbService,
	)

	mux := http.NewServeMux()

	// Register routes
	mux.Handle("/api/", http.StripPrefix("/api", openapiHandler))
	if s.alchemyLikeCollectiveEthLogWebhookSigningKey != "" {
		mux.Handle("/alchemy/", http.StripPrefix(
			"/alchemy",
			alchemy.NewAlchemyHandler(
				s.alchemyLikeCollectiveEthLogWebhookSigningKey,
				s.db,
				s.likeCollectiveLogConverter,
			),
		))
	}
	mux.HandleFunc("/", s.HelloWorldHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
