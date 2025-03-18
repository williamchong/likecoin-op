import type { NuxtAxiosInstance } from '@nuxtjs/axios';
import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { LikeCoinMigrationSchema } from './models/likeCoinMigration';

const ResponseSchema = z.object({
  migration: LikeCoinMigrationSchema,
});

export const makeGetLatestLikeCoinMigrationAPI =
  (axiosInstance: NuxtAxiosInstance) => (cosmosAddress: string) =>
    makeAPI({
      method: 'GET',
      url: `/likecoin/migration/${cosmosAddress}`,
      responseSchema: ResponseSchema,
    })(axiosInstance)();

export type GetLatestLikeCoinMigration = ReturnType<
  typeof makeGetLatestLikeCoinMigrationAPI
>;
