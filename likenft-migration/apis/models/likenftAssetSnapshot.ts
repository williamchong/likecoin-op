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

const LikeNFTAssetSnapshotSchemaBase = z.object({
  id: z.number(),
  created_at: z.coerce.date(),
  cosmos_address: z.string(),
  block_height: z.string().nullable(),
  block_time: z.coerce.date().nullable(),
  status: LikeNFTAssetSnapshotStatusSchema,
  failed_reason: z.string().nullable(),
});

export const EmptyLikeNFTAssetSnapshotSchema =
  LikeNFTAssetSnapshotSchemaBase.extend({
    classes: z.tuple([]),
    nfts: z.tuple([]),
  });

export type EmptyLikeNFTAssetSnapshot = z.infer<
  typeof EmptyLikeNFTAssetSnapshotSchema
>;

export const NonEmptyLikeNFTAssetSnapshotSchema = z.union([
  LikeNFTAssetSnapshotSchemaBase.extend({
    classes: z.array(LikeNFTAssetSnapshotClassSchema).nonempty(),
    nfts: z.array(LikeNFTAssetSnapshotNFTSchema),
  }),
  LikeNFTAssetSnapshotSchemaBase.extend({
    classes: z.array(LikeNFTAssetSnapshotClassSchema),
    nfts: z.array(LikeNFTAssetSnapshotNFTSchema).nonempty(),
  }),
]);

export type NonEmptyLikeNFTAssetSnapshot = z.infer<
  typeof NonEmptyLikeNFTAssetSnapshotSchema
>;

export const LikeNFTAssetSnapshotSchema = z.union([
  EmptyLikeNFTAssetSnapshotSchema,
  NonEmptyLikeNFTAssetSnapshotSchema,
]);

export type LikeNFTAssetSnapshot = z.infer<typeof LikeNFTAssetSnapshotSchema>;

export function isEmptyLikeNFTAssetSnapshot(
  s: LikeNFTAssetSnapshot
): s is EmptyLikeNFTAssetSnapshot {
  return s.nfts.length === 0 && s.classes.length === 0;
}

export function isNonEmptyLikeNFTAssetSnapshot(
  s: LikeNFTAssetSnapshot
): s is NonEmptyLikeNFTAssetSnapshot {
  return !isEmptyLikeNFTAssetSnapshot(s);
}
