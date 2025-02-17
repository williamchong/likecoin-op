import { z } from 'zod';

export const LikeNFTAssetSnapshotClassSchema = z.object({
  id: z.number(),
  nft_snapshot_id: z.number(),
  created_at: z.coerce.date(),
  cosmos_class_id: z.string(),
  name: z.string(),
  image: z.string(),
});

export type LikeNFTAssetSnapshotClass = z.infer<
  typeof LikeNFTAssetSnapshotClassSchema
>;

export const LikeNFTAssetSnapshotNFTSchema = z.object({
  id: z.number(),
  nft_snapshot_id: z.number(),
  created_at: z.coerce.date(),
  cosmos_class_id: z.string(),
  cosmos_nft_id: z.string(),
  name: z.string(),
  image: z.string(),
});

export type LikeNFTAssetSnapshotNFT = z.infer<
  typeof LikeNFTAssetSnapshotNFTSchema
>;

export const LikeNFTAssetSnapshotStatusSchema = z.enum([
  'init',
  'in_progress',
  'completed',
  'failed',
]);

export const LikeNFTAssetSnapshotSchema = z.object({
  id: z.number(),
  created_at: z.coerce.date(),
  cosmos_address: z.string(),
  block_height: z.string().nullable(),
  block_time: z.coerce.date().nullable(),
  status: LikeNFTAssetSnapshotStatusSchema,
  failed_reason: z.string().nullable(),
  classes: z.array(LikeNFTAssetSnapshotClassSchema),
  nfts: z.array(LikeNFTAssetSnapshotNFTSchema),
});

export type LikeNFTAssetSnapshot = z.infer<typeof LikeNFTAssetSnapshotSchema>;
