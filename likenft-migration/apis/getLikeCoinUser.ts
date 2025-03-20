import { z } from 'zod';

import { makeAPI } from './makeAPI';

export const LikeCoinUserSchema = z.object({
  user: z.string(),
  displayName: z.string(),
  avatar: z.string().nullish(),
  cosmosWallet: z.string().nullish(),
  likeWallet: z.string().nullish(),
  evmWallet: z.string().nullish(),
});

export const ResponseSchema = LikeCoinUserSchema;

export const makeGetLikeCoinUserAPI = (cosmosWalletAddress: string) =>
  makeAPI({
    method: 'GET',
    url: `/users/addr/${cosmosWalletAddress}/min`,
    responseSchema: ResponseSchema,
  });
