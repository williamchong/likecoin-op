# Project likenft-indexer

One Paragraph of project description goes here

## Pre-requisite

Install atlas db migration tool

```bash
curl -sSf https://atlasgo.sh | sh
```

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Development

Init env file

```bash
make secret
```

Start api server

```bash
make start-api
```

#### Workers

Tasks:

- `cmd/worker/task`
- `cmd/worker/cmd/worker.go`

To start worker on destinated queues

```bash
make start-worker-default
```

```bash
make start-worker-index-action
```

To start worker on all queues

```bash
make start-worker-all
```

Start task scheduler

```bash
make start-worker-scheduler
```

## MakeFile

Init env file

```bash
make secret
```

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Create DB container

```bash
make docker-run
```

Shutdown DB Container

```bash
make docker-down
```

DB Integrations Test:

```bash
make itest
```

Live reload the application:

```bash
make watch
```

Run the test suite:

```bash
make test
```

Clean up binary from the last build:

```bash
make clean
```
