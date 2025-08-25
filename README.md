# Likecoin 3.0

This repo contains the code for the Likecoin 3.0, which is hosting on the optimism chain.

## Folder overview

- `likecoin`: ERC20 token
- `likecoin-migration`: the migration frontend, from cosmos like to evm like
- `likenft`: LikeProtocol smart contract
- `likenft-indexer`: provide support for downstream application to quert NFT around LikeProtocol.
- `likenft-migration`: the code for the likerID and likeNFT migration frontend, from cosmos x/nft to evm likenft
- `signer-backend`: Signing the EVM TX for migration
- `migration-backend`: the code for the backend program, for migrating the data from cosmos likecoin and likecoin nft
- `migration-admin`: Simple CRUD interface for CX to support end user migration
- `operation`: holding script for operation support
- `deploy`: the scripts for deploying the migration program

## Quick start up on backend

To setup the evn files and kick start basic backend for the first time
```
make setup
make start
```

The infra will be setup with the following services

- migration-backend
- migration-backend-worker
- migration-backend-scheduler
- signer-backend
- eth-node
    - The contract will also be deployed
- db-migration-backend
- db-signer-backend
- redis

The console will stream the backend logs

To deploy the smartcontract locally, run following in another console
```
make local-contracts
```

To stop services,

```
$ make stop
```

> Please note that the `eth-node` have no mechanism to persist state.
> To prevent stucking at nonce inconsistent, after restarting the eth-node for any reason,
> the dbs related to transactions should also be reset.
>
> ```
> $ make clean-transaction-volumes
> ```
>
> The stop command above has included this command

## Setup on frontends

For respective frontend, `likenft-migration`, `likecoin-migration` and `migration-admin`, please naigate to respective folder and follow instruction there.
