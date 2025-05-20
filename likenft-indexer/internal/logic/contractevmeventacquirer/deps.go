package contractevmeventacquirer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ABIManager interface {
	GetLikeProtocolABI() *abi.ABI
	GetBookNFTABI() *abi.ABI
}

type EvmEventQueryClient interface {
	QueryEvents(
		ctx context.Context,
		contractAddresses []common.Address,
		startBlock uint64,
		endBlock uint64,
	) ([]types.Log, error)
}

type EvmClient interface {
	GetHeaderMapByBlockNumbers(
		ctx context.Context,
		blockNumbers []uint64,
	) (map[uint64]*types.Header, error)
	BlockNumber(
		ctx context.Context,
	) (uint64, error)
	ChainID(
		ctx context.Context,
	) (*big.Int, error)
}
