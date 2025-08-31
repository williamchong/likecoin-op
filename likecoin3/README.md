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

#### Verification

ignition should verify the contract on etherscan & sourcify. In case the request fails (like rate limit), run `DOTENV_CONFIG_PATH=.env npx hardhat verify --network sepolia 0x1EE5DD1794C28F559f94d2cc642BaE62dC3be5cf`, in some testnet, etherscan will fail while sourcify will success.