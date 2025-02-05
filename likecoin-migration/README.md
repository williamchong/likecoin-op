# Likecoin Migration

This repo contains the code for the migration program, from cosmos like to evm likecoin.

## Program flow

```mermaid
sequenceDiagram
    Browser->>Browser: Get user's cosmos address
    Browser->>Browser: Get user's evm address
    Browser->>Backend: Send the cosmos and evm address to backend
    Backend->>Browser: Request user using EVM wallet to signin
    Browser->>Browser: Request user to sign the cosmos send transaction with signed EVM login message
    Browser->>Cosmos: Send all LIKE to destinated address
    Browser->>Backend: Send the TX hash for verify
    Backend->>Cosmos: Querying the cosmos transaction to verify
    Backend->>EVM: Send the LIKE to user's evm address
    Backend->>Browser: Return the EVM TX hash
    Browser->>Browser: Show the migration result
```

Note:

- The migration program will paid for the gas fee of the evm transaction.
- The destinated address of migrated LIKE will be redelegate across the validator for protecting the security of the chain.

## Pre-requisite

Node 20

## Dev

```bash
$ make setup
```

```bash
$ make dev
```
