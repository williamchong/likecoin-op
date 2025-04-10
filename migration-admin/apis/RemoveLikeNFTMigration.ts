import { z } from "zod";

import { makeAPI } from "./makeAPI";
import { LikeNFTMigrationSchema } from "./models/likenftMigration";

export const ResponseSchema = z.object({
  migration: LikeNFTMigrationSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeRemoveLikeNFTMigrationsAPI = (migrationId: number) =>
  makeAPI({
    url: `/admin/likenft/migration/${migrationId}`,
    method: "DELETE",
    responseSchema: ResponseSchema,
  });
