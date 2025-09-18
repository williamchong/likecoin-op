.PHONY: setup
setup:
	npm install

.PHONY: build
build:
	npm run build

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
	DOTENV_CONFIG_PATH=.env npm run deploy:local -- --network localhost

.PHONY: mint-local
mint-local:
	npx hardhat mint --network localhost --amount $(AMOUNT) --to $(TO)
