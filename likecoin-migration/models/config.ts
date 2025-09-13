import { z } from 'zod';

export const ConfigSchema = z.object({
  transitionLoadingScreenDelayMs: z.number().default(0),
  isTestnet: z.boolean().catch(true),
  authcoreRedirectUrl: z.string(),
  apiBaseURL: z.string(),
  cosmosDepositAddress: z.string(),
  cosmosExplorerBaseURL: z.string(),
  evmExplorerBaseURL: z.string(),
  cosmosLikeCoinNetworkConfigPath: z.string(),
  googleAnalyticsTagId: z.string().default(''),
  intercomAppId: z.string(),
});

export type Config = z.infer<typeof ConfigSchema>;
