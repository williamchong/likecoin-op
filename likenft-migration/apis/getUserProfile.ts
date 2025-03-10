import { z } from 'zod';

import { makeAPI } from './makeAPI';

export const UserProfileSchema = z.object({
  cosmos_wallet_address: z.string(),
  liker_id: z.string().nullable(),
  avatar: z.string().nullable(),
  eth_wallet_address: z.string().nullable(),
});

export const ResponseSchema = z.object({
  user_profile: UserProfileSchema,
});

export const makeGetUserProfileAPI = (cosmosWalletAddress: string) =>
  makeAPI({
    method: 'GET',
    url: `/user/profile/${cosmosWalletAddress}`,
    responseSchema: ResponseSchema,
  });
