
GIT_SHORT_HASH ?= $(shell git rev-parse --short=8 HEAD)
TAG ?= $(GIT_SHORT_HASH)
DOCKER_REGISTRY ?= likecoin
DEPLOY_DEV_DOCKER_IMAGE_NAME ?= $(DOCKER_REGISTRY)/like-migration-backend
DOCKER_BUILD_ARGS ?=

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}` 

-include .env
export

.PHONY: secret
secret:
	@curl -s https://raw.githubusercontent.com/oursky/devsecops-secret/v1.0/generate-secret.sh \
		| bash -s -- \
		.env.example .env

.PHONY: vendor
vendor:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.55.2
	go mod download
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/rubenv/sql-migrate/...@latest

.PHONY: start
start:
	go run ./cmd/server

.PHONY: test
test:
	go test ./...

.PHONY: start-worker
start-worker:
	go run ./cmd/worker worker

.PHONY: start-scheduler
start-scheduler:
	go run ./cmd/worker scheduler

.PHONY: build
build:
	go build -o like-migration-backend ./cmd/server

.PHONY: build-cli
build-cli:
	go build -o like-migration-backend-cli ./cmd/cli

.PHONY: build-worker-cli
build-worker-cli:
	go build -o like-migration-backend-worker-cli ./cmd/worker

.PHONY: fmt
fmt:
	find ./cmd -name '*.go' -not -name 'wire_gen.go' -not -name '*_mock_test.go' | sort | xargs goimports -w -format-only -local github.com/likecoin/like-migration-backend

.PHONY: lint
lint:
	golangci-lint run ./cmd/... --timeout 7m
	go vet ./cmd/...

.PHONY: check-tidy
check-tidy:
	$(MAKE) fmt
	go mod tidy
	git status . --porcelain | grep '.*'; test $$? -eq 1

.PHONY: ci
ci: check-tidy lint test

.PHONY: create-migration
create-migration:
	go run github.com/rubenv/sql-migrate/...@latest new $(args)

.PHONY: run-migration
run-migration:
	go run github.com/rubenv/sql-migrate/...@latest up

.PHONY: abigen
abigen: SHELL=/bin/bash
abigen:
	mkdir -p pkg/likenft/evm/like_protocol
	mkdir -p pkg/likenft/evm/book_nft
	mkdir -p pkg/likecoin/evm/likecoin
	go run github.com/ethereum/go-ethereum/cmd/abigen@latest \
		--abi=<(cat ../abi/LikeProtocol.abi.json') \
		--pkg=like_protocol \
		--out=pkg/likenft/evm/like_protocol/like_protocol.go
	go run github.com/ethereum/go-ethereum/cmd/abigen@latest \
		--abi=<(cat ../abi/BookNFT.abi.json') \
		--pkg=book_nft \
		--out=pkg/likenft/evm/book_nft/book_nft.go

.PHONY: docker-image
docker-image:
	docker buildx build --platform linux/amd64 . \
		--build-arg 'BUILD_TAG=$(TAG)' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):latest' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):$(TAG)' \
		$(DOCKER_BUILD_ARGS)
