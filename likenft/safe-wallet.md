# Production guide on upgrade LikeProtocol

According to the operation guide, an offline key is ued to deploy the initial contract to keep predictable address. And then transfer the ownership to a community owned multisif Safe wallet. This guide is for operator to create proposal to the Safe wallet.

We assume the subsequence LikeProtocol/BookNFT are deployed by random address owned by the developer at the time, to ease the operation overhead, as the implementation address is not effecting user expiernece.

Following command are operate on sepolia.

## Upgrade only BookNFT

#### Step Overview

1. Deploy the new version of BookNFT
1. Propose upgrade of new BookNFT implementation on https://app.safe.global/

#### Check the protocol status

```
ERC721_PROXY_ADDRESS=0xC513ffcaab6f5aC669055D09f4dC0C9A3dA12c05 npm run script:sepolia scripts/queryBookNFTImpl.ts
```

#### Deploy V2 of BookNFT

```
npm run script:sepolia scripts/deployBookNFT.ts

New BookNFT implementation is deployed to: 0x0F8d5F709c6916D332c08Be447a88e8551A43EFf
```

#### At safe.global:

- `New transcation`
- `Transaction Builder`
- Fill in the LikeProtol Proxy address
- Paste the current version of LikeProtocol abi, which is find in last tag of `abi/LikeProtocol.json`
- The interface should show avalible function, pick upgradeTo
- Paste the above script output of BookNFT address to `newImplementation`
- Submit the proposal
- Ask the owner to sign with attached expected payload for cross check.
- Run `npm run script:sepolia scripts/queryBookNFTImpl.ts` to verify after all party signed

## Upgrade both LikeProtocol & BookNFT

Step Overview

1. Prepare the new version of BookNFT
1. Prepare the new version of LikeProtocol
1. Prepare the byte payload for `upgradeToAndCall`
1. Propose upgrade of new BookNFT and LikeProtocal in atomic call on https://app.safe.global/

#### First deploy the new version of BookNFT

```
npm run script:sepolia scripts/deployBookNFT.ts

New BookNFT implementation is deployed to: 0x0F8d5F709c6916D332c08Be447a88e8551A43EFf
```

Note: Dev verify the New version of BookNFT is properly initialized, it often have out of gas error.

#### Prepare the new version of LikeProtocol implementation

```
ERC721_PROXY_ADDRESS=0xC513ffcaab6f5aC669055D09f4dC0C9A3dA12c05 npm run script:sepolia scripts/prepareLikeProtocol.ts

> script:sepolia
> hardhat run --network sepolia scripts/prepareLikeProtocol.ts

Deployer: 0xC71fe89e4C0e5458a793fc6548EF6B392417A7Fb
Preparing Upgrade of LikeProtocol... 0xC513ffcaab6f5aC669055D09f4dC0C9A3dA12c05
LikeProtocol new implementation is deployed to: 0x0e43789dAe6E3F16B3411A43ecFA5c6bbDC96a71

... verify info
```

#### Prepare the byte payload for `upgradeToAndCall`

```
BOOKNFT_ADDRESS=0x0F8d5F709c6916D332c08Be447a88e8551A43EFf npm run script:sepolia scripts/prepareUpgradeToAndCall.ts

Target new BookNFT implementation address: 0x0F8d5F709c6916D332c08Be447a88e8551A43EFf
Upgrade to and call data: 0x3659cfe60000000000000000000000000f8d5f709c6916d332c08be447a88e8551a43eff
```

#### At safe.global

- `New transcation`
- `Transaction Builder`
- Fill in the LikeProtocol Proxy address
- Paste the current version of LikeProtocol abi, which is find in last tag of `abi/LikeProtocol.json`
- The interface should show avalible function, pick `upgradeToAndCall`
- Paste the above script output of BookNFT address to `newImplementation`, `data` with the above output of `prepareUpgradeToAndCall.ts` script
- Submit the proposal
- Ask the owner to sign with attached expected payload for cross check.
- Run `npm run script:sepolia scripts/queryBookNFTImpl.ts` to verify after all party signed
