import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { LikeNFTAssetSnapshotSchema } from './models/likenftAssetSnapshot';

export const RequestSchema = z.object({
  cosmos_address: z.string(),
});

export type Request = z.infer<typeof RequestSchema>;

export const ResponseSchema = z.object({
  preview: LikeNFTAssetSnapshotSchema,
});

export type Response = z.infer<typeof ResponseSchema>;

export const makeCreateMigrationPreviewAPI = makeAPI({
  url: `/likenft/migration-preview`,
  method: 'POST',
  requestSchema: RequestSchema,
  responseSchema: ResponseSchema,
});
