import { z } from 'zod';

export const ConfigSchema = z.object({
  isTestnet: z.boolean().catch(true),
  authcoreRedirectUrl: z.string(),
  apiBaseURL: z.string(),
  cosmosDepositAddress: z.string(),
  cosmosExplorerBaseURL: z.string(),
  evmExplorerBaseURL: z.string(),
  cosmosLikeCoinNetworkConfigPath: z.string(),
});

export type Config = z.infer<typeof ConfigSchema>;
