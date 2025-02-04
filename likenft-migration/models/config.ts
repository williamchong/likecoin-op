import { z } from 'zod';

export const ConfigSchema = z.object({
  isTestnet: z.boolean().catch(true),
});

export type Config = z.infer<typeof ConfigSchema>;
