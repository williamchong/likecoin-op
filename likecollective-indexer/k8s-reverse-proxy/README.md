# K8s Reserve Proxy

Ref: https://www.notion.so/oursky/Reverse-HTTPS-Proxy-with-Pandawork

## Setup

```sh
$ cp .env.example .env
```

Provide namespace to the env file

## Start ssh server and establish a port-forwarding

```sh
$ make rproxy-container
```

## Establish a ssh reverseh proxy

```sh
$ make rproxy-ssh
```

## Clean up

```sh
$ make rproxy-container/clean
```

## Test

```sh
$ make start-test-server
```

The following url should be able to return text from `test/index.html`.

```sh
$ make test
```
