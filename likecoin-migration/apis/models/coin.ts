import { z } from 'zod';

import { ChainDenomSchema } from '~/models/cosmosNetworkConfig';
import { parseCoin, serializeCoin } from '~/utils/cosmos';

export const CoinSchema = z
  .object({ denom: ChainDenomSchema, amount: z.string() })
  .transform(serializeCoin);

export const CoinStringSchema = z.string().transform(parseCoin);
