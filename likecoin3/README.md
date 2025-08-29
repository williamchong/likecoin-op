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

Make deployment to superism: `npx hardhat ignition deploy --network superism2 ignition/modules/Likecoin.ts --parameters ignition/parameters.json`
