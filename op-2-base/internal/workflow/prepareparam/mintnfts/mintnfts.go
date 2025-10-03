package mintnfts

import (
	"context"
	"log/slog"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/likecoin/likecoin-op/op-2-base/internal/workflow/prepareactions"
)

type Input struct {
	prepareactions.PrepareMintNFTActionOutput
}

type Output struct {
	Input
	FromTokenId uint64 `json:"from_token_id"`
	To          string `json:"to"`
	Memo        string `json:"memo"`
	Metadata    string `json:"metadata"`
}

type PrepareMintNFTsParam interface {
	Prepare(
		ctx context.Context,
		logger *slog.Logger,
		input *Input,
	) (*Output, error)
}

type prepareMintNFTsParam struct {
}

func NewPrepareMintNFTsParam() PrepareMintNFTsParam {
	return &prepareMintNFTsParam{}
}

func (p *prepareMintNFTsParam) Prepare(
	ctx context.Context,
	logger *slog.Logger,
	input *Input,
) (*Output, error) {
	mylogger := logger.WithGroup("PrepareMintNFTsParam").With("tokenId", input.TokenId)
	mylogger.Info("PrepareMintNFTsParam")

	fromTokenId := input.TokenId
	owner := common.HexToAddress(input.EvmOwner)

	memo := makeMemo(input.Memos)
	metadata := input.MetadataStr

	return &Output{
		Input:       *input,
		FromTokenId: fromTokenId,
		To:          owner.Hex(),
		Memo:        memo,
		Metadata:    metadata,
	}, nil
}

func makeMemo(memos []string) string {
	return strings.Join(memos, "\n\n")
}
