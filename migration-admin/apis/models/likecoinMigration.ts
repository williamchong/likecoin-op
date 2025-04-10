import { z } from "zod";

export enum LikeCoinMigrationStatus {
  PendingCosmosTxHash = "pending_cosmos_tx_hash",
  VerifyingCosmosTx = "verifying_cosmos_tx",
  EvmMinting = "evm_minting",
  EvmVerifying = "evm_verifying",
  Completed = "completed",
  Failed = "failed",
}

export const LikeCoinMigrationSchema = z.object({
  id: z.number(),
  created_at: z.coerce.date(),
  user_cosmos_address: z.string(),
  user_eth_address: z.string(),
  evm_signature: z.string(),
  amount: z.string(),
  status: z.nativeEnum(LikeCoinMigrationStatus),
  cosmos_tx_hash: z.string().nullable(),
  evm_tx_hash: z.string().nullable(),
  failed_reason: z.string().nullable(),
});

export type LikeCoinMigration = z.infer<typeof LikeCoinMigrationSchema>;
