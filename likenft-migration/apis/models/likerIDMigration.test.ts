import { LikerIDMigrationErrorSchema } from './likerIDMigration';

describe('LikerIDMigrationErrorSchema', () => {
  test('should recognize a code', () => {
    expect(
      LikerIDMigrationErrorSchema.parse({
        migrateBookUserError: 'EVM_WALLET_USED_BY_OTHER_USER',
        migrateLikerIdError: 'SOME_ERROR_LIKER_ID',
        migrateLikerLandError: 'SOME_ERROR_LIKER_LAND',
      })
    ).toEqual({
      isEnum: true,
      value: 'EVM_WALLET_USED_BY_OTHER_USER',
    });
  });

  test('should be unrecognized', () => {
    expect(
      LikerIDMigrationErrorSchema.parse({
        migrateBookUserError: 'SOME_OTHER_ERROR',
        migrateLikerIdError: 'SOME_ERROR_LIKER_ID',
        migrateLikerLandError: 'SOME_ERROR_LIKER_LAND',
      })
    ).toEqual({
      isEnum: false,
      value: 'SOME_OTHER_ERROR',
    });
  });

  test('should be likerIdError', () => {
    expect(
      LikerIDMigrationErrorSchema.parse({
        migrateBookUserError: '',
        migrateLikerIdError: 'SOME_ERROR_LIKER_ID',
        migrateLikerLandError: 'SOME_ERROR_LIKER_LAND',
      })
    ).toEqual({
      isEnum: false,
      value: 'SOME_ERROR_LIKER_ID',
    });
  });

  test('should be likerLandError', () => {
    expect(
      LikerIDMigrationErrorSchema.parse({
        migrateBookUserError: '',
        migrateLikerIdError: '',
        migrateLikerLandError: 'SOME_ERROR_LIKER_LAND',
      })
    ).toEqual({
      isEnum: false,
      value: 'SOME_ERROR_LIKER_LAND',
    });
  });
});
