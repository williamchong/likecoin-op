import { z } from "zod";

export const ConfigSchema = z.object({
  apiBaseURL: z.string(),
  isTestnet: z.boolean().catch(true),
});

export type Config = z.infer<typeof ConfigSchema>;
