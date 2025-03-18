import { z } from 'zod';

import { makeAPI } from './makeAPI';
import { CoinSchema } from './models/coin';
import { LikeCoinMigrationSchema } from './models/likeCoinMigration';

const RequestSchema = z.object({
  eth_address: z.string(),
  cosmos_address: z.string(),
  evm_signature: z.string(),
  evm_signature_message: z.string(),
  amount: CoinSchema,
});

const ResponseSchema = z.object({
  migration: LikeCoinMigrationSchema,
});

export const makeCreateLikeCoinMigrationAPI = makeAPI({
  method: 'POST',
  url: `/likecoin/migration`,
  responseSchema: ResponseSchema,
  requestSchema: RequestSchema,
});

export type CreateLikeCoinMigration = ReturnType<
  typeof makeCreateLikeCoinMigrationAPI
>;
