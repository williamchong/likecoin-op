
GIT_SHORT_HASH ?= $(shell git rev-parse --short=8 HEAD)
TAG ?= $(GIT_SHORT_HASH)
DEPLOY_DEV_DOCKER_IMAGE_NAME ?= chankiyu22/like-signer-backend
DOCKER_BUILD_ARGS ?=

ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

.PHONY: secret
secret:
	@curl -s https://raw.githubusercontent.com/oursky/devsecops-secret/v1.0/generate-secret.sh \
		| bash -s -- \
		.env.example .env

.PHONY: vendor
vendor:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v2.0.1
	go mod download
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: start
start:
	go run ./cmd/server

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build -o like-signer-backend ./cmd/server

.PHONY: build-worker-cli
build-worker-cli:
	go build -o like-signer-backend-worker-cli ./cmd/worker

.PHONY: fmt
fmt:
	find ./cmd -name '*.go' -not -name 'wire_gen.go' -not -name '*_mock_test.go' | sort | xargs goimports -w -format-only -local github.com/likecoin/like-signer-backend

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

.PHONY: docker-image
docker-image:
	docker buildx build --platform linux/amd64 . \
		--build-arg 'BUILD_TAG=$(TAG)' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):latest' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):$(TAG)' \
		$(DOCKER_BUILD_ARGS)
