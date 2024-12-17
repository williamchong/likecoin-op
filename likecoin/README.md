# Likecoin

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

```
$ npm run deploy:local

> deploy:local
> hardhat run --network localhost scripts/deploy.ts

Deploying EkilCoin...
EkilCoin deployed to: 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
```

```bash
$ npm run debug:local
```

```typescript
> const EkilCoin = await ethers.getContractFactory("EkilCoin");
> const ekilCoin = await EkilCoin.attach('0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512');
> await ekilCoin.mint('0x70997970C51812dc3A010C7d01b50e0d17dc79C8', 1000000000000000);
1n
```

3. Upgrade

```bash
$ npm run upgrade:local
> upgrade:local
> hardhat run --network localhost scripts/upgrade.ts

Upgrading EkilCoin...
EkilCoin deployed to: 0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512
```

Please note that the deployed address should be the same

## Testnet

1. Update `.env` to provide the wallet private key, default owner address and default minter address.
1. To add a new network, add a config in hardhat config.
1. Deploy to the testnet.

```bash
$ npm run deploy:optimism-sepolia

> deploy:optimism-sepolia
> DEBUG=@openzeppelin:* hardhat run --network optimism-sepolia scripts/deploy.ts

Deploying EkilCoin...
  @openzeppelin:upgrades:core manifest file: .openzeppelin/op-sepolia.json fallback file: .openzeppelin/unknown-11155420.json +0ms
  @openzeppelin:upgrades:core manifest file: .openzeppelin/op-sepolia.json fallback file: .openzeppelin/unknown-11155420.json +1s
  @openzeppelin:upgrades:core fetching deployment of implementation 89322429d44097607aa1735dcfd8d109798567e82798bafdab6318d00f17e571 +2ms
  @openzeppelin:upgrades:core found previous deployment 0xb3a36707e7709ca11c04ac71ab6dd5692c69efb667f1366f168ad54bb67fc0a1 +3ms
  @openzeppelin:upgrades:core resuming previous deployment 0xb3a36707e7709ca11c04ac71ab6dd5692c69efb667f1366f168ad54bb67fc0a1 +342ms
  @openzeppelin:upgrades:core polling timeout 0 polling interval 5000 +1ms
  @openzeppelin:upgrades:core verifying deployment tx mined 0xb3a36707e7709ca11c04ac71ab6dd5692c69efb667f1366f168ad54bb67fc0a1 +0ms
  @openzeppelin:upgrades:core succeeded verifying deployment tx mined 0xb3a36707e7709ca11c04ac71ab6dd5692c69efb667f1366f168ad54bb67fc0a1 +373ms
  @openzeppelin:upgrades:core verifying code in target address 0x04a772Dd557D39b331b23C963c75c686C4A4c521 +1ms
  @openzeppelin:upgrades:core code in target address found 0x04a772Dd557D39b331b23C963c75c686C4A4c521 +431ms
EkilCoin deployed to: 0xfB0A48F5Ca78C8aA27432b64e89aF36423988Af8
```

Markdown the deployed address and put it in `PROXY_ADDRESS` in `.env`

1. Upgrade

```bash
$ npm run upgrade:optimism-sepolia

> upgrade:optimism-sepolia
> DEBUG=@openzeppelin:* hardhat run --network optimism-sepolia scripts/upgrade.ts

Upgrading EkilCoin...
  @openzeppelin:upgrades:core manifest file: .openzeppelin/op-sepolia.json fallback file: .openzeppelin/unknown-11155420.json +0ms
  @openzeppelin:upgrades:core manifest file: .openzeppelin/op-sepolia.json fallback file: .openzeppelin/unknown-11155420.json +1s
  @openzeppelin:upgrades:core manifest file: .openzeppelin/op-sepolia.json fallback file: .openzeppelin/unknown-11155420.json +928ms
  @openzeppelin:upgrades:core fetching deployment of implementation 89322429d44097607aa1735dcfd8d109798567e82798bafdab6318d00f17e571 +1ms
  @openzeppelin:upgrades:core found previous deployment 0xb3a36707e7709ca11c04ac71ab6dd5692c69efb667f1366f168ad54bb67fc0a1 +2ms
  @openzeppelin:upgrades:core resuming previous deployment 0xb3a36707e7709ca11c04ac71ab6dd5692c69efb667f1366f168ad54bb67fc0a1 +445ms
  @openzeppelin:upgrades:core polling timeout 0 polling interval 5000 +2ms
  @openzeppelin:upgrades:core verifying deployment tx mined 0xb3a36707e7709ca11c04ac71ab6dd5692c69efb667f1366f168ad54bb67fc0a1 +0ms
  @openzeppelin:upgrades:core succeeded verifying deployment tx mined 0xb3a36707e7709ca11c04ac71ab6dd5692c69efb667f1366f168ad54bb67fc0a1 +322ms
  @openzeppelin:upgrades:core verifying code in target address 0x04a772Dd557D39b331b23C963c75c686C4A4c521 +1ms
  @openzeppelin:upgrades:core code in target address found 0x04a772Dd557D39b331b23C963c75c686C4A4c521 +424ms
EkilCoin deployed to: 0xfB0A48F5Ca78C8aA27432b64e89aF36423988Af8
```

Please note that the deployed address should be the same
