# LikeCoin Migration Backend

## Program flow

```mermaid
sequenceDiagram
    Browser->>Backend: Send used signature with intent to migrate
    Backend->>Cosmos: Verify the action on chain
    Backend->>Backend: Verify the action EVM personal_sign message
    Backend->>Signer: Translate the verified intention as a EVM payload, i.e. LikeNFT mint/LikeERC20 send
    Signer->>EVM: Submit the transaction
    Backend->>Signer: Query the transaction status
    Backend->>EVM: Query the transaction status
    Backend->>Browser: Send result success/fails
```

## Note

- The migration program will paid for the gas fee of the evm transaction.
- Backend will hold the logic while have no ability to sign the transaction. It's for segregration of duty.
- Signer will
    - Hold the hot wallet
    - Sign submitted transaction
    - Detect low balance
    - Detect any anomaly of the transaction

## Project setup

```sh
$ make secrets
$ make vendor
```

## Infra

```sh
$ docker compose up
```

## Migration

```sh
$ make run-migration
```

## Dev

```sh
$ make start
```

## Create migration

```sh
$ make create-migration
```
