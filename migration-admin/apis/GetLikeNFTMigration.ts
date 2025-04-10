import { z } from "zod";

import { makeAPI } from "./makeAPI";
import { LikeNFTMigrationDetailSchema } from "./models/likenftMigration";

export const ResponseSchema = z.object({
  migration: LikeNFTMigrationDetailSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeGetLikeNFTMigrationsAPI = (migrationId: number) =>
  makeAPI({
    url: `/admin/likenft/migration/${migrationId}`,
    method: "GET",
    responseSchema: ResponseSchema,
  });
