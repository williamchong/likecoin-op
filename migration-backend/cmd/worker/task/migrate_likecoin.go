package task

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hibiken/asynq"

	appcontext "github.com/likecoin/like-migration-backend/cmd/worker/context"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
	apptask "github.com/likecoin/like-migration-backend/pkg/task"
)

func HandleMigrateLikeCoinTask(ctx context.Context, t *asynq.Task) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	envCfg := appcontext.ConfigFromContext(ctx)
	db := appcontext.DBFromContext(ctx)
	logger := appcontext.LoggerFromContext(ctx)

	mylogger := logger.WithGroup("HandleMigrateLikeCoinTask")

	var p apptask.MigrateLikeCoinPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	mylogger = mylogger.With("CosmosAddress", p.CosmosAddress)

	ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
	if err != nil {
		return err
	}

	cosmosAPI := api.NewCosmosAPI(envCfg.CosmosNodeUrl)

	cosmosLikeCoinNetworkConfigData, err := os.ReadFile(
		envCfg.CosmosLikeCoinNetworkConfigPath,
	)

	if err != nil {
		panic(err)
	}

	cosmosLikeCoinNetworkConfig, err := model.LoadNetworkConfig(
		cosmosLikeCoinNetworkConfigData,
	)

	if err != nil {
		panic(err)
	}

	cosmosLikeCoinClient := cosmos.NewLikeCoin(
		logger,
		cosmosLikeCoinNetworkConfig,
	)

	contractAddress := common.HexToAddress(envCfg.EthTokenAddress)
	likeCoinClient := evm.NewLikeCoin(
		logger,
		ethClient,
		envCfg.EthChainId,
		contractAddress,
	)
	authedLikeCoinClient := likeCoinClient.Auth(envCfg.EthWalletPrivateKey)

	mylogger.Info("running migrate likecoin")

	migration, err := likecoin.DoMintLikeCoinByCosmosAddress(
		ctx,
		mylogger,
		db,
		ethClient,
		cosmosAPI,
		authedLikeCoinClient,
		cosmosLikeCoinClient,
		p.CosmosAddress,
	)

	if err != nil {
		mylogger.Error("running migrate likecoin failed", "error", err)
		return err
	}

	mylogger.Info("migrate likecoin completed", "evmTxHash", *migration.EvmTxHash)
	return nil
}
