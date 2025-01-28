package api

import (
	"encoding/json"
	"log"
	"net/http"

	"likenft-indexer/ent"
)

func SetupRoutes(r *http.ServeMux, db *ent.Client) {
	// Initialize handlers
	nftHandler := NewNFTHandler(db)
	nftClassHandler := NewNFTClassHandler(db)

	// NFT routes
	r.HandleFunc("GET /api/nfts", nftHandler.GetNFTs)
	r.HandleFunc("GET /api/nft/{id}", nftHandler.GetNFTByID)

	// NFT Class routes
	r.HandleFunc("GET /api/nftclasses", nftClassHandler.GetNFTClasses)
	r.HandleFunc("GET /api/nftclass/{id}", nftClassHandler.GetNFTClassByID)
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
