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
