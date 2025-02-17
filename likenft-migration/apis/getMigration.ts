import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { LikeNFTAssetMigrationSchema } from './models/likenftAssetMigration';

export const ResponseSchema = z.object({
  migration: LikeNFTAssetMigrationSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeGetMigrationAPI = (cosmosAddress: string) =>
  makeAPI({
    url: `/likenft/migration/${cosmosAddress}`,
    method: 'GET',
    responseSchema: ResponseSchema,
  });
