package db

import (
	"time"

	"github.com/likecoin/like-migration-backend/pkg/model"
)

func QueryNFTSigningMessageByCosmosAddress(
	tx TxLike,
	cosmosAddress string,
) (*model.NFTSigningMessage, error) {
	row := tx.QueryRow(
		`SELECT
	id,
	created_at,
	cosmos_address,
	liker_id,
	eth_address,
	nonce,
	issue_time,
	message
FROM nft_signing_message WHERE cosmos_address = $1`, cosmosAddress)

	nftSigningMessage := &model.NFTSigningMessage{}

	err := row.Scan(
		&nftSigningMessage.Id,
		&nftSigningMessage.CreatedAt,
		&nftSigningMessage.CosmosAddress,
		&nftSigningMessage.LikerID,
		&nftSigningMessage.EthAddress,
		&nftSigningMessage.Nonce,
		&nftSigningMessage.IssueTime,
		&nftSigningMessage.Message,
	)

	if err != nil {
		return nil, err
	}

	return nftSigningMessage, nil
}

func InsertNFTSigningMessage(
	tx TxLike,
	nftSigningMessage *model.NFTSigningMessage,
) error {
	_, err := tx.Exec(
		"INSERT INTO nft_signing_message (cosmos_address, liker_id, eth_address, nonce, issue_time, message) VALUES ($1, $2, $3, $4, $5, $6)",
		nftSigningMessage.CosmosAddress,
		nftSigningMessage.LikerID,
		nftSigningMessage.EthAddress,
		nftSigningMessage.Nonce,
		nftSigningMessage.IssueTime.UTC().Format(time.RFC3339),
		nftSigningMessage.Message,
	)

	if err != nil {
		return err
	}

	return nil
}
