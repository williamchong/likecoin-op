# Likecollective indexer

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Development

Init env file

```bash
make secret
```

Install tools

```bash
make vendor
```

Start api server

```bash
make start-api
```

Linting

```bash
make lint
```

Stop api server

```bash
make stop-api
```

Run the test suite:

```bash
make test
```

## MakeFile

Build the application

```bash
make build
```

Clean up binary from the last build:

```bash
make clean
```

### Alchemy config

Quick reference for setting up webhook in the alchemy webhook

URL For LikeCollective contract should be
`https://indexer-base.v3.like.co/collective/alchemy/like-collective/ethlog`

Webhooktype: Custom
Query as below: (all event)
```
{
  block {
    hash,
    number,
    timestamp,
    logs(filter: {addresses: ["0x4506Ac2dD1e9A470d92a3D1656E1a99C676E1c8E"]}) {
      # Account is the account which generated this log - this will always be a contract account.
      account {
        address,
      },
      # Topics is a list of 0-4 indexed topics for the log.
      topics,
      # Data is unindexed data for this log.
      data,
      # Transaction is the transaction that generated this log entry.
      transaction {
        hash,
        index,
      },
      index,
    }
  }
}
```

URL for LikeStakePosition should be
`https://indexer-base.v3.like.co/collective/alchemy/like-stake-position/ethlog`
Webhooktype: Custom
Query as below: (all event)
```
{
  block {
    hash,
    number,
    timestamp,
    logs(filter: {addresses: ["0x508610D3009cda82Ac1a40D2b322Ed31932D16b1"]}) {
      # Account is the account which generated this log - this will always be a contract account.
      account {
        address,
      },
      # Topics is a list of 0-4 indexed topics for the log.
      topics,
      # Data is unindexed data for this log.
      data,
      # Transaction is the transaction that generated this log entry.
      transaction {
        hash,
        index,
      },
      index,
    }
  }
}
```