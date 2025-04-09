import { z } from "zod";

export enum LikeNFTMigrationStatus {
  Init = "init",
  InProgress = "in_progress",
  Completed = "completed",
  Failed = "failed",
}

export const LikeNFTMigrationSchema = z.object({
  id: z.number(),
  created_at: z.coerce.date(),
  likenft_asset_snapshot_id: z.number(),
  cosmos_address: z.string(),
  eth_address: z.string(),
  status: z.nativeEnum(LikeNFTMigrationStatus),
  failed_reason: z.string().nullable(),
});

export type LikeNFTMigration = z.infer<typeof LikeNFTMigrationSchema>;
