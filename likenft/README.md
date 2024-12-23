# Like NFT

## Dev

1. Start local dev node

```bash
$ npm run dev
```

2. Lint

```bash
$ npm run format
$ npm run lint [-- --fix]
```

2. Test locally

```bash
$ npm run build
$ npm run deploy:local
```

Retrieve the smart contract address from deploy output for further use

```bash
$ npm run deploy:local

> deploy:local
> hardhat run --network localhost scripts/deploy.ts

Deploying LikeNFT...
LikeNFT deployed to: 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
```

```bash
$ npm run debug:local
```

```typescript
> const LikeNFT = await ethers.getContractFactory("LikeNFT");
> const likeNFT = await LikeNFT.attach('0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512');
> const tx = await likeNFT.newClass({
    creator: signer.address,
    parent: {
      type_: 1,
      iscn_id_prefix: "abcdggiii123",
    },
    input: {
      name: "My Book",
      symbol: "KOOB",
      description: "Description",
      uri: "",
      uri_hash: "",
      metadata: "ipfs://",
      config: {
        burnable: true,
        max_supply: 10,
        blind_box_config: {
          mint_periods: [],
          reveal_time: 0,
        },
      },
    },
  });
> await tx.wait();
```

3. Upgrade

```bash
$ npm run upgrade:local

> upgrade:local
> hardhat run --network localhost scripts/upgrade.ts

Upgrading LikeNFT...
LikeNFT deployed to: 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
```

Please note that the deployed address should be the same

## Testnet

1. Update `.env` to provide the wallet private key, default owner address and default minter address.
1. To add a new network, add a config in hardhat config.
1. Deploy to the testnet.

```bash
$ npm run deploy:optimism-sepolia

...

EkilCoin deployed to: 0xfB0A48F5Ca78C8aA27432b64e89aF36423988Af8
```

Markdown the deployed address and put it in `PROXY_ADDRESS` in `.env`

1. Upgrade

```bash
$ npm run upgrade:optimism-sepolia

...

EkilCoin deployed to: 0xfB0A48F5Ca78C8aA27432b64e89aF36423988Af8
```

Please note that the deployed address should be the same
