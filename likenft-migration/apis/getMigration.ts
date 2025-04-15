import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { LikeNFTAssetMigrationSchema } from './models/likenftAssetMigration';
import { LikeNFTAssetSnapshotSchema } from './models/likenftAssetSnapshot';

export const ResponseSchema = z.object({
  migration: LikeNFTAssetMigrationSchema,
  snapshot: LikeNFTAssetSnapshotSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeGetMigrationAPI = (cosmosAddress: string) =>
  makeAPI({
    url: `/likenft/migration/${cosmosAddress}`,
    method: 'GET',
    responseSchema: ResponseSchema,
  });
