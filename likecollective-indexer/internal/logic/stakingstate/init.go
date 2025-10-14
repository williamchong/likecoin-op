package stakingstate

import (
	"likecollective-indexer/internal/evm/like_collective"
	"likecollective-indexer/internal/evm/like_stake_position"
	"likecollective-indexer/internal/evm/util/logconverter"
)

var (
	likeCollectiveLogConverter    *logconverter.LogConverter
	likeStakePositionLogConverter *logconverter.LogConverter
)

func init() {
	likeCollectiveAbi, err := like_collective.LikeCollectiveMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	likeCollectiveLogConverter = logconverter.NewLogConverter(likeCollectiveAbi)

	likeStakePositionAbi, err := like_stake_position.LikeStakePositionMetaData.GetAbi()
	if err != nil {
		panic(err)
	}
	likeStakePositionLogConverter = logconverter.NewLogConverter(likeStakePositionAbi)
}
