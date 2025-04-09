import { z } from "zod";

import { makeAPI } from "./makeAPI";
import {
  LikeCoinMigrationSchema,
  LikeCoinMigrationStatus,
} from "./models/likecoinMigration";
import { NonNullParameter } from "~/utils/parameter";

export const RequestSchema = z
  .object({
    limit: z.number(),
    offset: z.number(),
    q: z.string().nullable(),
    status: z.nativeEnum(LikeCoinMigrationStatus).nullable(),
  })
  .transform((z) => {
    const params = new NonNullParameter(z).get();
    return new URLSearchParams(params);
  });

export type RequestParams = z.input<typeof RequestSchema>;

export const ResponseSchema = z.object({
  migrations: z.array(LikeCoinMigrationSchema),
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeListLikeCoinMigrationsAPI = makeAPI<Response, RequestParams>({
  url: `/admin/likecoin/migration`,
  method: "GET",
  requestSchema: RequestSchema,
  responseSchema: ResponseSchema,
});
