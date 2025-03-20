export interface MigrateUserEvmWalletMemoData {
  /**
   *  'migrate'.
   *
   *  Ref: https://github.com/likecoin/likecoin-api-public/blob/f5ff7c0e9898d51fddad8e070a9829922fcc87dd/src/routes/wallet/index.ts#L92
   */
  action: 'migrate';

  /**
   * should match input wallet.
   *
   * ref: https://github.com/likecoin/likecoin-api-public/blob/8417033b761550e1e4cd430306f8d74e103233db/src/util/api/users/index.ts#L186
   */
  cosmosWallet: string;

  /**
   * should match input wallet.
   *
   * ref: https://github.com/likecoin/likecoin-api-public/blob/8417033b761550e1e4cd430306f8d74e103233db/src/util/api/users/index.ts#L186
   */
  likeWallet: string;

  /**
   * Current timestamp.
   *
   * Will be checked if within 5 mins.
   *
   * ref: https://github.com/likecoin/likecoin-api-public/blob/8417033b761550e1e4cd430306f8d74e103233db/src/util/api/users/index.ts#L189
   */
  ts: number;

  /**
   * The evm wallet to submit
   *
   * ref: https://github.com/likecoin/likecoin-api-public/blob/f5ff7c0e9898d51fddad8e070a9829922fcc87dd/src/routes/wallet/index.ts#L97
   */
  evm_wallet: string;
}

export function makeMigrateUserEvmWalletMemoData(
  obj: MigrateUserEvmWalletMemoData
): MigrateUserEvmWalletMemoData {
  return obj;
}
