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

.PHONY: gen-go-binding
gen-go-binding:
	npx solcjs \
		--include-path node_modules/ \
		--base-path . \
		--abi contracts/LikeProtocol.sol \
		--optimize \
		--optimize-runs=200 \
		-o build
	npx solcjs \
		--include-path node_modules/ \
		--base-path . \
		--bin contracts/LikeProtocol.sol \
		--optimize \
		--optimize-runs=200 \
		-o build
	npx solcjs \
		--include-path node_modules/ \
		--base-path . \
		--abi contracts/LikeNFTClass.sol \
		--optimize \
		--optimize-runs=200 \
		-o build
	npx solcjs \
		--include-path node_modules/ \
		--base-path . \
		--bin contracts/LikeNFTClass.sol \
		--optimize \
		--optimize-runs=200 \
		-o build
	go run github.com/ethereum/go-ethereum/cmd/abigen@latest \
		--bin=build/contracts_LikeProtocol_sol_LikeProtocol.bin \
		--abi=build/contracts_LikeProtocol_sol_LikeProtocol.abi \
		--pkg=like_protocol \
		--out=build/like_protocol.go
	go run github.com/ethereum/go-ethereum/cmd/abigen@latest \
		--bin=build/contracts_LikeNFTClass_sol_LikeNFTClass.bin \
		--abi=build/contracts_LikeNFTClass_sol_LikeNFTClass.abi \
		--pkg=likenft_class \
		--out=build/likenft_class.go

.PHONY: deploy-local
deploy-local:
	npm run deploy:local
