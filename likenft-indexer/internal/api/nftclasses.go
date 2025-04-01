package api

import (
	"context"
	"net/http"
	"strconv"

	"likenft-indexer/ent"
)

type NFTClassHandler struct {
	db *ent.Client
}

func NewNFTClassHandler(db *ent.Client) *NFTClassHandler {
	return &NFTClassHandler{
		db: db,
	}
}

func (h *NFTClassHandler) GetNFTClasses(w http.ResponseWriter, r *http.Request) {
	classes, err := h.db.NFTClass.Query().All(context.Background())
	if err != nil {
		http.Error(w, "Failed to fetch NFT classes", http.StatusInternalServerError)
		return
	}

	sendJSON(w, classes)
}

type NFTClassRequestParams struct {
	ID int
}

func NewNFTClassRequestParams(r *http.Request) (*NFTClassRequestParams, error) {
	var req NFTClassRequestParams
	id := r.PathValue("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	req = NFTClassRequestParams{
		ID: _id,
	}
	return &req, nil
}

func (h *NFTClassHandler) GetNFTClassByID(w http.ResponseWriter, r *http.Request) {
	params, err := NewNFTClassRequestParams(r)
	if err != nil {
		http.Error(w, "Invalid request params", http.StatusBadRequest)
		return
	}

	class, err := h.db.NFTClass.Get(context.Background(), params.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			http.Error(w, "NFT class not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to fetch NFT class", http.StatusInternalServerError)
		return
	}

	sendJSON(w, class)
}
