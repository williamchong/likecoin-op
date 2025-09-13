import { z } from 'zod';

export const EVMChainConfigSchema = z.object({
  chainName: z.string(),
  chainId: z.string(),
  rpcUrls: z.array(z.string()),
  nativeCurrency: z.object({
    name: z.string(),
    symbol: z.string(),
    decimals: z.number(),
  }),
  blockExplorerUrls: z.array(z.string()),
});

export type EVMChainConfig = z.infer<typeof EVMChainConfigSchema>;
