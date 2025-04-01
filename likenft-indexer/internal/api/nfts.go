package api

import (
	"context"
	"net/http"
	"strconv"

	"likenft-indexer/ent"
)

type NFTHandler struct {
	// Add your dependencies here (e.g., database service)
	client *ent.Client
}

func NewNFTHandler(client *ent.Client) *NFTHandler {
	return &NFTHandler{
		client: client,
	}
}

func (h *NFTHandler) GetNFTs(w http.ResponseWriter, r *http.Request) {
	nfts, err := h.client.NFT.Query().
		Limit(10).
		All(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, nfts)
}

type NFTRequestParams struct {
	ID int
}

func NewNFTRequestParams(r *http.Request) (*NFTRequestParams, error) {
	var req NFTRequestParams
	id := r.PathValue("id")
	_id, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	req = NFTRequestParams{
		ID: _id,
	}

	return &req, nil
}

func (h *NFTHandler) GetNFTByID(w http.ResponseWriter, r *http.Request) {
	params, err := NewNFTRequestParams(r)
	if err != nil {
		http.Error(w, "Invalid request params", http.StatusBadRequest)
		return
	}

	nft, err := h.client.NFT.Get(context.Background(), params.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sendJSON(w, nft)
}
