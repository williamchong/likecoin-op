import { Coin } from '@keplr-wallet/types';
import { Decimal } from 'decimal.js';
import { z } from 'zod';

export const ViewDenomSchema = z.enum(['LIKE', 'EKIL']);
export type ViewDenom = z.infer<typeof ViewDenomSchema>;
export const ChainDenomSchema = z.enum(['nanoekil', 'nanolike']);
export type ChainDenom = z.infer<typeof ChainDenomSchema>;

export function isViewDenom(denom: string): denom is ViewDenom {
  const { success } = ViewDenomSchema.safeParse(denom);
  return success;
}

export function isChainDenom(denom: string): denom is ChainDenom {
  const { success } = ChainDenomSchema.safeParse(denom);
  return success;
}

export const CoinLookupItemSchema = z.object({
  viewDenom: ViewDenomSchema,
  chainDenom: ChainDenomSchema,
  chainToViewConversionFactor: z.string().transform(Decimal),
  icon: z.string(),
  coinGeckoId: z.string(),
});

export type CoinLookupItem = z.infer<typeof CoinLookupItemSchema>;

export const FeeOptionSchema = z.object({
  denom: ViewDenomSchema,
  amount: z.number().transform(String),
});

export type FeeOption = z.infer<typeof FeeOptionSchema>;

export const CosmosNetworkConfigSchema = z.object({
  coinLookup: z.tuple([CoinLookupItemSchema]),
});

export type CosmosNetworkConfig = z.infer<typeof CosmosNetworkConfigSchema>;

export interface ChainCoin {
  denom: ChainDenom;
  amount: string;
}

export interface ViewCoin {
  denom: ViewDenom;
  amount: string;
}

export function isChainCoin(coin: Coin): coin is ChainCoin {
  return isChainDenom(coin.denom);
}

export function isViewCoin(coin: Coin): coin is ViewCoin {
  return isViewDenom(coin.denom);
}

export function convertViewCoinToChainCoin(
  coin: ViewCoin,
  networkConfig: CosmosNetworkConfig
): ChainCoin {
  const coinLookup = networkConfig.coinLookup[0];
  return {
    denom: coinLookup.chainDenom,
    amount: Decimal(coin.amount)
      .div(Decimal(coinLookup.chainToViewConversionFactor))
      .toString(),
  };
}

export function convertChainCoinToViewCoin(
  coin: ChainCoin,
  networkConfig: CosmosNetworkConfig
): ViewCoin {
  const coinLookup = networkConfig.coinLookup[0];
  return {
    denom: coinLookup.viewDenom,
    amount: Decimal(coin.amount)
      .mul(Decimal(coinLookup.chainToViewConversionFactor))
      .toString(),
  };
}
