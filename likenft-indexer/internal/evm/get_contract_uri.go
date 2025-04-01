package evm

import (
	"context"

	"likenft-indexer/internal/evm/book_nft"
	"likenft-indexer/internal/util/jsondatauri"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (c *EvmClient) GetContractURI(
	ctx context.Context,
	contractAddress common.Address,
) (jsondatauri.JSONDataUri, error) {
	bookNFTClient, err := book_nft.NewBookNft(contractAddress, c.client)
	if err != nil {
		return "", err
	}
	contractURI, err := bookNFTClient.ContractURI(&bind.CallOpts{})
	if err != nil {
		return "", err
	}
	return jsondatauri.JSONDataUri(contractURI), nil
}
