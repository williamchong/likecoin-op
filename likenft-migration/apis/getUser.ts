interface User {
  cosmos_wallet_address: string;
  liker_id: string | null;
  evm_address: string | null;
  is_authcore: boolean;
}
export function getUser(cosmosWalletAddress: string): Promise<User> {
  return new Promise((resolve) => {
    console.log(`TODO: getUser
${{ cosmosWalletAddress }}`);
    setTimeout(() => {
      resolve({
        cosmos_wallet_address: cosmosWalletAddress,
        liker_id: 'my-liker-id',
        evm_address: null,
        is_authcore: false,
      });
    }, 2000);
  });
}
