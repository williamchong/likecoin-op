package task

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hibiken/asynq"
	appcontext "github.com/likecoin/like-migration-backend/cmd/worker/context"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
	apptask "github.com/likecoin/like-migration-backend/pkg/task"
)

func HandleMigrateClassTask(ctx context.Context, t *asynq.Task) error {
	envCfg := appcontext.ConfigFromContext(ctx)
	db := appcontext.DBFromContext(ctx)
	logger := appcontext.LoggerFromContext(ctx)

	mylogger := logger.WithGroup("HandleMigrateClassTask")

	var p apptask.MigrateClassPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	mylogger = mylogger.With("LikenftAssetMigrationClassId", p.LikenftAssetMigrationClassId)

	likenftClient := &cosmos.LikeNFTCosmosClient{
		HTTPClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		NodeURL: envCfg.CosmosNodeUrl,
	}

	ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
	if err != nil {
		return err
	}
	privateKey, err := crypto.HexToECDSA(envCfg.EthWalletPrivateKey)
	if err != nil {
		return err
	}
	contractAddress := common.HexToAddress(envCfg.EthLikeNFTContractAddress)

	evmLikeProtocolClient := evm.NewLikeProtocol(
		ethClient,
		privateKey,
		envCfg.EthChainId,
		contractAddress,
	)
	evmLikeNFTClient := evm.NewLikeNFTClass(
		ethClient,
		privateKey,
		envCfg.EthChainId,
	)

	mylogger.Info("running migrate class")
	mc, err := likenft.MigrateClassFromAssetMigration(
		logger,
		db,
		likenftClient,
		&evmLikeProtocolClient,
		&evmLikeNFTClient,
		envCfg.InitialNewClassOwner,
		envCfg.InitialNewClassMinter,
		envCfg.InitialNewClassUpdater,
		p.LikenftAssetMigrationClassId,
	)

	if err != nil {
		mylogger.Error("running migrate class failed", "error", err)
		return err
	}

	mylogger.Info("migrate class completed", "evmTxHash", *mc.EvmTxHash)
	return nil
}
