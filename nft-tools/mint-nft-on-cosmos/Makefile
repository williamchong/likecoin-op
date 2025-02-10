.PHONY: setup
setup:
	cp .env.example .env

NODE_ENV ?= docker run --rm -v $$PWD:/src -w /src  --env-file .env node:14-alpine

.PHONY: npm
npm: ARGS=$(filter-out $@, $(MAKECMDGOALS))
npm:
	$(NODE_ENV) npm $(ARGS)

.PHONY: npx
npx: ARGS=$(filter-out $@, $(MAKECMDGOALS))
npx:
	$(NODE_ENV) npx $(ARGS)

.PHONY: mint-nft-new-iscn
mint-nft-new-iscn: NFT_COUNT=$(shell echo $$(($$( grep -c . data/nfts.csv) - 1)))
mint-nft-new-iscn:
	$(NODE_ENV) node index.js --nft-count $(NFT_COUNT) --create-new-iscn

.PHONY: mint-nft-existing-iscn
mint-nft-existing-iscn: NFT_COUNT=$(shell echo $$(($$( grep -c . data/nfts.csv) - 1)))
mint-nft-existing-iscn:
	$(NODE_ENV) node index.js --nft-count $(NFT_COUNT) --iscn-id ${ISCN_ID}
