import { z } from "zod";

export const ConfigSchema = z.object({
  apiBaseURL: z.string(),
});

export type Config = z.infer<typeof ConfigSchema>;
