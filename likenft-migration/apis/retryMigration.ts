import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { LikeNFTAssetMigrationSchema } from './models/likenftAssetMigration';

export const BookNFTSchema = z.object({
  class_id: z.string(),
  nft_id: z.string(),
});

export const RetryMigrationRequestSchema = z.object({
  book_nft_collection: z.array(z.string()),
  book_nft: z.array(BookNFTSchema),
});

export type RetryMigrationRequest = z.infer<typeof RetryMigrationRequestSchema>;

export const RetryMigrationResponseSchema = z.object({
  migration: LikeNFTAssetMigrationSchema,
});

export type RetryMigrationResponse = z.infer<
  typeof RetryMigrationResponseSchema
>;

export const makeRetryMigrationAPI = (cosmosAddress: string) =>
  makeAPI({
    url: `/likenft/migration/${cosmosAddress}`,
    method: 'PUT',
    requestSchema: RetryMigrationRequestSchema,
    responseSchema: RetryMigrationResponseSchema,
  });
