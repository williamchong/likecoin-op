package likenft

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/likecoin/like-migration-backend/cmd/cli/config"
	likecoin_api "github.com/likecoin/like-migration-backend/pkg/likecoin/api"
	"github.com/likecoin/like-migration-backend/pkg/logic/likenft"
)

var submitEvmBookMigrated = &cobra.Command{
	Use:   "submit-evm-book-migrated like-class-id evm-class-id",
	Short: "Submit EVM Book Migrated",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			_ = cmd.Usage()
			return
		}

		ctx := cmd.Context()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		likeClassId := args[0]
		evmClassId := args[1]

		envCfg := cmd.Context().Value(config.ContextKey).(*config.EnvConfig)
		db, err := sql.Open("postgres", envCfg.DbConnectionStr)
		if err != nil {
			panic(err)
		}

		logger := slog.New(slog.Default().Handler()).
			WithGroup("submitEvmBookMigrated").
			With("likeClassId", likeClassId).
			With("evmClassId", evmClassId)

		likecoinAPI := likecoin_api.NewLikecoinAPI(envCfg.LikecoinAPIUrlBase)

		response, err := likenft.SubmitEvmBookMigrated(ctx, logger, db, likecoinAPI, likeClassId, evmClassId)
		if err != nil {
			panic(err)
		}

		fmt.Printf("response: %+v\n", response)
	},
}

func init() {
	LikeNFTCmd.AddCommand(submitEvmBookMigrated)
}
