.PHONY: setup
setup:
	npm install

.PHONY: lint
lint:
	npm run lint

.PHONY: format
format:
	npm run format

.PHONY: build
build:
	npm run build

.PHONY: test
test:
	npm run test

.PHONY: ci
ci:
	cp .env.local .env
	npm run check-format
	npm run lint
	npm run test

.PHONY: deploy-local
deploy-local:
	npm run deploy:local
