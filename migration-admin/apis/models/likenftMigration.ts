import { makeImageUrl } from "~/utils/imageUrl";
import { z } from "zod";

export enum LikeNFTMigrationStatus {
  Init = "init",
  InProgress = "in_progress",
  Completed = "completed",
  Failed = "failed",
}

export const LikeNFTAssetMigrationClassSchema = z.object({
  id: z.number(),
  nft_asset_migration_id: z.number(),
  created_at: z.coerce.date(),
  cosmos_class_id: z.string(),
  name: z.string(),
  image: z.string().transform((d) => makeImageUrl(d)),
  status: z.nativeEnum(LikeNFTMigrationStatus),
  enqueue_time: z.coerce.date().nullable(),
  finish_time: z.coerce.date().nullable(),
  evm_tx_hash: z.string().nullable(),
  failed_reason: z.string().nullable(),
});

export const LikeNFTAssetMigrationNFTSchema = z.object({
  id: z.number(),
  nft_asset_migration_id: z.number(),
  created_at: z.coerce.date(),
  cosmos_class_id: z.string(),
  cosmos_nft_id: z.string(),
  name: z.string(),
  image: z.string().transform((d) => makeImageUrl(d)),
  status: z.nativeEnum(LikeNFTMigrationStatus),
  enqueue_time: z.coerce.date().nullable(),
  finish_time: z.coerce.date().nullable(),
  evm_tx_hash: z.string().nullable(),
  failed_reason: z.string().nullable(),
});

export const LikeNFTMigrationSchema = z.object({
  id: z.number(),
  created_at: z.coerce.date(),
  likenft_asset_snapshot_id: z.number(),
  cosmos_address: z.string(),
  eth_address: z.string(),
  status: z.nativeEnum(LikeNFTMigrationStatus),
  failed_reason: z.string().nullable(),
});

export const LikeNFTMigrationDetailSchema = LikeNFTMigrationSchema.extend({
  classes: z.array(LikeNFTAssetMigrationClassSchema),
  nfts: z.array(LikeNFTAssetMigrationNFTSchema),
});

export type LikeNFTMigration = z.infer<typeof LikeNFTMigrationSchema>;
export type LikeNFTMigrationDetail = z.infer<
  typeof LikeNFTMigrationDetailSchema
>;
