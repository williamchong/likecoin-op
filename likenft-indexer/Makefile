GIT_SHORT_HASH ?= $(shell git rev-parse --short=8 HEAD)
TAG ?= $(GIT_SHORT_HASH)
DOCKER_REGISTRY ?= likecoin
DEPLOY_DEV_DOCKER_IMAGE_NAME ?= $(DOCKER_REGISTRY)/likenft-indexer
DOCKER_BUILD_ARGS ?=

-include .env
export

.PHONY: secret
secret:
	@curl -s https://raw.githubusercontent.com/oursky/devsecops-secret/v1.0/generate-secret.sh \
		| bash -s -- \
		.env.example .env

# Build the application
all: build test
templ-install:
	@if ! command -v templ > /dev/null; then \
		read -p "Go's 'templ' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/a-h/templ/cmd/templ@latest; \
			if [ ! -x "$$(command -v templ)" ]; then \
				echo "templ installation failed. Exiting..."; \
				exit 1; \
			fi; \
		else \
			echo "You chose not to install templ. Exiting..."; \
			exit 1; \
		fi; \
	fi
tailwind-install:
	@if [ ! -f tailwindcss ]; then curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o tailwindcss; fi
	
	@chmod +x tailwindcss

generate:
	@echo "Generating..."
	go generate ./ent
	go generate ./openapi

.PHONY: abigen
abigen: SHELL=/bin/bash
abigen:
		mkdir -p internal/evm/like_protocol
		mkdir -p internal/evm/book_nft
		go run github.com/ethereum/go-ethereum/cmd/abigen@latest \
			--abi=<(cat ../abi/LikeProtocol.abi.json) \
			--pkg=like_protocol \
			--out=internal/evm/like_protocol/like_protocol.go
		go run github.com/ethereum/go-ethereum/cmd/abigen@latest \
			--abi=<(cat ../abi/BookNFT.abi.json) \
			--pkg=book_nft \
			--out=internal/evm/book_nft/book_nft.go

.PHONY: start-api
start-api:
	docker compose --profile api up -d && docker compose --profile api logs -f

.PHONY: stop-api
stop-api:
	docker compose --profile api stop && docker compose --profile api rm -f

.PHONY: start-scheduler
start-scheduler:
	docker compose --profile scheduler up -d && docker compose --profile scheduler logs -f

.PHONY: stop-scheduler
stop-scheduler:
	docker compose --profile scheduler stop && docker compose --profile scheduler rm -f

.PHONY: start-worker
start-worker:
	docker compose --profile worker up -d && docker compose --profile worker logs -f

.PHONY: stop-worker
stop-worker:
	docker compose --profile worker stop && docker compose --profile worker rm -f

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

.PHONY: fmt
fmt:
	find ./cmd -name '*.go' -not -name 'wire_gen.go' -not -name '*_mock_test.go' | sort | xargs goimports -w -format-only -local likenft-indexer
	find ./ent/schema -name '*.go' | sort | xargs goimports -w -format-only -local likenft-indexer
	find ./internal -name '*.go' | sort | xargs goimports -w -format-only -local likenft-indexer

build: tailwind-install templ-install generate
	@echo "Building..."
	templ generate
	./tailwindcss -i cmd/web/styles/input.css -o cmd/web/assets/css/output.css
	go build -o main cmd/api/main.go

.PHONY: build-api
build-api:
	go build -o likenft-indexer-api ./cmd/api

.PHONY: build-cli
build-cli:
	go build -o likenft-indexer-cli ./cmd/cli

.PHONY: build-worker-cli
build-worker-cli:
	go build -o likenft-indexer-worker-cli ./cmd/worker

# Run the application
run:
	@go run cmd/api/main.go
# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: docker-image
docker-image:
	docker buildx build --platform linux/amd64 . \
		--build-arg 'BUILD_TAG=$(TAG)' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):latest' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):$(TAG)' \
		$(DOCKER_BUILD_ARGS)

.PHONY: all build run test clean watch tailwind-install docker-run docker-down itest templ-install
