const IPFS_REGEX = /^ipfs:\/\/(.+)/;
const AR_REGEX = /^ar:\/\/(.+)/;

export function makeImageUrl(u: string): string {
  const ipfsMatched = IPFS_REGEX.exec(u);
  if (ipfsMatched != null) {
    return `https://ipfs.io/ipfs/${ipfsMatched[1]}`;
  }

  const arMatched = AR_REGEX.exec(u);
  if (arMatched != null) {
    return `https://arweave.net/${arMatched[1]}`;
  }

  return u;
}
