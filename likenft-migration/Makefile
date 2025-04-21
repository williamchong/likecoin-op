GIT_SHORT_HASH ?= $(shell git rev-parse --short=8 HEAD)
TAG ?= $(GIT_SHORT_HASH)
DOCKER_REGISTRY ?= likecoin
DEPLOY_DEV_DOCKER_IMAGE_NAME ?= $(DOCKER_REGISTRY)/likenft-migration
DOCKER_BUILD_ARGS ?=

.PHONY: setup
setup:
	npm install

.PHONY: dev
dev:
	npm run dev

.PHONY: build
build:
	npm run generate

.PHONY: lint
lint:
	npm run lint

.PHONY: ci
ci:
	npm run check-format
	npm run lint
	npm run check-translations

.PHONY: docker-image
docker-image:
	docker buildx build --platform linux/amd64 . \
		--build-arg 'BUILD_TAG=$(TAG)' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):latest' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):$(TAG)' \
		$(DOCKER_BUILD_ARGS)
