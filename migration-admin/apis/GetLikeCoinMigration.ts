import { z } from "zod";

import { makeAPI } from "./makeAPI";
import { LikeCoinMigrationSchema } from "./models/likecoinMigration";

export const ResponseSchema = z.object({
  migration: LikeCoinMigrationSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeGetLikeCoinMigrationsAPI = (migrationId: number) =>
  makeAPI({
    url: `/admin/likecoin/migration/${migrationId}`,
    method: "GET",
    responseSchema: ResponseSchema,
  });
