.PHONY: setup
setup:
	npm install
	cp .env.local .env

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
ci: build
	npm run check-format
	npm run lint
	npm run test

.PHONY: deploy-local
deploy-local:
	npm run script:local scripts/deploy.ts
	sleep 2
	npm run script:local scripts/transferOwner.ts
