package stakingstate

import (
	"likecollective-indexer/internal/evm/like_collective"
	"likecollective-indexer/internal/evm/util/logconverter"
)

var (
	logConverter *logconverter.LogConverter
)

func init() {
	abi, err := like_collective.LikeCollectiveMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	logConverter = logconverter.NewLogConverter(abi)
}
