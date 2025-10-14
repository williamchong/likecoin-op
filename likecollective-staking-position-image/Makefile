GIT_SHORT_HASH ?= $(shell git rev-parse --short=8 HEAD)
TAG ?= $(GIT_SHORT_HASH)
DOCKER_REGISTRY ?= likecoin
DEPLOY_DEV_DOCKER_IMAGE_NAME ?= $(DOCKER_REGISTRY)/likecollective-staking-position-image
DOCKER_BUILD_ARGS ?=

PYTHON_VERSION = python3.13

PY_MODULES = staking_image

WITH_DOCKER = true

ifeq ($(WITH_DOCKER), true)
POETRY_COMMAND = docker compose run --rm dev poetry
else
POETRY_COMMAND = poetry
endif

.PHONY: setup-dev
setup-dev:
	pip install poetry
	poetry install

.PHONY: format
format:
	$(POETRY_COMMAND) run black $(PY_MODULES)
	$(POETRY_COMMAND) run isort $(PY_MODULES)

.PHONY: dev
dev:
ifeq ($(WITH_DOCKER), false)
	poetry run fastapi dev staking_image/server.py
else
	$(MAKE) stop
	docker compose up -d dev
	docker compose logs -f dev
endif

.PHONY: stop
stop:
ifeq ($(WITH_DOCKER), true)
	docker compose stop dev
	docker compose rm -f dev
endif

.PHONY: poetry
poetry: ARGS = $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
poetry:
	$(POETRY_COMMAND) $(ARGS)

.PHONY: docker-image
docker-image:
	docker buildx build --platform linux/amd64 . \
		--build-arg 'BUILD_TAG=$(TAG)' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):latest' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):$(TAG)' \
		$(DOCKER_BUILD_ARGS)
