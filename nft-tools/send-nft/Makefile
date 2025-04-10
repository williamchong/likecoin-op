NODE_ENV ?= docker run --rm -v $$PWD:/src -w /src  --env-file .env node:14-alpine

-include .env
export

ifeq (${IS_TESTNET}, TRUE)
	RPC_ENDPOINT=https://node.testnet.like.co
else
	RPC_ENDPOINT=https://mainnet-node.like.co
endif

.PHONY: npm
npm: ARGS=$(filter-out $@, $(MAKECMDGOALS))
npm:
	$(NODE_ENV) npm $(ARGS)

.PHONY: npx
npx: ARGS=$(filter-out $@, $(MAKECMDGOALS))
npx:
	$(NODE_ENV) npx $(ARGS)

.PHONY: send-nft
send-nft:
	$(NODE_ENV) node index.js

.PHONY: query-nft-event
query-nft-event: ARGS=$(filter-out $@, $(MAKECMDGOALS))
query-nft-event: CLASSID=$(word 1, $(ARGS))
query-nft-event: NFTID=$(word 2, $(ARGS))
query-nft-event:
	curl '$(RPC_ENDPOINT)/likechain/likenft/v1/event?class_id=$(CLASSID)&nft_id=$(NFTID)&action_type=%2Fcosmos.nft.v1beta1.MsgSend&action_type=buy_nft&action_type=sell_nft&action_type=new_class&pagination.limit=100&pagination.reverse=true'
