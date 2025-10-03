# Likecoin V3

Likecoin contract for superchain deployment

## Usage

### Pre-commit

To run pre-commit check, execute the following

```
npm run format
npm run lint
```

### Running Tests

To run all the tests in the project, execute the following command:

```shell
npm run test
```

### Local superchain development

```
docker compose up
```

To tail individual chain logs, run something like: `docker compose exec -it superism tail -f /tmp/anvil-chain-902-{timestamp}`

### Deployment

Make deployment to superism & testnet/mainnet:

```
npm run deploy:noverify -- --network superism1
DOTENV_CONFIG_PATH=.env npm run deploy -- --network baseSepolia
```

For upgrading the LikeProtocol
Swipe if there is previous deployment: 

```
npx hardhat ignition wipe chain-84532 LikeProtocolModule#LikeProtocolImpl
```
```
DOTENV_CONFIG_PATH=.env \
    npx hardhat ignition deploy \
    ignition/modules/LikeProtocol.ts \
    --verify --strategy create2 \
    --parameters ignition/parameters.json --network baseSepolia
```

Manually switch version, it's not managed by ignition


```
cast send 0xfb5cbb1973a092E6C77af02EA1E74B14870AbeC5 \
    "upgradeToAndCall(address newImplementation, bytes data)" \
    0x05857EE837AB29fF79C7BB1d4c642b2C9dd10FA5 \
    0x \
    --rpc-url https://base-sepolia.g.alchemy.com/v2/OM1XAvx0Dwavrz6MQn5aG \
    --account likecoin-deployer.eth 
```

cast

#### Verification

ignition should verify the contract on etherscan & sourcify. In case the request fails (like rate limit), run `DOTENV_CONFIG_PATH=.env npx hardhat verify --network sepolia 0x1EE5DD1794C28F559f94d2cc642BaE62dC3be5cf`, in some testnet, etherscan will fail while sourcify will success.

#### Upgrading the staking implementation

For local superism deploy, normally you will run following

```
npm run deploy:local -- --network superism1
```

If you want to upgrade the contract, you may want to swipe the implementation future as below.

```
cat impl-future | xargs -t -L 1 npx hardhat ignition wipe chain-901
```
