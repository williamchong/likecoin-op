import { z } from 'zod';

export const ConfigSchema = z.object({
  isTestnet: z.boolean().catch(true),
  apiBaseURL: z.string(),
  likeCoinAPIBaseURL: z.string(),
  likerlandUrlBase: z.string(),
  crispWebsiteId: z.string(),
});

export type Config = z.infer<typeof ConfigSchema>;
