import type { NuxtAxiosInstance } from '@nuxtjs/axios';
import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { LikeCoinMigrationSchema } from './models/likeCoinMigration';

const RequestSchema = z.object({
  cosmos_tx_hash: z.string(),
});

type Request = z.infer<typeof RequestSchema>;

const ResponseSchema = z.object({
  migration: LikeCoinMigrationSchema,
});

export const makeUpdateLikeCoinMigrationCosmosTxHash =
  (axiosInstance: NuxtAxiosInstance) => (cosmosAddress: string, req: Request) =>
    makeAPI({
      method: 'PUT',
      url: `/likecoin/migration/${cosmosAddress}/cosmos-tx-hash`,
      responseSchema: ResponseSchema,
      requestSchema: RequestSchema,
    })(axiosInstance)(req);

export type UpdateLikeCoinMigrationCosmosTxHash = ReturnType<
  typeof makeUpdateLikeCoinMigrationCosmosTxHash
>;
