GIT_SHORT_HASH ?= $(shell git rev-parse --short=8 HEAD)
TAG ?= $(GIT_SHORT_HASH)
DOCKER_REGISTRY ?= likecoin
DEPLOY_DEV_DOCKER_IMAGE_NAME ?= $(DOCKER_REGISTRY)/likecollective-indexer
DOCKER_BUILD_ARGS ?=

-include .env
export

PROJECT_NAME := likecollective-indexer

.PHONY: secret
secret:
	@curl -s https://raw.githubusercontent.com/oursky/devsecops-secret/v1.0/generate-secret.sh \
		| bash -s -- \
		.env.example .env

.PHONY: vendor
vendor:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $$(go env GOPATH)/bin v2.4.0
	go mod download

.PHONY: generate
generate:
	go generate ./...

.PHONY: fmt-go-changed
fmt-go-changed:
	git ls-files -mo | grep -e '.go$$' | sort | xargs \
		go run golang.org/x/tools/cmd/goimports@v0.36.0 -w -format-only -local likecollective-indexer

fmt-go:
	find ./cmd -name '*.go' \
		-not -name 'wire_gen.go' \
		-not -name '*_mock_test.go' \
		| sort \
		| xargs \
			go run golang.org/x/tools/cmd/goimports@v0.36.0 -w -format-only -local likecollective-indexer
	find ./internal -name '*.go' \
		-not -name 'wire_gen.go' \
		-not -name '*_mock_test.go' \
		| sort \
		| xargs \
			go run golang.org/x/tools/cmd/goimports@v0.36.0 -w -format-only -local likecollective-indexer

.PHONY: fmt
fmt: fmt-go

.PHONY: check-tidy
check-tidy:
	$(MAKE) fmt
	go mod tidy
	git status --porcelain . | grep '.*'; test $$? -eq 1

.PHONY: lint
lint:
	golangci-lint run ./cmd/... --timeout 7m
	go vet ./...

.PHONY: test
test:
	go test ./...

.PHONY: ci
ci:
	$(MAKE) check-tidy
	$(MAKE) lint
	$(MAKE) test

.PHONY: start-api
start-api:
	docker compose --profile api up -d && docker compose --profile api logs -f

.PHONY: stop-api
stop-api:
	docker compose --profile api stop && docker compose --profile api rm -f

.PHONY: migration
migration:
	atlas migrate diff "${MIGRATION_NAME}" \
		--dir file://ent/migrate/migrations \
		--to ent://ent/schema/ \
		--dev-url "docker://postgres/15/test?search_path=${DB_SCHEMA}"

.PHONY: migrate-up
migrate-up:
	@echo "Migrating up..."
	atlas migrate apply \
		--dir "file://ent/migrate/migrations" \
		--url "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable&search_path=${DB_SCHEMA}"

.PHONY: migrate-down
migrate-down:
	@echo "Migrating down..."
	atlas migrate down \
		--dir "file://ent/migrate/migrations" \
		--dev-url "docker://postgres/15/test?search_path=${DB_SCHEMA}" \
		--url "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable&search_path=${DB_SCHEMA}"

API_ARTIFACT := $(PROJECT_NAME)-api

$(API_ARTIFACT):
	go build -o $(API_ARTIFACT) ./cmd/api

.PHONY: build
build: $(PROJECT_NAME)-api

clean:
	rm -f $(API_ARTIFACT)

.PHONY: docker-image
docker-image:
	docker buildx build --platform linux/amd64 . \
		--build-arg 'BUILD_TAG=$(TAG)' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):latest' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):$(TAG)' \
		$(DOCKER_BUILD_ARGS)
