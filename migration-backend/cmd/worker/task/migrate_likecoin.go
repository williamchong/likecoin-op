package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hibiken/asynq"

	appcontext "github.com/likecoin/like-migration-backend/cmd/worker/context"
	"github.com/likecoin/like-migration-backend/pkg/cosmos/api"
	"github.com/likecoin/like-migration-backend/pkg/ethereum"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos"
	cosmosmodel "github.com/likecoin/like-migration-backend/pkg/likecoin/cosmos/model"
	"github.com/likecoin/like-migration-backend/pkg/likecoin/evm"
	"github.com/likecoin/like-migration-backend/pkg/logic/likecoin"
	"github.com/likecoin/like-migration-backend/pkg/model"
	"github.com/likecoin/like-migration-backend/pkg/signer"
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

	mylogger = mylogger.With(
		"CosmosAddress", p.CosmosAddress,
		"LikeCoinMigrationId", p.LikeCoinMigrationId,
	)

	ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
	if err != nil {
		return err
	}

	cosmosAPI := api.NewCosmosAPI(
		envCfg.CosmosNodeUrl,
		time.Duration(envCfg.CosmosNodeHTTPTimeoutSeconds),
	)

	cosmosLikeCoinNetworkConfigData, err := os.ReadFile(
		envCfg.CosmosLikeCoinNetworkConfigPath,
	)

	if err != nil {
		panic(err)
	}

	cosmosLikeCoinNetworkConfig, err := cosmosmodel.LoadNetworkConfig(
		cosmosLikeCoinNetworkConfigData,
	)

	if err != nil {
		panic(err)
	}

	cosmosLikeCoinClient := cosmos.NewLikeCoin(
		logger,
		cosmosLikeCoinNetworkConfig,
	)

	signer := signer.NewSignerClient(
		&http.Client{
			Timeout: 10 * time.Second,
		},
		envCfg.EthSignerBaseUrl,
		envCfg.EthSignerAPIKey,
	)

	ethereumClient := ethereum.NewClient(logger, ethClient, signer)

	contractAddress := common.HexToAddress(envCfg.EthTokenAddress)
	likeCoinClient, err := evm.NewLikeCoin(
		logger,
		ethClient,
		signer,
		contractAddress,
	)
	if err != nil {
		return err
	}

	mylogger.Info("running migrate likecoin")

	var migration *model.LikeCoinMigration
	if p.CosmosAddress != "" {
		migration, err = likecoin.DoMintLikeCoinByCosmosAddress(
			ctx,
			mylogger,
			db,
			ethClient,
			cosmosAPI,
			likeCoinClient,
			cosmosLikeCoinClient,
			p.CosmosAddress,
		)
	} else if p.LikeCoinMigrationId != 0 {
		migration, err = likecoin.DoMintLikeCoinByLikeCoinMigrationId(
			ctx,
			mylogger,
			db,
			ethClient,
			cosmosAPI,
			likeCoinClient,
			cosmosLikeCoinClient,
			p.LikeCoinMigrationId,
		)
	} else {
		return fmt.Errorf("cosmos address or likecoin migration id is required")
	}

	if err != nil {
		mylogger.Error("running migrate likecoin failed", "error", err)
		return err
	}

	err = airdropEth(
		ctx,
		mylogger,
		ethereumClient,
		common.HexToAddress(migration.UserEthAddress),
		envCfg.EthAmountToAirdropOnMigrate,
		migration.EvmTxHash,
	)
	if err != nil {
		mylogger.Error("airdrop eth failed", "error", err)
		// The error is ignored
	}

	mylogger.Info("migrate likecoin completed", "evmTxHash", *migration.EvmTxHash)
	return nil
}

func airdropEth(
	ctx context.Context,
	logger *slog.Logger,
	ethereumClient ethereum.EthereumClient,
	toAddress common.Address,
	ethAmount string,
	likecoinMigrationTx *string,
) error {
	amountEth, ok := new(big.Float).SetString(ethAmount)
	if !ok {
		logger.Error("new(big.Float).SetString", "error", fmt.Errorf("invalid amount: %s", ethAmount))
		return fmt.Errorf("invalid amount: %s", ethAmount)
	}
	amountWei := ethereum.ConvertToWei(amountEth)
	if amountWei.Cmp(big.NewInt(0)) == 0 {
		logger.Info("skipped airdrop because amount is zero")
		return nil
	}
	_, _, err := likecoin.DoAirdropEth(ctx, logger, ethereumClient, toAddress, amountWei, likecoinMigrationTx)
	if err != nil {
		logger.Error("likecoin.DoAirdropEth", "error", err)
		return err
	}
	return nil
}
