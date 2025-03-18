import { z } from 'zod';

export const ConfigSchema = z.object({
  isTestnet: z.boolean().catch(true),
  apiBaseURL: z.string(),
  cosmosDepositAddress: z.string(),
});

export type Config = z.infer<typeof ConfigSchema>;
