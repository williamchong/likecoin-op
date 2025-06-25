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
});

export const LoadingLikeNFTAssetSnapshotSchema =
  LikeNFTAssetSnapshotSchemaBase.extend({
    status: z.enum(['init', 'in_progress']),
  });

export type LoadingLikeNFTAssetSnapshot = z.infer<
  typeof LoadingLikeNFTAssetSnapshotSchema
>;

export const EmptyLikeNFTAssetSnapshotSchema =
  LikeNFTAssetSnapshotSchemaBase.extend({
    status: z.enum(['completed']),
    classes: z.tuple([]),
    nfts: z.tuple([]),
  });

export type EmptyLikeNFTAssetSnapshot = z.infer<
  typeof EmptyLikeNFTAssetSnapshotSchema
>;

export const NonEmptyLikeNFTAssetSnapshotSchema = z.union([
  LikeNFTAssetSnapshotSchemaBase.extend({
    status: z.enum(['completed']),
    classes: z.array(LikeNFTAssetSnapshotClassSchema).nonempty(),
    nfts: z.array(LikeNFTAssetSnapshotNFTSchema),
  }),
  LikeNFTAssetSnapshotSchemaBase.extend({
    status: z.enum(['completed']),
    classes: z.array(LikeNFTAssetSnapshotClassSchema),
    nfts: z.array(LikeNFTAssetSnapshotNFTSchema).nonempty(),
  }),
]);

export type NonEmptyLikeNFTAssetSnapshot = z.infer<
  typeof NonEmptyLikeNFTAssetSnapshotSchema
>;

export const FailedLikeNFTAssetSnapshotSchema =
  LikeNFTAssetSnapshotSchemaBase.extend({
    status: z.enum(['failed']),
    failed_reason: z.string(),
  });

export type FailedLikeNFTAssetSnapshot = z.infer<
  typeof FailedLikeNFTAssetSnapshotSchema
>;

export const LikeNFTAssetSnapshotSchema = z.union([
  LoadingLikeNFTAssetSnapshotSchema,
  EmptyLikeNFTAssetSnapshotSchema,
  NonEmptyLikeNFTAssetSnapshotSchema,
  FailedLikeNFTAssetSnapshotSchema,
]);

export type LikeNFTAssetSnapshot = z.infer<typeof LikeNFTAssetSnapshotSchema>;

export function isLoadingLikeNFTAssetSnapshot(
  s: LikeNFTAssetSnapshot
): s is LoadingLikeNFTAssetSnapshot {
  return s.status === 'init' || s.status === 'in_progress';
}

export function isEmptyLikeNFTAssetSnapshot(
  s: LikeNFTAssetSnapshot
): s is EmptyLikeNFTAssetSnapshot {
  return (
    s.status === 'completed' && s.nfts.length === 0 && s.classes.length === 0
  );
}

export function isNonEmptyLikeNFTAssetSnapshot(
  s: LikeNFTAssetSnapshot
): s is NonEmptyLikeNFTAssetSnapshot {
  return (
    s.status === 'completed' && (s.nfts.length !== 0 || s.classes.length !== 0)
  );
}

export function isFailedLikeNFTAssetSnapshot(
  s: LikeNFTAssetSnapshot
): s is FailedLikeNFTAssetSnapshot {
  return s.status === 'failed';
}
