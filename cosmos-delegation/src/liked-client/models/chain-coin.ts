import { z } from "zod";

export const ChainCoinSchema = z.object({
  denom: z.string(),
  amount: z.coerce.bigint(),
});

export type ChainCoin = z.infer<typeof ChainCoinSchema>;
