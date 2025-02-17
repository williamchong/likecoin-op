import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { LikeNFTAssetSnapshotSchema } from './models/likenftAssetSnapshot';

export const ResponseSchema = z.object({
  preview: LikeNFTAssetSnapshotSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeGetMigrationPreviewAPI = (cosmosAddress: string) =>
  makeAPI({
    url: `/likenft/migration-preview/${cosmosAddress}`,
    method: 'GET',
    responseSchema: ResponseSchema,
  });
