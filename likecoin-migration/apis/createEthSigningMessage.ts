import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { CoinSchema } from './models/coin';

const RequestSchema = z.object({
  amount: CoinSchema,
});

const ResponseSchema = z.object({
  signing_message: z.string(),
});

export const makeCreateEthSigningMessageAPI = makeAPI({
  method: 'POST',
  url: `/likecoin/migration/eth-signing-message`,
  responseSchema: ResponseSchema,
  requestSchema: RequestSchema,
});

export type CreateEthSigningMessage = ReturnType<
  typeof makeCreateEthSigningMessageAPI
>;
