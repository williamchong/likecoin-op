import { ImageURLConfig, makeImageUrl } from './imageUrl';

test('makeImageUrl', () => {
  const imageURLConfig: ImageURLConfig = {
    ipfsHTTPBaseURL: 'https://ipfs.io/ipfs',
    arHTTPBaseURL: 'https://arweave.net',
  };
  expect(
    makeImageUrl(
      imageURLConfig,
      'ar://yY1IjxpTtp_QItTzirhOREYmFZEeW9IjlEmRcp-Cz5w'
    )
  ).toBe('https://arweave.net/yY1IjxpTtp_QItTzirhOREYmFZEeW9IjlEmRcp-Cz5w');
  expect(
    makeImageUrl(
      imageURLConfig,
      'ipfs://bafybeierwqwwtj7wynjaud2jwi5yjxfqnnthvxfky66suih5wlpjuofvey'
    )
  ).toBe(
    'https://ipfs.io/ipfs/bafybeierwqwwtj7wynjaud2jwi5yjxfqnnthvxfky66suih5wlpjuofvey'
  );
  expect(makeImageUrl(imageURLConfig, 'https://placehold.co/600x400')).toBe(
    'https://placehold.co/600x400'
  );
});
