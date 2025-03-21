import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { CoinSchema } from './models/coin';

const RequestSchema = z.object({
  signature: z.string(),
  ethAddress: z.string(),
  amount: CoinSchema,
});

const ResponseSchema = z.object({
  memo_data: z.string(),
});

export const makeCreateCosmosMemoDataAPI = makeAPI({
  method: 'POST',
  url: `/likecoin/migration/cosmos-memo-data`,
  responseSchema: ResponseSchema,
  requestSchema: RequestSchema,
});

export type CreateCosmosMemoData = ReturnType<
  typeof makeCreateCosmosMemoDataAPI
>;
