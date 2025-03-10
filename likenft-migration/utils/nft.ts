export function cosmosClassUrl(base: string, cosmosClassId: string): string {
  return new URL(`/nft/class/${cosmosClassId}`, base).toString();
}

export function cosmosNFTUrl(
  base: string,
  cosmosClassId: string,
  cosmosNFTId: string
): string {
  return new URL(`/nft/class/${cosmosClassId}/${cosmosNFTId}`, base).toString();
}
