import { ImageURLConfig } from '~/models/config';

const IPFS_REGEX = /^ipfs:\/\/(.+)/;
const AR_REGEX = /^ar:\/\/(.+)/;

export function makeImageUrl(
  imageURLConfig: ImageURLConfig,
  u: string
): string {
  const ipfsMatched = IPFS_REGEX.exec(u);
  if (ipfsMatched != null) {
    return `${imageURLConfig.ipfsHTTPBaseURL}/${ipfsMatched[1]}`;
  }

  const arMatched = AR_REGEX.exec(u);
  if (arMatched != null) {
    return `${imageURLConfig.arHTTPBaseURL}/${arMatched[1]}`;
  }

  return u;
}

export { ImageURLConfig };
