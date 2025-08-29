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

API_ARTIFACT := $(PROJECT_NAME)-api

$(API_ARTIFACT):
	go build -o $(API_ARTIFACT) ./cmd/api

.PHONY: build
build: $(PROJECT_NAME)-api

clean:
	rm -f $(API_ARTIFACT)
