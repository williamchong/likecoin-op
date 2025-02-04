export function getSignMessage(
  likerID: string | null,
  cosmosAddress: string
): Promise<string> {
  return new Promise((resolve) => {
    console.log(`TODO: getSignMessage
${{ likerID, cosmosAddress }}`);
    setTimeout(() => {
      resolve(`Liker ID: ${likerID}
Cosmos address: ${cosmosAddress}
Nonce: some-nonce`);
    }, 2000);
  });
}
