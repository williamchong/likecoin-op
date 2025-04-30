import { makeImageUrl } from './imageUrl';

test('makeImageUrl', () => {
  expect(makeImageUrl('ar://yY1IjxpTtp_QItTzirhOREYmFZEeW9IjlEmRcp-Cz5w')).toBe(
    'https://arweave.net/yY1IjxpTtp_QItTzirhOREYmFZEeW9IjlEmRcp-Cz5w'
  );
  expect(
    makeImageUrl(
      'ipfs://bafybeierwqwwtj7wynjaud2jwi5yjxfqnnthvxfky66suih5wlpjuofvey'
    )
  ).toBe(
    'https://ipfs.io/ipfs/bafybeierwqwwtj7wynjaud2jwi5yjxfqnnthvxfky66suih5wlpjuofvey'
  );
  expect(makeImageUrl('https://placehold.co/600x400')).toBe(
    'https://placehold.co/600x400'
  );
});
