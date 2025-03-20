import { z } from 'zod';

import { makeAPI } from './makeAPI';

export const RequestSchema = z.object({
  cosmos_address: z.string(),
  cosmos_signature: z.string(),
  cosmos_public_key: z.string(),
  cosmos_signature_content: z.string(),
  signMethod: z.enum(['']),
});

export const ResponseSchema = z.object({
  isMigratedLikerId: z.boolean(),
  isMigratedLikerLand: z.boolean(),
  migratedLikerId: z.string().nullable(),
  migratedLikerLandUser: z.string().nullable(),
  migrateLikerIdError: z.string().nullable(),
  migrateLikerLandError: z.string().nullable(),
});

export const makeMigrateUserEvmWallet = makeAPI({
  url: '/wallet/evm/migrate/user',
  method: 'POST',
  requestSchema: RequestSchema,
  responseSchema: ResponseSchema,
});
