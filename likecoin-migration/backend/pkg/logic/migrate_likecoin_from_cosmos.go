package logic

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/likecoin/likecoin-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/likecoin-migration-backend/pkg/ethereum"
)

type Memo struct {
	Signature  string `json:"signature"`
	EthAddress string `json:"ethAddress"`
	Amount     string `json:"amount"`
	Denom      string `json:"denom"`
}

func getMessage(amount string, denom string) string {
	return fmt.Sprintf(`You are going to deposit %v %v to migration program.

This sign make sure the address is correct.`, amount, denom)
}

func MigrateLikeCoinFromCosmos(
	cosmosAPI *api.CosmosAPI,
	ethNetworkPublicRPCURL string,
	ethWalletPrivateKey string,
	ethTokenAddress string,
	cosmosTxHash string,
) (*types.Transaction, error) {
	txResponse, err := cosmosAPI.QueryTransaction(cosmosTxHash)
	if err != nil {
		return nil, err
	}

	// TODO: Compare the cosmos wallet address which receives token
	// should be the one desire

	memoString := txResponse.Tx.Body.Memo
	var memo Memo
	err = json.Unmarshal([]byte(memoString), &memo)
	if err != nil {
		return nil, err
	}
	message := getMessage(memo.Amount, memo.Denom)

	recoveredAddr, err := ethereum.RecoverAddress(memo.Signature, []byte(message))
	if err != nil {
		return nil, err
	}

	// TODO: Use worker
	tx, err := ethereum.TransferToken(
		ethNetworkPublicRPCURL,
		ethWalletPrivateKey,
		*recoveredAddr,
		common.HexToAddress(ethTokenAddress),
		memo.Amount)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
