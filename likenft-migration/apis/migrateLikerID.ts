import { z } from 'zod';

import { makeAPI } from './makeAPI';

export const RequestSchema = z.object({
  cosmos_pub_key: z.string(),
  like_id: z.string().nullable(),
  eth_address: z.string(),
  cosmos_signature: z.string(),
  eth_signature: z.string(),
  signing_message: z.string(),
});

export const ResponseSchema = z.object({
  message: z.string(),
});

export const makeMigrateLikerIDAPI = makeAPI({
  url: '/likenft/likerid/migration',
  method: 'POST',
  requestSchema: RequestSchema,
  responseSchema: ResponseSchema,
});
