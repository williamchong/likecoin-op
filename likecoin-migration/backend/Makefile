GIT_SHORT_HASH ?= $(shell git rev-parse --short=8 HEAD)
TAG ?= $(GIT_SHORT_HASH)
DEPLOY_DEV_DOCKER_IMAGE_NAME ?= chankiyu22/likecoin-migration-backend
DOCKER_BUILD_ARGS ?=

.PHONY: vendor
vendor:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.55.2
	go mod download
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: start
start:
	go run ./cmd/server

.PHONY: build
build:
	go build -o likecoin-migration-backend ./cmd/server

.PHONY: fmt
fmt:
	find ./cmd -name '*.go' -not -name 'wire_gen.go' -not -name '*_mock_test.go' | sort | xargs goimports -w -format-only -local github.com/likecoin/likecoin-migration-backend

.PHONY: lint
lint:
	golangci-lint run ./cmd/... --timeout 7m
	go vet ./cmd/...

.PHONY: check-tidy
check-tidy:
	$(MAKE) fmt
	go mod tidy
	git status . --porcelain | grep '.*'; test $$? -eq 1

.PHONY: docker-image
docker-image:
	docker buildx build --platform linux/amd64 . \
		--build-arg 'BUILD_TAG=$(TAG)' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):latest' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):$(TAG)' \
		$(DOCKER_BUILD_ARGS)
