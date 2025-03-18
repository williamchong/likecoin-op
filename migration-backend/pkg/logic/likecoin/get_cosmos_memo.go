package likecoin

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/types"
)

type memoJsonData struct {
	Signature  string `json:"signature"`
	EthAddress string `json:"ethAddress"`
	Amount     string `json:"amount"`
}

type MemoData struct {
	Signature  string
	EthAddress string
	Amount     types.Coin
}

func ParseMemoData(memoStr string) (*MemoData, error) {
	var memoJsonData memoJsonData
	err := json.Unmarshal([]byte(memoStr), &memoJsonData)
	if err != nil {
		return nil, err
	}

	amount, err := types.ParseCoinNormalized(memoJsonData.Amount)
	if err != nil {
		return nil, err
	}

	return &MemoData{
		Signature:  memoJsonData.Signature,
		EthAddress: memoJsonData.EthAddress,
		Amount:     amount,
	}, nil
}

func EncodeCosmosMemoData(m *MemoData) (string, error) {
	memoJsonData := &memoJsonData{
		Signature:  m.Signature,
		EthAddress: m.EthAddress,
		Amount:     m.Amount.String(),
	}
	memoJsonStr, err := json.Marshal(memoJsonData)
	if err != nil {
		return "", err
	}
	return string(memoJsonStr), nil
}
