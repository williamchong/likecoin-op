import { z } from 'zod';

import { CoinStringSchema } from './coin';

export const LikeCoinMigrationStatusSchema = z.enum([
  'pending_cosmos_tx_hash',
  'verifying_cosmos_tx',
  'evm_minting',
  'evm_verifying',
  'completed',
  'failed',
]);

export type LikeCoinMigrationStatus = z.infer<
  typeof LikeCoinMigrationStatusSchema
>;

export const LikeCoinMigrationSchema = z.object({
  id: z.number(),
  created_at: z.coerce.date(),
  user_cosmos_address: z.string(),
  user_eth_address: z.string(),
  evm_signature: z.string(),
  amount: CoinStringSchema,
  status: LikeCoinMigrationStatusSchema,
  cosmos_tx_hash: z.string().nullable(),
  evm_tx_hash: z.string().nullable(),
  failed_reason: z.string().nullable(),
});

export type LikeCoinMigration = z.output<typeof LikeCoinMigrationSchema>;

export type Pending<M extends LikeCoinMigration> = M & {
  status: 'pending_cosmos_tx_hash';
};

export type Polling<M extends LikeCoinMigration> = M & {
  status: 'verifying_cosmos_tx' | 'evm_minting' | 'evm_verifying';
};

export type Completed<M extends LikeCoinMigration> = M & {
  status: 'completed';
};

export type Failed<M extends LikeCoinMigration> = M & {
  status: 'failed';
  failed_reason: string;
};
