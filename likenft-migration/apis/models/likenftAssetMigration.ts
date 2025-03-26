import { z } from 'zod';

export const LikeNFTAssetMigrationClassStatusSchema = z.enum([
  'init',
  'in_progress',
  'completed',
  'failed',
]);

export const LikeNFTAssetMigrationClassSchema = z.object({
  id: z.number(),
  nft_asset_migration_id: z.number(),
  created_at: z.coerce.date(),
  cosmos_class_id: z.string(),
  name: z.string(),
  image: z.string(),
  status: LikeNFTAssetMigrationClassStatusSchema,
  enqueue_time: z.coerce.date().nullable(),
  finish_time: z.coerce.date().nullable(),
  evm_tx_hash: z.string().nullable(),
  failed_reason: z.string().nullable(),
});

export type LikeNFTAssetMigrationClass = z.infer<
  typeof LikeNFTAssetMigrationClassSchema
>;

export const LikeNFTAssetMigrationNFTStatusSchema = z.enum([
  'init',
  'in_progress',
  'completed',
  'failed',
]);

export const LikeNFTAssetMigrationNFTSchema = z.object({
  id: z.number(),
  nft_asset_migration_id: z.number(),
  created_at: z.coerce.date(),
  cosmos_class_id: z.string(),
  cosmos_nft_id: z.string(),
  name: z.string(),
  image: z.string(),
  status: LikeNFTAssetMigrationNFTStatusSchema,
  enqueue_time: z.coerce.date().nullable(),
  finish_time: z.coerce.date().nullable(),
  evm_tx_hash: z.string().nullable(),
  failed_reason: z.string().nullable(),
});

export type LikeNFTAssetMigrationNFT = z.infer<
  typeof LikeNFTAssetMigrationNFTSchema
>;

export const LikeNFTAssetMigrationStatusSchema = z.enum([
  'init',
  'in_progress',
  'completed',
  'failed',
]);

export const LikeNFTAssetMigrationSchema = z.object({
  id: z.number(),
  created_at: z.coerce.date(),
  likenft_asset_snapshot_id: z.number(),
  cosmos_address: z.string(),
  eth_address: z.string(),
  status: LikeNFTAssetMigrationStatusSchema,
  failed_reason: z.string().nullable(),
  classes: z.array(LikeNFTAssetMigrationClassSchema),
  nfts: z.array(LikeNFTAssetMigrationNFTSchema),
});

export type LikeNFTAssetMigration = z.infer<typeof LikeNFTAssetMigrationSchema>;

export type CompletedLikeNFTAssetMigration = LikeNFTAssetMigration & {
  status: 'completed';
};

export function isMigrationCompleted(
  m: LikeNFTAssetMigration
): m is CompletedLikeNFTAssetMigration {
  return m.status === 'completed';
}

export type FailedLikeNFTAssetMigration = LikeNFTAssetMigration & {
  status: 'failed';
};

export function isMigrationFailed(
  m: LikeNFTAssetMigration
): m is FailedLikeNFTAssetMigration {
  return m.status === 'failed';
}
