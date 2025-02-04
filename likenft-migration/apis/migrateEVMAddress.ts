export function migrateEVMAddress(
  cosmosWalletAddress: string,
  evmWalletAddress: string,
  likerID: string | null,
  signature: string,
  nonce: string
): Promise<void> {
  return new Promise<void>((resolve) => {
    console.log(`TODO: migrateEVMAddress
${{ cosmosWalletAddress, evmWalletAddress, likerID, signature, nonce }}`);
    setTimeout(() => {
      resolve();
    }, 2000);
  });
}
