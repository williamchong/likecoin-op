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

```
npx hardhat ignition deploy \
    ignition/modules/{module}.ts \
    --strategy create2 --parameters ignition/parameters.local.json \
    --network superism1
```

If you want to upgrade the contract, you may want to swipe the implementation future as below.

```
cat impl-future | xargs -t -L 1 npx hardhat ignition wipe chain-901
```

### Deployment

Make `LikeCollective` deployment to testnet/mainnet:

```
DOTENV_CONFIG_PATH=.env \
    npx hardhat ignition deploy \
    ignition/modules/LikeCollective.ts \
    --verify --strategy create2 --parameters ignition/parameters.json \
    --network baseSepolia
```

#### LikeProtocol

Ref command for deploying `LikeProtocolV1`, useful for dev to set local state.

Note: the `LikeProtocol` and `BookNFT` ignition are there for dev purpose, when
deploy to public chain, create a new `BookNFTV{num}.ts` for tracking
`deployed_address.json`

```
DOTENV_CONFIG_PATH=.env \
    npx hardhat ignition deploy \
    ignition/modules/LikeProtocolV1.ts \
    --verify --strategy create2 \
    --parameters ignition/parameters.json --network baseSepolia
```

For upgrading the LikeProtocol

Reference swipe comment for testing:

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
    --rpc-url https://base-sepolia.g.alchemy.com/v2/<redacted> \
    --account likecoin-deployer.eth
```

For upgrading the BookNFT

Reference swipe comment for testing:

```
npx hardhat ignition wipe chain-84532 BookNFT#BookNFTImpl
```

Create a new `BookNFTV{num}.ts` for tracking in `deployed_address.json`

```
DOTENV_CONFIG_PATH=.env \
    npx hardhat ignition deploy \
    ignition/modules/BookNFTV2.ts \
    --verify --strategy create2 \
    --parameters ignition/parameters.json --network baseSepolia
```

```
cast send 0xfb5cbb1973a092E6C77af02EA1E74B14870AbeC5 \
    "upgradeTo(address newImplementation)" \
    0x98AA03bAdD6a013bcc79CDfc00131e0d90E65AC6 \
    --rpc-url https://base-sepolia.g.alchemy.com/v2/<redacted> \
    --account likecoin-deployer.eth
```

#### Verification

ignition should verify the contract on etherscan & sourcify. In case the request fails (like rate limit), run `DOTENV_CONFIG_PATH=.env npx hardhat verify --network sepolia 0x1EE5DD1794C28F559f94d2cc642BaE62dC3be5cf`, in some testnet, etherscan will fail while sourcify will success.

## Simulation

### Simulate with adhoc node

The simulation will be run immediately but will be destroied after simulation completed.

```bash
npx hardhat simulate:likecollective \
    --outputfile simulate/likecollective/simulations/case1.output.json \
    simulate/likecollective/simulations/case1.yaml
```

### Simulate with local node

A local eth node process should be setup.

```bash
docker compose up
```

The simulation involves contract deployments so may take long time to setup.

```bash
npx hardhat simulate:likecollective --network localhost \
    --outputfile simulate/likecollective/simulations/case1.output.json \
    simulate/likecollective/simulations/case1.yaml
```

The state will be persisted as long as the node is not shut down.
The node is then be able to be queried for further processing.

For more details, can see [likecollective-indexer](../likecollective-indexer/README.md)
