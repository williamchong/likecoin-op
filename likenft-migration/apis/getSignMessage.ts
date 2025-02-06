import { z } from 'zod';

import { makeAPI } from './makeAPI';

export const RequestSchema = z.object({
  cosmos_address: z.string(),
  liker_id: z.string().nullable(),
  eth_address: z.string(),
});

export type Request = z.infer<typeof RequestSchema>;

export const ResponseSchema = z.object({
  message: z.string(),
});

export type Response = z.infer<typeof ResponseSchema>;

export const getSignMessage = makeAPI({
  url: '/likenft/signing_message',
  method: 'POST',
  requestSchema: RequestSchema,
  responseSchema: ResponseSchema,
});
