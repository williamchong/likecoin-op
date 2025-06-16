import { z } from 'zod';

export const ImageURLConfigSchema = z.object({
  ipfsHTTPBaseURL: z.string().default('https://ipfs.io/ipfs'),
  arHTTPBaseURL: z.string().default('https://arweave.net'),
});

export type ImageURLConfig = z.infer<typeof ImageURLConfigSchema>;

export const ConfigSchema = z
  .object({
    transitionLoadingScreenDelayMs: z.number().default(0),
    isTestnet: z.boolean().catch(true),
    authcoreRedirectUrl: z.string(),
    apiBaseURL: z.string(),
    likerlandUrlBase: z.string(),
    crispWebsiteId: z.string(),
    googleAnalyticsTagId: z.string().default(''),
  })
  .merge(ImageURLConfigSchema);

export type Config = z.infer<typeof ConfigSchema>;
