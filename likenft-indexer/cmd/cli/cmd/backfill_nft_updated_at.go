package cmd

import (
	"likenft-indexer/cmd/cli/context"
	"likenft-indexer/internal/database"

	"github.com/spf13/cobra"
)

var BackfillNFTUpdatedAtCmd = &cobra.Command{
	Use:   "backfill-nft-updated-at",
	Short: "Backfill nfts.updated_at from Transfer event block timestamps",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		logger := context.LoggerFromContext(ctx)
		mylogger := logger.WithGroup("BackfillNFTUpdatedAt")

		dbService := database.New()
		nftRepository := database.MakeNFTRepository(dbService)

		updatedCount, err := nftRepository.BackfillUpdatedAtFromTransferEvents(ctx, mylogger)

		if err != nil {
			panic(err)
		}

		mylogger.Info("done", "updated", updatedCount)
	},
}
