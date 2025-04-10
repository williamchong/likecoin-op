const IPFS_REGEX = /^ipfs:\/\/(.+)/;

export function makeImageUrl(u: string): string {
  const ipfsMatched = IPFS_REGEX.exec(u);
  if (ipfsMatched != null) {
    return `https://ipfs.io/ipfs/${ipfsMatched[1]}`;
  }
  return u;
}
