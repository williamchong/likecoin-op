import { z } from "zod";

import { makeAPI } from "./makeAPI";
import {
  LikeNFTMigrationSchema,
  LikeNFTMigrationStatus,
} from "./models/likenftMigration";
import { NonNullParameter } from "~/utils/parameter";

export const RequestSchema = z
  .object({
    limit: z.number(),
    offset: z.number(),
    q: z.string().nullable(),
    status: z.nativeEnum(LikeNFTMigrationStatus).nullable(),
  })
  .transform((z) => {
    const params = new NonNullParameter(z).get();
    return new URLSearchParams(params);
  });

export type RequestParams = z.input<typeof RequestSchema>;

export const ResponseSchema = z.object({
  migrations: z.array(LikeNFTMigrationSchema),
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeListLikeNFTMigrationsAPI = makeAPI<Response, RequestParams>({
  url: `/admin/likenft/migration`,
  method: "GET",
  requestSchema: RequestSchema,
  responseSchema: ResponseSchema,
});
