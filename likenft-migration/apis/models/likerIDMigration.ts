import { z } from 'zod';

import { EmptyStringToNullSchema } from '../../schemas/EmptyStringToNullSchema';
import {
  RecognizableEnum,
  UnrecognizableEnum,
} from '../../schemas/RecognizableEnum';

const MIGRATE_BOOK_USER_ERRORS = ['EVM_WALLET_USED_BY_OTHER_USER'] as const;

export const LikerIDMigrationErrorSchema = z
  .union([
    z
      .object({
        migrateBookUserError: EmptyStringToNullSchema.pipe(
          RecognizableEnum(MIGRATE_BOOK_USER_ERRORS)
        ),
      })
      .transform((migrateBookUserError) => ({
        type: 'bookUserError' as const,
        ...migrateBookUserError,
      })),
    z
      .object({
        migrateLikerIdError: EmptyStringToNullSchema.pipe(UnrecognizableEnum),
      })
      .transform((migrateLikerIdError) => ({
        type: 'likerIdError' as const,
        ...migrateLikerIdError,
      })),
    z
      .object({
        migrateLikerLandError: EmptyStringToNullSchema.pipe(UnrecognizableEnum),
      })
      .transform((migrateLikerLandError) => ({
        type: 'likerLandError' as const,
        ...migrateLikerLandError,
      })),
  ])
  .transform((v) => {
    switch (v.type) {
      case 'bookUserError':
        return v.migrateBookUserError;
      case 'likerIdError':
        return v.migrateLikerIdError;
      case 'likerLandError':
        return v.migrateLikerLandError;
    }
  });

export type LikerIDMigrationError = z.infer<typeof LikerIDMigrationErrorSchema>;
