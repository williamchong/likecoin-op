import { Decimal } from 'decimal.js';
import { z } from 'zod';

import { ViewCoin, ViewDenomSchema } from '../models/cosmosNetworkConfig';
import { makeAPI } from './makeAPI';

const ResponseSchema = z
  .object({
    denom: ViewDenomSchema,
    amount: z.string().transform(Decimal),
    decimals: z.number(),
  })
  .transform(
    (data): ViewCoin => ({
      denom: data.denom,
      amount: data.amount.div(Decimal(10).pow(data.decimals)).toString(),
    })
  );

export const getEvmPoolBalanceAPI = makeAPI({
  method: 'GET',
  url: `/likecoin/migration/evm-pool-balance`,
  responseSchema: ResponseSchema,
});

export type GetEvmPoolBalance = ReturnType<typeof getEvmPoolBalanceAPI>;
