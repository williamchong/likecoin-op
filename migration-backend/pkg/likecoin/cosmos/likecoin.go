package cosmos

import (
	"log/slog"

	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos/model"
)

type LikeCoin struct {
	Logger        *slog.Logger
	NetworkConfig *model.NetworkConfig
}

func NewLikeCoin(
	logger *slog.Logger,
	networkConfig *model.NetworkConfig,
) *LikeCoin {
	return &LikeCoin{
		Logger:        logger,
		NetworkConfig: networkConfig,
	}
}
