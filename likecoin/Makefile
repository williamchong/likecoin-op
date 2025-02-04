.PHONY: setup
setup:
	npm install

.PHONY: lint
lint:
	npm run lint

.PHONY: format
format:
	npm run format

.PHONY: test
test:
	npm run test

.PHONY: ci
ci:
	npm run check-format
	npm run lint
	npm run test

.PHONY: deploy-local
deploy-local:
	npm run deploy:local
