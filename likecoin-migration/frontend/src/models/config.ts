import z from "zod";

export const ConfigSchema = z.object({
  cosmosChainId: z.string(),
  cosmosDenom: z.string(),
  cosmosDepositAddress: z.string(),
  cosmosFeeAmount: z.number(),
  cosmosFeeGas: z.number(),

  apiBaseUri: z.string(),
});

export type Config = z.infer<typeof ConfigSchema>;
