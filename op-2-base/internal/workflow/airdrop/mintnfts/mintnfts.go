package mintnfts

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/likecoin/like-migration-backend/pkg/signer"

	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	preparemintnftsparam "github.com/likecoin/likecoin-op/op-2-base/internal/workflow/prepareparam/mintnfts"
)

type Input struct {
	preparemintnftsparam.Output
}

type Output struct {
	Input
	TxHash string `json:"tx_hash"`
}

type MintNFTsAirdrop interface {
	Airdrop(
		ctx context.Context,
		logger *slog.Logger,
		inputs []*Input,
	) ([]*Output, error)
}

type mintNFTsAirdrop struct {
	baseEthClient *ethclient.Client
	baseSigner    *signer.SignerClient
}

func NewMintNFTsAirdrop(
	baseEthClient *ethclient.Client,
	baseSigner *signer.SignerClient,
) MintNFTsAirdrop {
	return &mintNFTsAirdrop{
		baseEthClient,
		baseSigner,
	}
}

func (m *mintNFTsAirdrop) Airdrop(ctx context.Context, logger *slog.Logger, inputs []*Input) ([]*Output, error) {
	mylogger := logger.WithGroup("MintNFTsAirdrop")
	mylogger.Info("MintNFTsAirdrop")

	baseEvmClassId, err := m.assertSameClassId(inputs)
	if err != nil {
		return nil, fmt.Errorf("class id mismatch: %v", err)
	}

	baseBookNFT, err := evm.NewBookNFT(
		logger,
		m.baseEthClient,
		m.baseSigner,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create base book nft: %v", err)
	}

	fromTokenId, err := m.assertTokenIdSequence(inputs)
	if err != nil {
		return nil, fmt.Errorf("token id sequence mismatch: %v", err)
	}

	tos := make([]common.Address, len(inputs))
	for i, input := range inputs {
		tos[i] = common.HexToAddress(input.To)
	}
	memos := make([]string, len(inputs))
	for i, input := range inputs {
		memos[i] = input.Memo
	}
	metadataList := make([]string, len(inputs))
	for i, input := range inputs {
		metadataList[i] = input.Metadata
	}

	tx, _, err := baseBookNFT.MintNFTs(
		ctx, logger,
		common.HexToAddress(baseEvmClassId),
		fromTokenId,
		tos,
		memos,
		metadataList,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to mint nfts: %v", err)
	}

	mylogger.Info(
		"mint nfts",
		"classId", baseEvmClassId,
		"tx_hash", tx.Hash().Hex(),
	)

	outputs := make([]*Output, len(inputs))
	for i, input := range inputs {
		outputs[i] = &Output{
			Input:  *input,
			TxHash: tx.Hash().Hex(),
		}
	}

	return outputs, nil
}

func (m *mintNFTsAirdrop) assertSameClassId(inputs []*Input) (string, error) {
	classId := inputs[0].EvmClassId
	for _, input := range inputs {
		if input.EvmClassId != classId {
			return "", fmt.Errorf("class id mismatch: %s != %s", input.EvmClassId, classId)
		}
	}
	return classId, nil
}

func (m *mintNFTsAirdrop) assertTokenIdSequence(inputs []*Input) (*big.Int, error) {
	fromTokenId := big.NewInt(0).SetUint64(inputs[0].TokenId)
	for i, input := range inputs {
		tokenId := big.NewInt(0).SetUint64(input.TokenId)
		if tokenId.Cmp(big.NewInt(0).Add(fromTokenId, big.NewInt(int64(i)))) != 0 {
			return nil, fmt.Errorf("token id mismatch: %s != %s", tokenId.String(), big.NewInt(0).Add(fromTokenId, big.NewInt(int64(i))).String())
		}
	}
	return fromTokenId, nil
}
