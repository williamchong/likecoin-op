import { parseCoins } from '@cosmjs/stargate';

import { ChainCoin } from '~/models/cosmosNetworkConfig';

/**
 * Serializing coin to string format such that it can be parsed
 * by `parseCoin` from `@cosmjs/stargate`
 * @param coin
 * @returns string
 */
export function serializeCoin(coin: ChainCoin): string {
  return `${coin.amount}${coin.denom}`;
}

export function parseCoin(coinStr: string): ChainCoin {
  return parseCoins(coinStr)[0] as ChainCoin;
}
