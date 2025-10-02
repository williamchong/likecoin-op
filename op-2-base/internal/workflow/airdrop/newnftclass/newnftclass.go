package newnftclass

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm/like_protocol"
	preparenewnftclassparam "github.com/likecoin/likecoin-op/op-2-base/internal/workflow/prepareparam/newnftclass"
)

type Input struct {
	preparenewnftclassparam.Output
}

type Output struct {
	Input
	BaseEvmClassId string `json:"base_evm_class_id"`
	TxHash         string `json:"tx_hash"`
}

type NewNFTClassAirdrop interface {
	Airdrop(
		ctx context.Context,
		logger *slog.Logger,
		input *Input,
	) (*Output, error)
}

type newNFTClassAirdrop struct {
	baseLikeProtocol *evm.LikeProtocol
}

func NewNewNFTClassAirdrop(
	baseLikeProtocol *evm.LikeProtocol,
) NewNFTClassAirdrop {
	return &newNFTClassAirdrop{
		baseLikeProtocol,
	}
}

func (n *newNFTClassAirdrop) Airdrop(
	ctx context.Context,
	logger *slog.Logger,
	input *Input,
) (*Output, error) {
	mylogger := logger.WithGroup("NewNFTClassAirdrop")
	mylogger.Info("NewNFTClassAirdrop")

	signerAddressStr, err := n.baseLikeProtocol.Signer.GetSignerAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to get signer address: %v", err)
	}
	signerAddress := common.HexToAddress(*signerAddressStr)

	salt, err := evm.ComputeNewBookNFTSalt(signerAddress, [2]byte{0, 0}, input.Metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to compute new book nft salt: %v", err)
	}

	updaters := make([]common.Address, len(input.Updaters))
	for i, updater := range input.Updaters {
		updaters[i] = common.HexToAddress(updater)
	}

	minters := make([]common.Address, len(input.Minters))
	for i, minter := range input.Minters {
		minters[i] = common.HexToAddress(minter)
	}

	maxSupply := input.Config.MaxSupply
	if err != nil {
		return nil, fmt.Errorf("failed to parse max supply: %v", err)
	}

	defaultRoyaltyFraction, ok := big.NewInt(0).SetString(input.DefaultRoyaltyFraction, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse default royalty fraction: %s", input.DefaultRoyaltyFraction)
	}

	tx, txReceipt, err := n.baseLikeProtocol.NewBookNFTWithRoyaltyAndSalt(
		ctx, logger,
		salt,
		like_protocol.MsgNewBookNFT{
			Creator:  common.HexToAddress(input.Creator),
			Updaters: updaters,
			Minters:  minters,
			Config: like_protocol.BookConfig{
				Name:      input.Config.Name,
				Symbol:    input.Config.Symbol,
				Metadata:  input.Config.Metadata,
				MaxSupply: maxSupply,
			},
		},
		defaultRoyaltyFraction,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to new book nft with royalty and salt: %v", err)
	}

	newClassId, err := n.baseLikeProtocol.GetClassIdFromNewClassTransaction(txReceipt)
	if err != nil {
		return nil, fmt.Errorf("failed to get class id from new class transaction: %v", err)
	}

	mylogger.Info("new book nft with royalty and salt",
		"classId", newClassId.Hex(),
		"tx_hash", tx.Hash().Hex(),
	)

	return &Output{
		Input:          *input,
		BaseEvmClassId: newClassId.Hex(),
		TxHash:         tx.Hash().Hex(),
	}, nil
}
