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

# Batch mint iscns

.PRECIOUS: batch-data/output/%/logs
batch-data/output/%/logs: batch-data/output/%/iscn.json
	mkdir -p $(shell dirname -- $@)
	cp $< data/iscn.json
	$(MAKE) mint-nft-new-iscn 2>&1 | tee $@

.PRECIOUS: batch-data/output/%/iscn.json
batch-data/output/%/iscn.json:
	mkdir -p $(shell dirname -- $@)
	sed "s/{id}/$*/g" batch-data/iscn.json.template > $@

.PHONY: batch-data
batch-data: FROM?=0
batch-data: TO?=10
batch-data: IDS?=$(shell seq $(FROM) 1 $(TO))
batch-data: RESULT_FILES=$(patsubst %,$@/output/%/logs,$(IDS))
batch-data:
	$(MAKE) $(RESULT_FILES)

.PHONY: batch-mint-nft-new-iscns
batch-mint-nft-new-iscns: FROM?=0
batch-mint-nft-new-iscns: TO?=10
batch-mint-nft-new-iscns: IDS?=$(shell seq $(FROM) 1 $(TO))
batch-mint-nft-new-iscns:
	$(MAKE) batch-data FROM=$(FROM) TO=$(TO) IDS="$(IDS)"

.PHONY: clean-batch-data
clean-batch-data:
	rm -rf batch-data/output
