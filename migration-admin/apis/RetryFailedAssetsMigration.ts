import { z } from "zod";

import { makeAPI } from "./makeAPI";
import { LikeNFTMigrationDetailSchema } from "./models/likenftMigration";

export const BookNFTSchema = z.object({
  class_id: z.string(),
  nft_id: z.string(),
});

export const RetryFailedAssetsMigrationRequestSchema = z.object({
  book_nft_collection: z.array(z.string()),
  book_nft: z.array(BookNFTSchema),
});

export type RetryFailedAssetsMigrationRequest = z.infer<
  typeof RetryFailedAssetsMigrationRequestSchema
>;

export const RetryFailedAssetsMigrationResponseSchema = z.object({
  migration: LikeNFTMigrationDetailSchema,
});

export type RetryFailedAssetsMigrationResponse = z.infer<
  typeof RetryFailedAssetsMigrationResponseSchema
>;

export const makeRetryFailedAssetsMigrationAPI = (cosmosAddress: string) =>
  makeAPI({
    url: `/likenft/migration/${cosmosAddress}`,
    method: "PUT",
    requestSchema: RetryFailedAssetsMigrationRequestSchema,
    responseSchema: RetryFailedAssetsMigrationResponseSchema,
  });
