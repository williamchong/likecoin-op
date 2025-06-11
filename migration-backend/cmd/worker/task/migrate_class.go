package task

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hibiken/asynq"

	appcontext "github.com/likecoin/like-migration-backend/cmd/worker/context"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/likenft/cosmos"
	"github.com/likecoin/like-migration-backend/pkg/likenft/evm"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/cosmosnftidclassifier"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/erc721externalurl"
	"github.com/likecoin/like-migration-backend/pkg/likenft/util/nftidmatcher"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
	"github.com/likecoin/like-migration-backend/pkg/signer"
	apptask "github.com/likecoin/like-migration-backend/pkg/task"
)

func HandleMigrateClassTask(ctx context.Context, t *asynq.Task) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	envCfg := appcontext.ConfigFromContext(ctx)
	db := appcontext.DBFromContext(ctx)
	logger := appcontext.LoggerFromContext(ctx)

	mylogger := logger.WithGroup("HandleMigrateClassTask")

	cosmosNFTIdClassifier := cosmosnftidclassifier.MakeCosmosNFTIDClassifier(
		nftidmatcher.MakeNFTIDMatcher(),
	)

	erc721ExternalURLBuilder, err := erc721externalurl.MakeErc721ExternalURLBuilder3ook(
		envCfg.ERC721MetadataExternalURLBase3ook,
	)

	if err != nil {
		panic(err)
	}

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

	likecoinAPI := likecoin_api.NewLikecoinAPI(
		envCfg.LikecoinAPIUrlBase,
		time.Duration(envCfg.LikecoinAPIHTTPTimeoutSeconds),
	)

	ethClient, err := ethclient.Dial(envCfg.EthNetworkPublicRPCURL)
	if err != nil {
		return err
	}

	signer := signer.NewSignerClient(
		&http.Client{
			Timeout: 10 * time.Second,
		},
		envCfg.EthSignerBaseUrl,
		envCfg.EthSignerAPIKey,
	)

	contractAddress := common.HexToAddress(envCfg.EthLikeNFTContractAddress)

	evmLikeProtocolClient := evm.NewLikeProtocol(
		logger,
		ethClient,
		signer,
		contractAddress,
	)
	evmLikeNFTClient, err := evm.NewBookNFT(
		logger,
		ethClient,
		signer,
	)

	if err != nil {
		return err
	}

	mylogger.Info("running migrate class")
	mc, err := likenft.MigrateClassFromAssetMigration(
		ctx,
		logger,
		db,
		likenftClient,
		likecoinAPI,
		&evmLikeProtocolClient,
		&evmLikeNFTClient,
		cosmosNFTIdClassifier,
		erc721ExternalURLBuilder,
		envCfg.ShouldPremintAllNFTsWhenNewClass,
		envCfg.PremintAllNFTsWhenNewClassShouldPremintArbitraryNFTIDs,
		envCfg.InitialNewClassOwner,
		envCfg.InitialNewClassMinters,
		envCfg.InitialNewClassUpdater,
		envCfg.InitialBatchMintNFTsOwner,
		envCfg.BatchMintItemPerPage,
		new(big.Int).SetUint64(envCfg.DefaultRoyaltyFraction),
		p.LikenftAssetMigrationClassId,
	)

	if err != nil {
		mylogger.Error("running migrate class failed", "error", err)
		return err
	}

	mylogger.Info("migrate class completed", "evmTxHash", *mc.EvmTxHash)
	return nil
}
