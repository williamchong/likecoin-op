GIT_SHORT_HASH ?= $(shell git rev-parse --short=8 HEAD)
TAG ?= $(GIT_SHORT_HASH)
DOCKER_REGISTRY ?= likecoin
DEPLOY_DEV_DOCKER_IMAGE_NAME ?= $(DOCKER_REGISTRY)/likecollective-staking-position-image
DOCKER_BUILD_ARGS ?=

PYTHON_VERSION = python3.13

PY_MODULES = staking_image

VENV = venv

.PRECIOUS: $(VENV)
$(VENV):
	$(PYTHON_VERSION) -m venv $(VENV)

.PHONY: setup-dev
setup-dev: $(VENV)
	source $(VENV)/bin/activate; \
		pip install poetry; \
		poetry install

.PHONY: format
format: $(VENV)
	source $(VENV)/bin/activate; \
		poetry run black $(PY_MODULES); \
		poetry run isort $(PY_MODULES)

.PHONY: dev
dev: $(VENV)
	source $(VENV)/bin/activate; \
		fastapi dev staking_image/server.py

.PHONY: docker-image
docker-image:
	docker buildx build --platform linux/amd64 . \
		--build-arg 'BUILD_TAG=$(TAG)' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):latest' \
		-t '$(DEPLOY_DEV_DOCKER_IMAGE_NAME):$(TAG)' \
		$(DOCKER_BUILD_ARGS)
